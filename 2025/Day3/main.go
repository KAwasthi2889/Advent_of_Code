package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	CHOOSE = 12
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var jolts int64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		skips := len(line) - CHOOSE // we can skip this number of times
		stack := make([]byte, 0, CHOOSE)

		for i := range line {
			char := line[i] // char is current
			for skips > 0 && len(stack) > 0 && stack[len(stack)-1] < char {
				pop(&stack) // pop is skips are available, stack is non empty & curr > top
				skips--     // reduce skips on each pop
			}
			if len(stack) == 0 || skips == 0 || stack[len(stack)-1] >= char {
				// push if stack is empty, no skips are available or curr >= top
				push(char, &stack, &skips) // push ignores the input if stack is full
			}
		}

		// Create a number with stack
		var sum int64
		for _, val := range stack {
			num := int(val - '0')     // convert byte to int
			sum = sum*10 + int64(num) // make the number
		}

		fmt.Println(sum)
		jolts += sum
	}

	fmt.Println(jolts)
}

func pop(stack_ptr *[]byte) { // pop the topmost element
	top := len(*stack_ptr) - 1
	*stack_ptr = (*stack_ptr)[:top]
}

func push(char byte, stack_ptr *[]byte, skips *int) { // ignores char if already filled
	if cap(*stack_ptr) > len(*stack_ptr) {
		*stack_ptr = append(*stack_ptr, char)
	} else {
		*skips--
	}
}
