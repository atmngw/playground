package main

import "fmt"

func pow(init int) chan int {
	n := init
	yield := make(chan int)
	go func() {
		for {
			yield <- n
			/* A */
			n *= n
		}
	}()
	return yield
}

func main() {
	generator := pow(5)

	fmt.Printf("num=%d\n", <-generator)
	fmt.Printf("num=%d\n", <-generator)
	fmt.Printf("num=%d\n", <-generator)
}
