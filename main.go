package main

import (
	"os"
	"strings"

	"github.com/ravilock/go-cowsay/cowsay"
)

const LINE_LENGTH_LIMIT = 40

const (
	FIRST_LINE = iota
	NORMAL_LINE
	LAST_LINE
)

func main() {
	text := strings.Join(os.Args[1:], " ")

	cowsay.Say(text)
}