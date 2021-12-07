package main

import (
	"strconv"
	"testing"
)

func TestRecurseOxygen(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	int_input := convBinaryStrArrayToInts(input)

	got := recurseOxygen(int_input, 4)
	want := 23

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestRecurseCO2(t *testing.T) {
	input := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	var int_input []int
	for _, l := range input {
		n, _ := strconv.ParseInt(l, 2, 0)
		int_input = append(int_input, int(n))
	}

	got := recurseCO2(int_input, 4)
	want := 10

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
