package cowsay

import (
	"strconv"
	"strings"
	"fmt"
)

const LINE_LENGTH_LIMIT = 39
const SPACE = " "

const (
	FIRST_LINE = iota
	NORMAL_LINE
	LAST_LINE
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
		finalIndex := startIndex + LINE_LENGTH_LIMIT
		if finalIndex > stringSize {
			finalIndex = stringSize
		}

		lastSpaceIndex := strings.LastIndex(s[startIndex : finalIndex], SPACE)
		if lastSpaceIndex != -1 && (finalIndex - startIndex) == LINE_LENGTH_LIMIT {
			finalIndex = startIndex + lastSpaceIndex + 1
		}
		
		line := s[startIndex : finalIndex]
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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func buildBaloonText(lines []string, biggestLineSize int) string {
	var baloonText string

	numberOfLines := len(lines)

	for lineIndex, line := range lines {
		line := fixLineSize(line, biggestLineSize)
		lineType := getLineType(lineIndex, numberOfLines)
		lineDelimiters := getLineDelimiters(lineType, numberOfLines)

		baloonText += lineDelimiters[0] + line + lineDelimiters[1] + "\n"
	}

	return baloonText
}

func fixLineSize(line string, size int) string {
	padString := "%-" + strconv.Itoa(size) + "s"
	return fmt.Sprintf(padString, line)
}

func getLineType(lineIndex, numberOfLines int) int {
	switch lineIndex {
	case 0:
		return FIRST_LINE
	case numberOfLines - 1:
		return LAST_LINE
	default:
		return NORMAL_LINE
	}
}

func getLineDelimiters(lineType, numberOfLines int) [2]string {
	if numberOfLines == 1 {
		return [2]string{"< ", " >"}
	}

	switch lineType {
	case FIRST_LINE:
		return [2]string{"/ ", " \\"}
	case LAST_LINE:
		return [2]string{"\\ ", " /"}
	default:
		return [2]string{"| ", " |"}
	}
}

func buildBaloonStart(baloonWidth int) string {
	return " " + strings.Repeat("_", baloonWidth + 2) + "\n"
}

func buildBaloonEnd(baloonWidth int) string {
	return " " + strings.Repeat("-", baloonWidth + 2) + "\n"
}

func buildCow() string {
	return "        \\   ^__^\n         \\  (oo)\\_______\n            (__)\\       )\\/\\\n                ||----w |\n                ||     ||\n"
}
