/*
Hands-on exercise #1
get this code working:
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}

1. using func literal, aka, anonymous self-executing func

2. a buffered channel

*/

package main

import (
	"fmt"
)

func main() {
	//Using func literal, aka, anonymous self-executing func
	c := make(chan int)
	go func(){
		c <- 42
	}()
	fmt.Println(<-c)
	
	//Using Buffered Channel
	c2 := make(chan int,1)

	c2 <- 42

	fmt.Println(<-c2)

}
