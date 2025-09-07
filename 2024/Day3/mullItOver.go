package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"bufio"
	"regexp"
)

func part1(filename string) int {
	file, err := os.Open(filename)

	defer file.Close()

	if err != nil {
		log.Fatalf("can't open file %s err: %s", filename, err.Error())
	}

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pattern := `mul\(([0-9]{1,3}),([0-9]{1,3})\)`
		re := regexp.MustCompile(pattern)

		matches := re.FindAllStringSubmatch(line, -1)

		for _, m := range matches {
			val1, _ := strconv.Atoi(m[1])
			val2, _ := strconv.Atoi(m[2])

			total += val1 * val2

		}
	}
	return total
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %s err: %s", filename, err.Error())
	}

	defer file.Close()

	pattern := `mul\(([0-9]{1,3}),([0-9]{1,3})\)|do\(\)|don't\(\)`

	total := 0
	scanner := bufio.NewScanner(file)
    enabled := true
	re := regexp.MustCompile(pattern)

	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllStringSubmatch(line, -1)

		for _, m := range matches {
            switch {
            case m[0] == "do()":
                enabled = true
            case m[0] == "don't()":
                enabled = false
            default:
                if enabled {
                    val1, _ := strconv.Atoi(m[1])
                    val2, _ := strconv.Atoi(m[2])

                    total += val1 * val2
				}
			}
		}
	}
	return total
}

func main() {
	filename := "input.txt"
	fmt.Println(part2(filename))
}
