package main

import (
	"fmt"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	"github.com/tmc/keyring"
)

func main() {
	secretValue, err := keyring.Get("progo-keyring-test", "password")
	if err == keyring.ErrNotFound {
		fmt.Printf("Secret Value is not found. Please Type:")
		pw, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			panic(err)
		}

		err = keyring.Set("progo-keyring-test", "password", string(pw))
		if err != nil {
			panic(err)
		}
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("Secret Value: %s\n", secretValue)
	}
}
