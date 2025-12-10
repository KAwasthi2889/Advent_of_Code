package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	position = 50
	count    = 0
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error in Input File")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		num, _ := strconv.Atoi(input[1:])
		count += num / 100
		num %= 100

		if input[0] == 'L' {
			num = -num
		}

		if num != 0 {
			position += num
			if position <= 0 {
				if position != num {
					count++
				}
				if position != 0 {
					position += 100
				}
			} else if position >= 100 {
				count++
				position -= 100
			}
		}
		// fmt.Println(position, count)
	}

	fmt.Println(count)
}
