package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() ([]int, [][]int) {
	/*
		Returns:
			- "Random" numbers as []int
			- Array of boards as one dimensional arrays. [][]int

	*/

	file, err := os.Open("day4/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	numberLine := scanner.Text()
	split := strings.Split(numberLine, ",")
	var numbers []int
	for _, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		numbers = append(numbers, n)
	}

	var boards [][]int
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			var new []int
			boards = append(boards, new)
		} else {
			split := strings.Fields(line)
			for _, s := range split {
				n, err := strconv.Atoi(s)
				if err != nil {
					fmt.Println(err)
				}
				boards[len(boards)-1] = append(boards[len(boards)-1], n)
			}
		}
	}

	file.Close()

	return numbers, boards
}

func sumNonNegatives(arr []int) int {
	sum := 0
	for _, n := range arr {
		if n > 0 {
			sum += n
		}
	}
	return sum
}

func hasBingo(board []int) bool {
	// return true if there is a row or column which has a product of -1
	for i := 0; i < 5; i++ {
		var a, b = 1, 1
		for j := 0; j < 5; j++ {
			a = a * board[i*5+j]
			b = b * board[i+j*5]
		}
		if a == -1 || b == -1 {
			return true
		}
	}
	return false
}

func pt1(numbers []int, boards [][]int) int {
	for _, number := range numbers {
		for _, board := range boards {

			// replace matching digit with -1
			for i, v := range board {
				if v == number {
					board[i] = -1
				}
			}

			if hasBingo(board[:]) {
				sum := sumNonNegatives(board[:])
				return sum * number
			}

		}
	}
	return -1
}

func pt2(numbers []int, boards [][]int) int {
	for i, number := range numbers {
		for _, board := range boards {

			// replace matching digit with -1
			for i, v := range board {
				if v == number {
					board[i] = -1
				}
			}
		}

		// remove boards containing bingo
		index := 0
		for _, board := range boards {
			if !hasBingo(board[:]) {
				boards[index] = board
				index++
			}
		}
		boards = boards[:index]

		if index == 1 {
			return pt1(numbers[i:], boards[:])
		}
	}
	return -1
}

func main() {
	numbers, boards := readInput()
	fmt.Println(pt1(numbers, boards))
	fmt.Println(pt2(numbers, boards))
}
