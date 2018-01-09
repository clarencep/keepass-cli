package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var inputFilePath string
	var outputFilePath string

	flag.StringVar(&inputFilePath, "input", "", "specify the input .kdbx file")
	flag.StringVar(&outputFilePath, "output", "", "specify the output .kdbx file")
	flag.Parse()

	if inputFilePath == "" {
		dieWithError("Input .kdbx file should not be empty!")
	}

	if outputFilePath == "" {
		dieWithError("Output .kdbx file should not be empty!")
	}

	kp, err := OpenKeepassDb(inputFilePath)
	failIfErrNotNil(err)

	defer kp.Close()

	oldPasswd, err := readPassword("Please input your old password: ")
	failIfErrNotNil(err)

	err = kp.UnlockWithPassword(oldPasswd)
	failIfErrNotNil(err)

	newPasswd, err := readPassword("Please input your new password: ")
	failIfErrNotNil(err)

	newPasswd2, err := readPassword("Please input your new password again: ")
	failIfErrNotNil(err)

	if newPasswd2 != newPasswd {
		dieWithError("Your new password confirmation failed!")
	}

	fmt.Println("Changing password...")
	err = kp.ChPassword(newPasswd)
	failIfErrNotNil(err)

	fmt.Printf("Writing to %s...\n", outputFilePath)
	err = kp.SaveTo(outputFilePath)
	failIfErrNotNil(err)

	fmt.Println("Done.")
}

func failIfErrNotNil(err interface{}) {
	if err != nil {
		log.Fatal(err)
	}
}

func dieWithError(msg string) {
	fmt.Fprintln(os.Stderr, "Error: "+msg)
	os.Exit(1)
}
