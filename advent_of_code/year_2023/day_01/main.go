package main

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/billwallis/advent-of-code-go/advent_of_code/utils"
)

//go:embed input.data
var input string

func main() {
	fmt.Println("--- year 2023 day 01 ---")
	data := strings.TrimSuffix(input, "\n")
	fmt.Println("    part 1:", part1(data))
	fmt.Println("    part 2:", part2(data))
}

func part1(data string) any {
	sum := 0
	for _, line := range strings.Split(data, "\n") {
		var left, right string
		for i := 0; i < len(line); i++ {
			if left = string(line[i]); strings.ContainsAny(left, "0123456789") {
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if right = string(line[i]); strings.ContainsAny(right, "0123456789") {
				break
			}
		}

		sum += utils.ToInt(left + right)
	}

	return sum
}

func part2(data string) any {
	var digits = map[string]string{
		"0":     "0",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	sum := 0
	for _, line := range strings.Split(data, "\n") {
		var left, right string
		var startsWith bool

		for i := 0; i < len(line); i++ {
			if left, startsWith = stringStartsWithDigit(line[i:], digits); startsWith {
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if right, startsWith = stringStartsWithDigit(line[i:], digits); startsWith {
				break
			}
		}

		sum += utils.ToInt(left + right)
	}

	return sum
}

func stringStartsWithDigit(s string, digits map[string]string) (string, bool) {
	for word, digit := range digits {
		if stringHasPrefix(s, word) {
			return digit, true
		}
	}
	return "", false
}

func stringHasPrefix(s string, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	return s[:len(prefix)] == prefix
}
