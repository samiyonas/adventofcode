package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(file *os.File) int {
	scanner := bufio.NewScanner(file)
	curr := 50
	password := 0

	for scanner.Scan() {
		line := scanner.Text()
		turn, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("failed to type cast")
		}

		if line[0] == 'L' {
			curr = ((curr - turn + 100) % 100)
		} else {
			curr = (curr + turn) % 100
		}

		if curr == 0 {
			password += 1
		}

	}

	return password
}

func part2(file *os.File) int {
	scanner := bufio.NewScanner(file)
	curr := 50
	password := 0

	for scanner.Scan() {
		line := scanner.Text()
		turn, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal("type casting failed")
		}

		direction := line[0]
		for range turn {
			if direction == 'L' {
				curr = (curr - 1 + 100) % 100
			} else {
				curr = (curr + 1) % 100
			}

			if curr == 0 {
				password += 1
			}

		}
	}

	return password
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("couldn't open the file")
	}

	defer file.Close()

	fmt.Println(part2(file))
}
