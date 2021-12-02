package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type move struct {
	direction string
	amount    int
}

func readInput() []move {

	file, err := os.Open("day2/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var result []move

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		direction := split[0]
		amount, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("failed to parse depth from string")
		}
		result = append(result, move{direction: direction, amount: amount})
	}

	file.Close()

	return result
}

func pt1(moves []move) int {
	var depth int = 0
	var distance int = 0

	for _, m := range moves {
		if m.direction == "forward" {
			distance += m.amount
		} else if m.direction == "up" {
			depth -= m.amount
		} else {
			depth += m.amount
		}
	}

	return depth * distance
}

func pt2(moves []move) int {
	var depth int = 0
	var distance int = 0
	var aim int = 0

	for _, m := range moves {
		if m.direction == "forward" {
			distance += m.amount
			depth += m.amount * aim
		} else if m.direction == "up" {
			aim -= m.amount
		} else {
			aim += m.amount
		}
	}

	return depth * distance
}

func main() {
	moves := readInput()
	fmt.Println(pt1(moves))
	fmt.Println(pt2(moves))
}
