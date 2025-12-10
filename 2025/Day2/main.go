package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type pair struct{ a, b int64 }

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Split(strings.TrimSpace(string(data)), ",")
	ranges := make([]pair, len(parts))

	for i, p := range parts {
		str := strings.Split(p, "-")
		a, _ := strconv.ParseInt(str[0], 10, 64)
		b, _ := strconv.ParseInt(str[1], 10, 64)
		ranges[i] = pair{a, b}
	}

	var sum int64

	for _, r := range ranges {
		for num := r.a; num <= r.b; num++ {
			if check(num) {
				sum += num
			}
		}
	}

	fmt.Println(sum)
}

func check(num int64) bool {
	str := strconv.FormatInt(num, 10)
	n := len(str)

	for i := 1; i <= n/2; i++ {
		if n%i != 0 {
			continue
		}

		pattern := str[:i]
		if strings.Repeat(pattern, n/i) == str {
			return true
		}
	}
	return false
}
