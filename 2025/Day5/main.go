package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	a, b int64
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	part1 := true

	pairs := []Pair{}
	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			part1 = false
			pairs = combine(pairs) // combine the intersecting ranges
			continue
		}

		if part1 {
			r := strings.Split(strings.TrimSpace(line), "-")
			n1, err1 := strconv.ParseInt(r[0], 10, 64)
			n2, err2 := strconv.ParseInt(r[1], 10, 64)
			if err1 != nil || err2 != nil {
				panic(fmt.Sprintln(err1, err2))
			}

			pairs = append(pairs, Pair{n1, n2})

		} else {
			num, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				panic(err)
			}

			if search(pairs, 0, len(pairs)-1, num) { // binary search, true if found
				count++
			}
		}
	}

	fmt.Println(count)
}

/*
Combine the intersecting ranges in the following manner:
  - Sort the pairs depending on starting value									 (curr)	   (new)
  - if ending value of current range is lower than starting value of new range like {3,5} & {10,14}
    just keep the new range
  - if the ending value is equal or larger like {10,14} & {12,18}, dont keep the new range, just change the
    ending value of current to bigger of the two
*/
func combine(pairs []Pair) []Pair {
	mergeSort(pairs, 0, len(pairs)-1)
	new := []Pair{pairs[0]}
	for i := 1; i < len(pairs); i++ {
		if pairs[i].a > new[len(new)-1].b {
			new = append(new, pairs[i])
		} else {
			new[len(new)-1].b = max(new[len(new)-1].b, pairs[i].b)
		}
	}
	return new
}

func search(pairs []Pair, low, high int, key int64) bool {
	if low > high {
		return false
	}

	mid := (low + high) / 2

	if pairs[mid].a <= key {
		if pairs[mid].b >= key {
			return true
		} else {
			return search(pairs, mid+1, high, key)
		}
	} else {
		return search(pairs, low, mid-1, key)
	}
}

func mergeSort(pairs []Pair, low, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSort(pairs, low, mid)
	mergeSort(pairs, mid+1, high)
	merge(pairs, low, mid, high)
}

func merge(pairs []Pair, low, mid, high int) {
	left := make([]Pair, mid-low+1)
	for i := range left {
		left[i] = pairs[low+i]
	}

	right := make([]Pair, high-mid)
	for i := range right {
		right[i] = pairs[mid+1+i]
	}

	a, b, i := 0, 0, low
	for (a < mid-low+1) && (b < high-mid) {
		if left[a].a < right[b].a {
			pairs[i] = left[a]
			a++
		} else {
			pairs[i] = right[b]
			b++
		}
		i++
	}

	for a < mid-low+1 {
		pairs[i] = left[a]
		a++
		i++
	}

	for b < high-mid {
		pairs[i] = right[b]
		b++
		i++
	}
}
