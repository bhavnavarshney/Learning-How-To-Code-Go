package main

import "fmt"

func main(){
    c := make(chan int)
    // send
    go foo(c)

    // receive
    bar(c)
    
    fmt.Println("Done")
}

//send
func foo(c chan<- int){
	c <- 100
}

//receive
func bar(c <-chan int){
	fmt.Println("Received:",<-c)
}