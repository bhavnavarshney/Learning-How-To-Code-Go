/*
Hands-on exercise #2
Get this code running:
cs := make(chan<- int)
	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)
	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
*/

package main

import (
	"fmt"
)

func main() {
	cs := make(chan int)  //Send only channel turned into Bidirectional Channel

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
