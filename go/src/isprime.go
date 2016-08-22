// isprime.go
package main

import (
	"fmt"
	"math"
)

type Result struct {
	isPrime bool
}

func (r Result) put() {
	if r.isPrime {
		fmt.Println("素数である")
	} else {
		fmt.Println("素数ではない")
	}
}

func main() {
	var number int
    fmt.Print("input number -->")
	fmt.Scan(&number)

	r := Result{isPrice(number)}
	r.put()
}

func isPrice(num int) bool {
	if num <= 1 {
		return false
	}

	if num > 2 {
		if num%2 == 0 {
			return false
		}
	}

	max := int(math.Sqrt(float64(num)))
	fmt.Println("num", num, "max", max)
	for i := 2; i <= max; i++ {
		//fmt.Println(i)
		if num%i == 0 {
			return false
		}
	}
	return true
}
