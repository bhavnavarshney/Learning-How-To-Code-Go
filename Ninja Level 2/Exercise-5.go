package main

import "fmt"

/*
Hands-on exercise #5
Create a variable of type string using a raw string literal. Print it.
*/

func main(){
	s :=  `Here there is something 
	interesting
	and here it goes
	"Raw String Literal"`
    fmt.Println(s)
}