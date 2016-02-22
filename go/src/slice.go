package main

import "fmt"

func main() {
	/* [...]で要素数を省略した配列型。[5]intと書いても同じ */
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Printf("length=%d,capacity=%d,first=%d\n", len(a), cap(a), a[0])
	/* [5]int全体を参照するslice */
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("length=%d,capacity=%d,first=%d\n", len(s), cap(s), s[0])
	/* [5]intの2要素目から3要素目までの参照するslice */
	p := a[1:3]
	fmt.Printf("length=%d,capacity=%d,first=%d\n", len(p), cap(p), p[0])
	/* capacity=100で先頭の１要素のみ0に初期化されたslice */
	b := make([]int, 1, 100)
	fmt.Printf("length=%d,capacity=%d,first=%d\n", len(b), cap(b), b[0])
}
