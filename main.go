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

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return
	}

	input := string(data)

	trimmedInput := strings.TrimSuffix(input, "\n")
	lines := strings.Split(trimmedInput, "\n")

	height := len(lines)
	width := len(lines[0])

	output := captureQuad(piscine.QuadA, width, height)

	if output == input {
		fmt.Println("QuadA matches!")
	} else {
		fmt.Println("QuadA does not match.")
	}

	fmt.Println("width: ", width)
	fmt.Println("Height: ", height)
}
