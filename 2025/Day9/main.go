package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	a, b int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	coordinates := []coordinate{}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		s := strings.Split(line, ",")

		// first num is col
		b, err := strconv.Atoi(s[0])
		if err != nil {
			panic(err)
		}

		// second num is row
		a, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}

		coordinates = append(coordinates, coordinate{a: a, b: b})
	}

	ans := 0
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	area := func(i, j int) int {
		// Add 1 for index converting index to value
		dx := coordinates[i].a - coordinates[j].a + 1
		dy := coordinates[i].b - coordinates[j].b + 1
		return abs(dy * dx)
	}

	for i := 0; i < len(coordinates); i++ {
		for j := i + 1; j < len(coordinates); j++ {
			ans = max(ans, area(i, j))
		}
	}

	fmt.Println(ans)
}
