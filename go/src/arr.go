package main

import "fmt"

type Math struct {
	i int
}

func main() {
	// 配列を引数にとってソートする関数をもったstructを作成

	arr := [...]int{10, 2, 5, 7, 79, 34, 609, 3, 20, 1}
	fmt.Println(arr)
	x := Math{1}
	x.sort(arr[:])
	fmt.Println(arr)
}

func (m Math) sort(arr []int) {
	var tmp int
	for i := 0; len(arr)-1 > i; i++ {
		for j := len(arr) - 1; j > i; j-- {

			if arr[j] < arr[j-1] {
				tmp = arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = tmp
			}
		}
	}

}
