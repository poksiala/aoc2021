package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

type line struct {
	start point
	end   point
}

func readInput() []line {
	file, err := os.Open("day5/input.txt")

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines []line
	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " -> ")
		start_split := strings.Split(split[0], ",")
		end_split := strings.Split(split[1], ",")
		start_x, err := strconv.Atoi(start_split[0])
		if err != nil {
			fmt.Println(err)
		}
		start_y, err := strconv.Atoi(start_split[1])
		if err != nil {
			fmt.Println(err)
		}
		end_x, err := strconv.Atoi(end_split[0])
		if err != nil {
			fmt.Println(err)
		}
		end_y, err := strconv.Atoi(end_split[1])
		if err != nil {
			fmt.Println(err)
		}
		lines = append(lines, line{point{start_x, start_y}, point{end_x, end_y}})
	}

	file.Close()

	return lines
}

func numRange(start int, end int) []int {
	// inclusive range between start and end
	delta := end - start
	absDelta := int(math.Abs(float64(delta)))
	if absDelta == 0 {
		return []int{start}
	}
	dir := delta / absDelta
	var result []int
	for i := 0; i <= absDelta; i++ {
		result = append(result, start+i*dir)
	}
	return result
}

func lineToPoints(l line) []point {
	// return points along the line. Inclusive
	var points []point
	x_range := numRange(l.start.x, l.end.x)
	y_range := numRange(l.start.y, l.end.y)

	if len(x_range) == len(y_range) {
		for i := range x_range {
			points = append(points, point{x_range[i], y_range[i]})
		}
	} else {
		for _, x := range x_range {
			for _, y := range y_range {
				points = append(points, point{x, y})
			}
		}
	}
	return points

}

func countHighpoints(lines []line) int {
	grid := make(map[point]int)
	for _, l := range lines {
		points := lineToPoints(l)
		for _, p := range points {
			grid[p] += 1
		}
	}
	res := 0
	for _, value := range grid {
		if value > 1 {
			res++
		}
	}
	return res
}

func pt1(input []line) int {

	// filter out diagonal lines
	var lines []line
	for _, l := range input {
		if l.start.x == l.end.x || l.start.y == l.end.y {
			lines = append(lines, l)
		}
	}
	return countHighpoints(lines)
}

func pt2(input []line) int {
	return countHighpoints(input)
}

func main() {
	lines := readInput()
	fmt.Println(pt1(lines))
	fmt.Println(pt2(lines))
}
