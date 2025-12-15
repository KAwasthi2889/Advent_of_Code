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

	data := [][]byte{}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := append([]byte(nil), sc.Bytes()...)
		data = append(data, line)
	}

	var answer, ans int64
	op := true // true for +, false for *

	for i := range data[0] {
		space := true // false if a digit is encountered
		var num int64

		if operation := data[len(data)-1][i]; operation == '+' {
			op = true
		} else if operation == '*' {
			op = false
		}

		for j := 0; j < len(data)-1; j++ {
			if data[j][i] != ' ' {
				space = false
				n := int64(data[j][i] - '0')
				num = num*10 + n
			}
		}

		if !space {
			if op {
				ans += num
			} else {
				if ans == 0 {
					ans = 1
				}
				ans *= num
			}
		} else {
			fmt.Println(ans)
			answer += ans
			ans = 0
		}
	}

	fmt.Println(answer + ans)
}
