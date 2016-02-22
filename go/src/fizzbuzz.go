package main

import "fmt"

type FizzBuzz struct {
	max int
}

func main() {
	fizzbuzz := FizzBuzz{100}
	fizzbuzz.do()
}

func (fb FizzBuzz) do() {

	for i := 1; i <= fb.max; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}

}
