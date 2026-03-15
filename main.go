// Package main provides the entry point for the cowsay application.
package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/ravilock/go-cowsay/cowsay"
)

const lineLengthLimit = 40

const (
	firstLine = iota
	normalLine
	lastLine
)

func getInput() string {
	if len(os.Args) > 1 {
		return strings.Join(os.Args[1:], " ")
	}

	data := os.Stdin
	scan := bufio.NewScanner(data)
	var stdinSlice []string
	for scan.Scan() {
		stdinSlice = append(stdinSlice, scan.Text()+"\n")
	}
	return strings.TrimSuffix(strings.Join(stdinSlice, ""), "\n")
}

func main() {
	text := getInput()

	cowsay.Say(text)
}
