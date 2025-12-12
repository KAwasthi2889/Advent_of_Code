package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pairs struct {
	a, b int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var rolls [][]bool

	for scanner.Scan() {
		line := scanner.Text()

		roll := make([]bool, len(line))

		for i, r := range line {
			if r == '@' {
				roll[i] = true
			}
		}

		rolls = append(rolls, roll)

	}

	count := 0
	var removables []Pairs // store removable indexes, neighboures of removed rolls

	for i, row := range rolls {
		for j, val := range row {
			// if the val is true and the neighbours are less than 4
			if val && check(i, j, rolls) {
				removables = append(removables, Pairs{i, j}) // add removables for next round
			}
		}
	}

	for len(removables) > 0 {
		// remove rolls
		for _, r := range removables {
			rolls[r.a][r.b] = false
			count++
		}

		remaining := make(map[Pairs]bool)

		for _, r := range removables {
			find(r.a, r.b, rolls, &remaining) // find remaining neighbours of removed, check them and add removables
		}

		removables = nil

		for p := range remaining {
			removables = append(removables, p)
		}

	}

	fmt.Println(count)

}

func find(i, j int, rolls [][]bool, remaining *map[Pairs]bool) {
	directions := [8][2]int{
		{1, 0},   // up
		{-1, 0},  // down
		{0, -1},  // left
		{0, 1},   // right
		{1, -1},  // up left
		{1, 1},   // up right
		{-1, -1}, // down left
		{-1, 1},  // down right
	}

	for _, direction := range directions {
		a, b := i+direction[0], j+direction[1]
		if a >= 0 && a < len(rolls) && b >= 0 && b < len(rolls[0]) {
			if rolls[a][b] && check(a, b, rolls) {
				(*remaining)[Pairs{a, b}] = true
			}
		}
	}
}

func check(i, j int, rolls [][]bool) bool {
	neighbours := 0

	directions := [8][2]int{
		{1, 0},   // up
		{-1, 0},  // down
		{0, -1},  // left
		{0, 1},   // right
		{1, -1},  // up left
		{1, 1},   // up right
		{-1, -1}, // down left
		{-1, 1},  // down right
	}

	for _, direction := range directions {
		a, b := i+direction[0], j+direction[1]
		if a >= 0 && a < len(rolls) && b >= 0 && b < len(rolls[0]) {
			if rolls[a][b] {
				neighbours++
			}
		}
	}

	return neighbours < 4
}
