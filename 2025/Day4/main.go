package main

import (
	"bufio"
	"fmt"
	"os"
)

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
	type Pairs struct {
		a, b int
	}
	var removables []Pairs // store removable indexes
	// Corners always have less than 4 neighbours

	// starting condition
	if rolls[0][0] {
		count++
	}
	removables = append(removables, Pairs{0, 0})

	for len(removables) > 0 {
		// remove every removables
		for _, pair := range removables {
			rolls[pair.a][pair.b] = false
		}

		// empty the removables slice
		removables = []Pairs{}

		for i, row := range rolls {
			for j, val := range row {
				// if the val is true and the neighbours are less than 4
				if val && check(i, j, rolls) {
					count++
					removables = append(removables, Pairs{i, j}) // add removables for next round
				}
			}
		}
	}

	fmt.Println(count)

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
