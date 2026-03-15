package cowsay

import (
	"fmt"
	"strconv"
	"strings"
)

// Package cowsay provides a simple cow-themed text formatter.

const lineLengthLimit = 39
const space = " "

const (
	firstLine = iota
	normalLine
	lastLine
)

func Say(text string) {
	text = strings.ReplaceAll(text, "\n", " ")

	cowSay := fmt.Sprintf("\n%s%s", buildBaloon(text), buildCow())

	fmt.Println(cowSay)
}

func buildBaloon(text string) string {
	lines := chunkString(text)
	baloonWidth := getBaloonWidth(lines)

	baloonStart := buildBaloonStart(baloonWidth)
	baloonText := buildBaloonText(lines, baloonWidth)
	baloonEnd := buildBaloonEnd(baloonWidth)

	return baloonStart + baloonText + baloonEnd
}

func chunkString(s string) []string {
	var chunks []string

	stringSize := len(s)

	for startIndex := 0; startIndex < stringSize; {
		finalIndex := startIndex + lineLengthLimit
		if finalIndex > stringSize {
			finalIndex = stringSize
		}

		lastSpaceIndex := strings.LastIndex(s[startIndex:finalIndex], space)
		if lastSpaceIndex != -1 && (finalIndex-startIndex) == lineLengthLimit {
			finalIndex = startIndex + lastSpaceIndex + 1
		}

		line := s[startIndex:finalIndex]
		line = strings.Trim(line, " ")
		chunks = append(chunks, line)
		startIndex = finalIndex
	}

	return chunks
}

func getBaloonWidth(lines []string) int {
	var biggestLineSize int

	for _, line := range lines {
		biggestLineSize = max(biggestLineSize, len(line))
	}

	return biggestLineSize
}

func buildBaloonText(lines []string, biggestLineSize int) string {
	var builder strings.Builder

	numberOfLines := len(lines)

	for lineIndex, line := range lines {
		line := fixLineSize(line, biggestLineSize)
		lineType := getLineType(lineIndex, numberOfLines)
		lineDelimiters := getLineDelimiters(lineType, numberOfLines)

		builder.WriteString(lineDelimiters[0])
		builder.WriteString(line)
		builder.WriteString(lineDelimiters[1])
		builder.WriteString("\n")
	}

	return builder.String()
}

func fixLineSize(line string, size int) string {
	padString := "%-" + strconv.Itoa(size) + "s"
	return fmt.Sprintf(padString, line)
}

func getLineType(lineIndex, numberOfLines int) int {
	switch lineIndex {
	case 0:
		return firstLine
	case numberOfLines - 1:
		return lastLine
	default:
		return normalLine
	}
}

func getLineDelimiters(lineType, numberOfLines int) [2]string {
	if numberOfLines == 1 {
		return [2]string{"< ", " >"}
	}

	switch lineType {
	case firstLine:
		return [2]string{"/ ", " \\"}
	case lastLine:
		return [2]string{"\\ ", " /"}
	default:
		return [2]string{"| ", " |"}
	}
}

func buildBaloonStart(baloonWidth int) string {
	return " " + strings.Repeat("_", baloonWidth+2) + "\n"
}

func buildBaloonEnd(baloonWidth int) string {
	return " " + strings.Repeat("-", baloonWidth+2) + "\n"
}

func buildCow() string {
	return "        \\   ^__^\n         \\  (oo)\\_______\n            (__)\\       )\\/\\\n                ||----w |\n                ||     ||\n"
}
