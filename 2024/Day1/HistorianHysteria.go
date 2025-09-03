package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
	"log"
	"sort"
)

func part1(file *os.File) int {
	left_location := make([]int, 0)
	right_location := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, "   ")
		left, err := strconv.Atoi(splitted[0])
		if err != nil {
			log.Fatalf("can't cast to int, err: %s", err.Error()) 
		}
		right, err := strconv.Atoi(splitted[1])
		if err != nil {
			log.Fatalf("can't cast to int, err: %s", err.Error()) 
		}

		left_location = append(left_location, left)
		right_location = append(right_location, right)
	}

	sort.Ints(left_location)
	sort.Ints(right_location)

	var total float64 = 0

	for i := 0; i < len(left_location); i++ {
		total += math.Abs(float64(left_location[i] - right_location[i]))
	}

	return int(total)
}

func part2(file *os.File) int {
	left_location := make([]int, 0)
	right_location := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, "   ")
		left, err := strconv.Atoi(splitted[0])
		if err != nil {
			log.Fatalf("can't cast to int, err: %s", err.Error()) 
		}
		right, err := strconv.Atoi(splitted[1])
		if err != nil {
			log.Fatalf("can't cast to int, err: %s", err.Error()) 
		}

		left_location = append(left_location, left)
		right_location = append(right_location, right)
	}

	rightMap := make(map[int]int, 0)
	total := 0

	for _, val := range right_location {
		rightMap[val] += 1
	}
	
	for _, val := range left_location {
		total += val * rightMap[val]
	}

	return total
}
func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file: %s: err", filename, err.Error())
	}
	fmt.Println(part2(file))
}
