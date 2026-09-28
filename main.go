package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func getDimensions(input string) (int, int, bool) {
	if input == "" {
		return 0, 0, false
	}

	input = strings.TrimSuffix(input, "\n")

	if input == "" {
		return 0, 0, false
	}

	lines := strings.Split(input, "\n")

	width := len([]rune(lines[0]))
	height := len(lines)

	if width == 0 {
		return 0, 0, false
	}

        // Every row must have the same width.
	for _, line := range lines {
		if len([]rune(line)) != width {
			return 0, 0, false
		}
	}

	return width, height, true
}

func expectedChar(quad string, x, y, width, height int) rune {
	switch quad {
	case "quadA":
		if y == 0 || y == height-1 {
			if x == 0 || x == width-1 {
				return 'o'
			}
			return '-'
		}

		if x == 0 || x == width-1 {
			return '|'
		}

		return ' '

	case "quadB":
		if y == 0 && x == 0 {
			return '/'
		}
		if y == 0 && x == width-1 {
			return '\\'
		}
		if y == height-1 && x == 0 {
			return '\\'
		}
		if y == height-1 && x == width-1 {
			return '/'
		}
		if y == 0 || y == height-1 || x == 0 || x == width-1 {
			return '*'
		}

		return ' '

	case "quadC":
		if y == 0 && (x == 0 || x == width-1) {
			return 'A'
		}
		if y == height-1 && (x == 0 || x == width-1) {
			return 'C'
		}
		if y == 0 || y == height-1 || x == 0 || x == width-1 {
			return 'B'
		}

		return ' '

	case "quadD":
		if (y == 0 || y == height-1) && x == 0 {
			return 'A'
		}
		if (y == 0 || y == height-1) && x == width-1 {
			return 'C'
		}
		if y == 0 || y == height-1 || x == 0 || x == width-1 {
			return 'B'
		}

		return ' '

	case "quadE":
		if y == 0 && x == 0 {
			return 'A'
		}
		if y == 0 && x == width-1 {
			return 'C'
		}
		if y == height-1 && x == 0 {
			return 'C'
		}
		if y == height-1 && x == width-1 {
			return 'A'
		}
		if y == 0 || y == height-1 || x == 0 || x == width-1 {
			return 'B'
		}

		return ' '
	}

	return 0
}

func checkQuad(input string, quad string, width, height int) bool {
	input = strings.TrimSuffix(input, "\n")
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		runes := []rune(line)

		for x, char := range runes {
			if char != expectedChar(quad, x, y, width, height) {
				return false
			}
		}
	}

	return true
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("Not a quad function")
		return
	}

	input := string(data)

	width, height, valid := getDimensions(input)

	if !valid {
		fmt.Println("Not a quad function")
		return
	}

	quads := []string{
		"quadA",
		"quadB",
		"quadC",
		"quadD",
		"quadE",
	}

	matches := []string{}

	for _, quad := range quads {
		if checkQuad(input, quad, width, height) {
			matches = append(
				matches,
				fmt.Sprintf("[%s] [%d] [%d]", quad, width, height),
			)
		}
	}

	if len(matches) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	fmt.Println(strings.Join(matches, " || "))
}
