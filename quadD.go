package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func expectedChar(quad string, x, y, width, height int) rune {
	switch quad {
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
	}
}
