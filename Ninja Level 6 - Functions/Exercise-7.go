/*
Hands-on exercise #7
Assign a func to a variable, then call that func

*/

package main

import "fmt"

func main(){
    x := func(){
		fmt.Println("Hey! I'm calling you")
	}
	x()
	fmt.Printf("%T",x)
}