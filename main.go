package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func findSteps(m map[int]string, posX int, posY int, steps int, dir rune) (bool, int) {
	line, exists := m[posY]
	if !exists {
		return false, 0
	}
	if len(line) < posX {
		return false, 0
	}
	pointVal := line[posX]
	steps++
	switch pointVal {
	case 'S':
		if dir == 'B' {
			atb, atv := findSteps(m, posX, posY+1, steps, 'N')
			if atb {
				return atb, atv
			}
			atb, atv = findSteps(m, posX, posY-1, steps, 'S')
			if atb {
				return atb, atv
			}
			return findSteps(m, posX-1, posY, steps, 'E')
		}
		return true, steps
	case '.':
		return false, 0
	case '|':
		if dir == 'N' {
			return findSteps(m, posX, posY+1, steps, 'N')
		} else if dir == 'S' {
			return findSteps(m, posX, posY-1, steps, 'S')

		} else {
			return false, 0
		}
	case '-':
		if dir == 'E' {
			return findSteps(m, posX-1, posY, steps, 'E')
		} else if dir == 'W' {
			return findSteps(m, posX+1, posY, steps, 'W')
		} else {
			return false, 0
		}
	case 'L':
		if dir == 'N' {
			return findSteps(m, posX+1, posY, steps, 'W')
		} else if dir == 'E' {
			return findSteps(m, posX, posY-1, steps, 'S')
		} else {
			return false, 0
		}

	case 'J':
		if dir == 'N' {
			return findSteps(m, posX-1, posY, steps, 'E')
		} else if dir == 'W' {
			return findSteps(m, posX, posY-1, steps, 'S')
		} else {
			return false, 0
		}
	case '7':
		if dir == 'S' {
			return findSteps(m, posX-1, posY, steps, 'E')
		} else if dir == 'W' {
			return findSteps(m, posX, posY+1, steps, 'N')
		} else {
			return false, 0
		}
	case 'F':
		if dir == 'S' {
			return findSteps(m, posX+1, posY, steps, 'W')
		} else if dir == 'E' {
			return findSteps(m, posX, posY+1, steps, 'N')
		} else {
			return false, 0
		}
	}
	return false, 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := make(map[int]string)
	pos := 0
	searching := true
	var posX int
	var posY int
	for scanner.Scan() {
		text := scanner.Text()
		m[pos] = text
		if searching {
			index := strings.Index(text, "S")
			if index > -1 {
				posX = index
				posY = pos
			}

		}
		pos++
	}
	isTrue, val := findSteps(m, posX, posY, 0, 'B')
	if isTrue {
		fmt.Println((val - 1) / 2)
	}
}
