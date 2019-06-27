/*
Hands-on exercise #8
Create a func which returns a func
assign the returned func to a variable
call the returned func
*/

package main

import "fmt"

func main(){
	yy := foo()
	fmt.Println(yy())
	fmt.Printf("%T",yy)
}

func foo() func() int {
	return func() int{
		return 100
	}
}