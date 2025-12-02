package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func part1(file *os.File) int64 {
	var sum int64

	scanner := bufio.NewScanner(file)
	var ranges []string

	for scanner.Scan() {
		ranges = strings.Split(scanner.Text(), ",")
	}

	for _, ra := range ranges {
		arr := strings.Split(ra, "-")
		start, _ := strconv.Atoi(arr[0])
		end, _ := strconv.Atoi(arr[1])

		for i := int64(start); i <= int64(end); i++ {
			val := strconv.Itoa(int(i))
			if len(val)%2 == 0 {
				if val[:len(val)/2] == val[len(val)/2:] {
					sum += i
				}
			}
		}
	}

	return sum
}

func part2(file *os.File) int64 {
	var sum int64

	scanner := bufio.NewScanner(file)
	var ranges []string

	for scanner.Scan() {
		ranges = strings.Split(scanner.Text(), ",")
	}

	for _, ra := range ranges {
		arr := strings.Split(ra, "-")
		start, _ := strconv.Atoi(arr[0])
		end, _ := strconv.Atoi(arr[1])

		for i := int64(start); i <= int64(end); i++ {
			val := strconv.Itoa(int(i))
			valen := int64(len(val))

			for j := int64(1); j <= valen/2; j += 1 {
				if valen%j != 0 {
					continue
				}
				block := val[:j]
				valid := true

				for k := j; k < valen; k += j {
					if val[k:k+j] != block {
						valid = false
						break
					}
				}
				if valid {
					sum += i
					break
				}
			}
		}
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("couldn't open input.txt")
	}

	log.Println(part2(file))
}
