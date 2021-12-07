package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput() []string {

	file, err := os.Open("day3/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var result []string

	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}

	file.Close()

	return result
}

func pt1(input []string) int {
	count := len(input)
	halfCount := count / 2
	bits := len(input[0])
	var popularity []int

	// initialize popularity array
	for i := 0; i < bits; i++ {
		popularity = append(popularity, 0)
	}

	// calculate bit popularity line by line and bit by bit
	for _, line := range input {
		num, err := strconv.ParseInt(line, 2, 64)
		if err != nil {
			fmt.Println(err)
		}
		for i := range popularity {
			popularity[i] += int((num >> i) & 1)
		}
	}

	// Calculate gamma and epsilon
	var gamma, epsilon int
	for i, n := range popularity {
		if n > halfCount {
			gamma += 1 << i
		} else {
			epsilon += 1 << i
		}
	}

	return gamma * epsilon
}

func convBinaryStrArrayToInts(input []string) []int {
	var result []int
	for _, l := range input {
		n, _ := strconv.ParseInt(l, 2, 0)
		result = append(result, int(n))
	}
	return result
}

func returnLongerOrFirst(a []int, b []int) []int {
	if len(a) >= len(b) {
		return a
	} else {
		return b
	}
}

func returnShorterOrFirst(a []int, b []int) []int {
	if len(a) <= len(b) {
		return a
	} else {
		return b
	}
}

func divideBasedOnBit(input []int, n int) ([]int, []int) {
	// divide input into two arrays (ones, zeroes) based on bit on nth position
	var ones, zeroes []int
	filter := 1 << n
	for _, line := range input {
		if line&filter == 0 {
			zeroes = append(zeroes, line)
		} else {
			ones = append(ones, line)
		}
	}
	return ones, zeroes
}

func recurseOxygen(input []int, n int) int {
	ones, zeroes := divideBasedOnBit(input, n)
	longer := returnLongerOrFirst(ones, zeroes)
	if len(longer) == 1 {
		return longer[0]
	} else {
		return recurseOxygen(longer, n-1)
	}
}

func recurseCO2(input []int, n int) int {
	ones, zeroes := divideBasedOnBit(input, n)
	shorter := returnShorterOrFirst(zeroes, ones)
	if len(shorter) == 1 {
		return shorter[0]
	} else {
		return recurseCO2(shorter, n-1)
	}
}

func pt2(input []string) int {
	bits := len(input[0])
	int_input := convBinaryStrArrayToInts(input)
	oxygen := recurseOxygen(int_input, bits-1)
	co2 := recurseCO2(int_input, bits-1)
	return oxygen * co2
}

func main() {
	input := readInput()
	fmt.Println(pt1(input))
	fmt.Println(pt2(input))
}
