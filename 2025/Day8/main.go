package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	type coordinate struct {
		x, y, z int
	}
	coordinates := []coordinate{}
	dist := func(a, b coordinate) float64 {
		dx := a.x - b.x
		dy := a.y - b.y
		dz := a.z - b.z
		return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()
		strs := strings.Split(line, ",")
		nums := make([]int, len(strs))
		for i, s := range strs {
			if nums[i], err = strconv.Atoi(s); err != nil {
				panic(err)
			}
		}
		coordinates = append(coordinates, coordinate{nums[0], nums[1], nums[2]})
	}

	// Calculate all the dist and put it in a min heap
	h := heap{}
	for i := range coordinates {
		for j := i + 1; j < len(coordinates); j++ {
			h.Add(distance{i, j, dist(coordinates[i], coordinates[j])})
		}
	}

	var dsu Disjoint
	dsu.parent = make([]int, len(coordinates))
	for i := range dsu.parent {
		// All circuits have themselves as root
		dsu.parent[i] = i
	}

	dsu.size = make([]int, len(coordinates))
	for i := range dsu.size {
		// All circuits have size atleast 1
		dsu.size[i] = 1
	}

	last_i, last_j := 0, 0

	for len(h) > 0 {
		i, j := h.Pop()
		// If they have same root, they are in same circuit. If not, mearge them
		if ri, rj := dsu.root(i), dsu.root(j); ri != rj {
			last_i, last_j = coordinates[i].x, coordinates[j].x
			if dsu.size[ri] >= dsu.size[rj] {
				// root of j will have root of i as parent
				dsu.parent[rj] = ri
				// size of root of i will grow by the size of root of j
				dsu.size[ri] += dsu.size[rj]
			} else {
				dsu.parent[ri] = rj
				dsu.size[rj] += dsu.size[ri]
			}
		}
	}

	fmt.Println(last_i * last_j)
}

type Disjoint struct {
	parent []int // Who is the root of this circuit
	size   []int // What is the size of this circuit
}

// Find the root/parent of given index
func (dsu Disjoint) root(i int) int {
	seen := map[int]bool{}
	seen[i] = true
	root := dsu.parent[i]
	for !seen[root] {
		seen[root] = true
		root = dsu.parent[root]
	}
	return root
}

type distance struct {
	i, j int
	d    float64
}

type heap []distance

func (h *heap) Add(x distance) {
	*h = append(*h, x)

	i := len(*h) - 1
	for i > 0 {
		p := (i - 1) / 2
		if (*h)[i].d >= (*h)[p].d {
			break
		}
		(*h)[i], (*h)[p] = (*h)[p], (*h)[i]
		i = p
	}
}

func (h *heap) Pop() (int, int) {
	n := len(*h)
	if n == 0 {
		panic("Pop from empty heap")
	}

	root := (*h)[0]

	// Move last element to root and shrink
	last := (*h)[n-1]
	*h = (*h)[:n-1]

	if len(*h) == 0 {
		return root.i, root.j
	}

	(*h)[0] = last

	// Heapify down
	i := 0
	for {
		left := 2*i + 1
		right := 2*i + 2
		smallest := i

		if left < len(*h) && (*h)[left].d < (*h)[smallest].d {
			smallest = left
		}
		if right < len(*h) && (*h)[right].d < (*h)[smallest].d {
			smallest = right
		}
		if smallest == i {
			break
		}

		(*h)[i], (*h)[smallest] = (*h)[smallest], (*h)[i]
		i = smallest
	}
	return root.i, root.j
}
