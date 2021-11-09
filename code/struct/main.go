package main

import (
    "fmt"
)

func main() {

	s := make([]int, 0, 3)
	s = append(s, 10, 20)

	s1 := append(s, 30)
	s2 := append(s, 30, 40)

	PrintSlice(s)
	PrintSlice(s1)
	PrintSlice(s2)
}

func PrintSlice (s []int) {
	fmt.Printf("%v, cap %v, len %v, %p\n", s, cap(s), len(s), s)
}