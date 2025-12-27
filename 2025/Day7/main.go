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

	path := []string{}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		path = append(path, sc.Text())
	}

	type pair struct {
		a, b int
	}
	seen := make(map[pair]int)

	var travel func(i, j int) int
	travel = func(i, j int) int {
		if j < 0 || j >= len(path[0]) {
			return 0
		}

		if i >= len(path) {
			return 1
		}

		if res := seen[pair{i, j}]; res != 0 {
			return res
		}

		if path[i][j] == '^' {
			left := travel(i, j-1)
			right := travel(i, j+1)
			seen[pair{i, j}] = left + right
			return left + right
		} else {
			res := travel(i+1, j)
			seen[pair{i, j}] = res
			return res
		}
	}

	for j, char := range path[0] {
		if char == 'S' {
			fmt.Println(travel(1, j))
			break
		}
	}
}
