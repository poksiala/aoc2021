package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() [9]int {
	file, err := os.Open("day6/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := scanner.Text()
	file.Close()
	split := strings.Split(line, ",")
	timers := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	for _, s := range split {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		timers[n] += 1
	}

	file.Close()

	return timers
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

func pt1(timers [9]int, days int) int {
	for d := 0; d < days; d++ {
		zeroes := timers[0]
		for i := 0; i < 9; i++ {
			if i == 8 {
				timers[i] = zeroes
			} else if i == 6 {
				timers[i] = timers[i+1] + zeroes
			} else {
				timers[i] = timers[i+1]
			}
		}
	}
	return sumNonNegatives(timers[:])
}

func main() {
	timers := readInput()
	fmt.Println(pt1(timers, 80))
	fmt.Println(pt1(timers, 256))
}
