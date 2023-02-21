package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operators = []string{"+", "-", "*", "/"}

func solve(numbers []int, target int, operators []string) bool {
	if len(numbers) == 1 {
		return numbers[0] == target
	}

	for i := 0; i < len(operators); i++ {
		for j := 1; j < len(numbers); j++ {
			var newNumbers []int
			switch operators[i] {
			case "+":
				newNumbers = append(newNumbers, numbers[j-1]+numbers[j])
			case "-":
				newNumbers = append(newNumbers, numbers[j-1]-numbers[j])
			case "*":
				newNumbers = append(newNumbers, numbers[j-1]*numbers[j])
			case "/":
				if numbers[j] == 0 || numbers[j-1]%numbers[j] != 0 {
					continue
				}
				newNumbers = append(newNumbers, numbers[j-1]/numbers[j])
			}
			newNumbers = append(newNumbers, numbers[j+1:]...)
			if solve(newNumbers, target, operators) {
				fmt.Printf("%d %s %d = %d\n", numbers[j-1], operators[i], numbers[j], target)
				return true
			}
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := make([]int, 0)
		for _, s := range strings.Split(line, " ") {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error parsing input:", err)
				os.Exit(1)
			}
			numbers = append(numbers, num)
		}
		target := numbers[len(numbers)-1]
		numbers = numbers[:len(numbers)-1]

		if !solve(numbers, target, operators) {
			fmt.Println()
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
}


