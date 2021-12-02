package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput() []uint {

	file, err := os.Open("day1/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var result []uint

	for scanner.Scan() {
		line := scanner.Text()
		tmp, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("failed to parse uint from string")
		}
		result = append(result, uint(tmp))
	}

	file.Close()

	return result
}

func windowSum(input []uint) uint {
	// Why is there no builtin for this
	var result uint = 0
	for _, v := range input {
		result += v
	}
	return result
}

func formWindows(input []uint) []uint {
	// convert array of uints to a sliding window array
	const windowSize = 3
	windowCount := len(input) - windowSize + 1
	var result []uint

	for index := 0; index < windowCount; index++ {
		window := input[index : index+windowSize]
		result = append(result, windowSum(window))
	}
	return result
}

func main() {
	depths := readInput()
	windows := formWindows(depths)

	for i := range windows {
		fmt.Println(depths[i], windows[i])
	}
	fmt.Println(len(depths), len(windows))

	var increses uint = 0
	var previous uint = ^uint(0)

	for _, current := range windows {

		if current > previous {
			increses++
		}
		previous = current
	}

	fmt.Println(increses)

}
