package main

import "fmt"

func main() {
	/* len=0, cap=0のslice */
	s1 := []int{}
	for i := 0; i < 10000; i++ {
		s1 = append(s1, i)
	}
	fmt.Printf("1:cap=%d\n", cap(s1))

	/* len=0, cap=10000のslice */
	s2 := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		s2 = append(s2, i)
	}
	fmt.Printf("2:cap=%d\n", cap(s2))
}
