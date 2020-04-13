package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"
	"log"
	"os"
)

// png format
// ----------------------------------------------------------------------------------------------------
// | signature(8bytes) | length(4bytes) | type(4byte) | data | CRC(4byte) | length, type, data, CRC...
// ----------------------------------------------------------------------------------------------------

func readChunks(file *os.File) []io.Reader {
	var chunks []io.Reader

	file.Seek(8, 0) // PNG の最初のシグネチャ8バイト分を読み飛ばす
	var offset int64 = 8

	for {
		var l int32
		err := binary.Read(file, binary.BigEndian, &l)
		if err == io.EOF {
			break
		}
		// offset - length(4) + type(4) + data(l) + CRC(4) までで SectionReader を作成
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(l)+12))

		// 今は length を読み終えたところなので、
		// 次の offset まで読み飛ばし (type(4) + data(l) + CRC(4))
		offset, _ = file.Seek(int64(l+8), 1)
	}

	return chunks
}

func dumpChunks(chunk io.Reader) {
	var length int32
	// chunk (length + type + data + CRC) から length を読む
	binary.Read(chunk, binary.BigEndian, &length)

	// 今は length を読み終えたところなので、次は type(4)
	buffer := make([]byte, 4)
	chunk.Read(buffer)

	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)

	if bytes.Equal(buffer, []byte("tEXt")) {
		rawText := make([]byte, length)
		chunk.Read(rawText)
		fmt.Println(string(rawText))
	}
}

func textChunk(text string) io.Reader {
	byteData := []byte(text)

	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, int32(len(byteData)))
	buffer.WriteString("tEXt")
	buffer.Write(byteData)

	// CRC を計算
	crc := crc32.NewIEEE()
	io.WriteString(crc, "tEXt")
	binary.Write(&buffer, binary.BigEndian, crc.Sum32())

	return &buffer
}

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	newFile, err := os.Create("Lenna2.png")
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	chunks := readChunks(file)

	// write signature
	io.WriteString(newFile, "\x89PNG\r\n\x1a\n")
	// write IHDR
	io.Copy(newFile, chunks[0])
	// add text chunk
	io.Copy(newFile, textChunk("ASCII PROGRAMMING++"))
	// add remaining chunk
	for _, chunk := range chunks[1:] {
		io.Copy(newFile, chunk)
	}

	newChunks := readChunks(newFile)

	for _, chunk := range newChunks {
		dumpChunks(chunk)
	}
}
