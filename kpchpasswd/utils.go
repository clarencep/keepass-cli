package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

func readPassword(prompt string) (passwd string, err error) {
	fmt.Fprint(os.Stdout, prompt)
	buf, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	passwd := strings.TrimSpace(string(buf))
	return
}
