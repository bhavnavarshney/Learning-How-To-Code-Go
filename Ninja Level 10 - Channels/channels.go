// Channel

package main

import "fmt"

func main(){
    c := make(chan int)
     go func(){
        c <- 45
    }()
	fmt.Println(<-c)
	
	c2 := make(chan int,1)
	c2 <- 100
	fmt.Println(<-c2)
}