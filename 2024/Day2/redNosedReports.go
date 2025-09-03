package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
	"strings"
)

func mapInt(list []string) []int {
	var result []int = make([]int, len(list))

	for i, v := range list {
		val, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("can't cast to int, err: %s", err.Error())
		}
		result[i] = val
	}

	return result
}

func increasing(list []int, n int) int {
	tolerate := 0
	for i := 1; i < n - 1; i++ {
		if tolerate > 1 {
			return 0
		}

		if list[i] >= list[i + 1] || list[i + 1] - list[i] < 1 || list[i + 1] - list[i] > 3 {
			tolerate += 1
		}
	}

	return 1
}

func decreasing(list []int, n int) int {
	tolerate := 0
	for i := 1; i < n - 1; i++ {
		if tolerate > 1 {
			return 0
		}
		if list[i] <= list[i + 1] || list[i] - list[i + 1] < 1 || list[i] - list[i + 1] > 3 {
			tolerate += 1
		}
	}

	return 1
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file, err: %s", err.Error())
	}

	safe := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s := mapInt(strings.Split(line, " "))

		if len(s) <= 1 {
			safe += 1
			continue
		}

		if s[0] > s[1] && (s[0] - s[1] >= 1 && s[0] - s[1] <= 3) {
			safe += decreasing(s, len(s))
		} else if s[0] < s[1] && (s[1] - s[0] >= 1 && s[1] - s[0] <= 3) {
			safe += increasing(s, len(s))
		}
	}

	fmt.Println(safe)
}
