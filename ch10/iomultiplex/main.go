package main

import (
	"fmt"
	"syscall"
)

func main() {
	kq, err := syscall.Kqueue()
	if err != nil {
		panic(err)
	}

	fd, err := syscall.Open("./test", syscall.O_RDONLY, 0) // 監視対象
	if err != nil {
		panic(err)
	}

	evl := syscall.Kevent_t{
		Ident:  uint64(fd),
		Filter: syscall.EVFILT_VNODE,
		Flags:  syscall.EV_ADD | syscall.EV_ENABLE | syscall.EV_ONESHOT,
		Fflags: syscall.NOTE_DELETE | syscall.NOTE_WRITE,
		Data:   0,
		Udata:  nil,
	} // 監視対象のイベント詳細

	for {
		events := make([]syscall.Kevent_t, 10)

		nev, err := syscall.Kevent(kq, []syscall.Kevent_t{evl}, events, nil)
		if err != nil {
			panic(err)
		}

		for i := 0; i < nev; i++ {
			fmt.Printf("Event [%d] -> %+v\n", i, events[i])
		}
	}
}
