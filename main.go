package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/ravilock/go-cowsay/cowsay"
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
