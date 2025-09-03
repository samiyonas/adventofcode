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

func increasing(list []int) bool {
	for i := 0; i < len(list) - 1; i++ {
		diff := list[i + 1] - list[i]
		if list[i] >= list[i + 1] || diff > 3 || diff < 1 {
			return false
		}

	}
	return true
}

func decreasing(list []int) bool {
	for i := 0; i < len(list) - 1; i++ {
		diff := list[i] - list[i + 1] 
		if list[i] <= list[i + 1] || diff > 3 || diff < 1 {
			return false
		}

	}
	return true
}

func isSafe(list []int) bool {
	if increasing(list) || decreasing(list) {
		return true
	}

	for i := 0; i < len(list); i++ {
		newList := append([]int{}, list[:i]...)
		newList = append(newList, list[i + 1:]...)

		if increasing(newList) || decreasing(newList) {
			return true
		}
	}

	return false
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

		if isSafe(s) {
			safe++
		}
	}

	fmt.Println(safe)
}
