package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"piscine"
)

func captureQuad(quad func(int, int), x, y int) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		return ""
	}

	oldStdout := os.Stdout
	os.Stdout = writer

	quad(x, y)

	writer.Close()
	os.Stdout = oldStdout

	data, err := io.ReadAll(reader)
	if err != nil {
		return ""
	}

	return string(data)
}

func checkQuad(input string, widht, height int, quad func(int, int)) bool {
	output := captureQuad(quad, widht, height)

	return output == input
}

func getDimensions(input string) (int, int, bool) {
	trimmedInput := strings.TrimSuffix(input, "\n")

	if trimmedInput == "" {
		return 0, 0, false
	}

	lines := strings.Split(trimmedInput, "\n")

	width := len(lines[0])
	height := len(lines)

	for _, line := range lines {
		if len(line) != width {
			return 0, 0, false
		}
	}

	return width, height, true
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	input := string(data)

	width, height, valid := getDimensions(input)

	if !valid {
		fmt.Println("Not a quad function")
		return
	}

	matches := []string{}

	if checkQuad(input, width, height, piscine.QuadA) {
		matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", width, height))
	}

	if checkQuad(input, width, height, piscine.QuadB) {
		matches = append(matches, fmt.Sprintf("[quadB] [%d] [%d]", width, height))
	}

	if checkQuad(input, width, height, piscine.QuadC) {
		matches = append(matches, fmt.Sprintf("[quadC] [%d] [%d]", width, height))
	}

	if checkQuad(input, width, height, piscine.QuadD) {
		matches = append(matches, fmt.Sprintf("[quadD] [%d] [%d]", width, height))
	}

	if checkQuad(input, width, height, piscine.QuadE) {
		matches = append(matches, fmt.Sprintf("[quadE] [%d] [%d]", width, height))
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	fmt.Println(strings.Join(matches, " || "))
}
