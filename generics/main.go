package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	ints := map[string]int64{
		"a": 1,
		"b": 2,
	}

	floats := map[string]float64{
		"a": 1.1,
		"b": 1.2,
	}

	totalInts := sumInts(ints)
	totalFloats := sumFloats(floats)

	fmt.Printf(" %v %v", totalInts, totalFloats)
	fmt.Printf(" %v %v", sumIntsOrFloats[string, int64](ints), sumIntsOrFloats[string, float64](floats))
	fmt.Printf(" %v %v", sumNumbers[string, int64](ints), sumNumbers[string, float64](floats))
}

func sumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s

}

func sumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}

	return s
}

func sumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}

func sumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}

	return s
}
