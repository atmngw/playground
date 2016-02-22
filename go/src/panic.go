package main

import "fmt"

func main() {

	defer func() {
		err := recover()
		fmt.Printf("%s\n", err)
		//localerr := "ERROR"
		//fmt.Printf("%s\n", localerr)
	}()

	panic("error")

	/*
		err := recover()
		fmt.Println(err)
	*/
}
