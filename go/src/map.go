package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "Finland"
	m[2] = "Norway"
	m[3] = "Sweden"
	mapfor(m)
}

/* mapのイテレーション*/
func mapfor(m map[int]string) {

	for k, v := range m {
		fmt.Printf("%s => %d\n", k, v)
	}
}
