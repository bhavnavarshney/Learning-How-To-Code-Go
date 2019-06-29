/*
Hands-on exercise #5
Show the comma ok idiom starting with this code.

*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func(){
		c <- 100
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)
	
	v, ok =	<-c
	fmt.Println(v, ok)
}

