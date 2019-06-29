/*
Hands-on exercise #4
Starting with this code, pull the values off the channel using a select statement

*/
package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c <-chan int,q chan int){
	for{
		select{
		case v := <-c :
			fmt.Println("Received:",v)
		case <-q :
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
    go func(){
	for i := 0; i < 100; i++ {
		c <- i
	}
	q <- 0
    close(c)
    }()

	return c
}
