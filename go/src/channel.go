package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		fmt.Printf("Hello, Another World!\n")
		ch <- 1 /* channelに入力 */
	}()

	fmt.Printf("Hello, World!\n")

	<-ch /* channelから受信して値は捨てる */
}
