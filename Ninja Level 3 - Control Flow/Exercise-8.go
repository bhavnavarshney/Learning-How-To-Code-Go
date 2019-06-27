/*
Hands-on exercise #8
Create a program that uses a switch statement with no switch expression specified.

*/

package main

import "fmt"

func main(){

	switch {
	case true :
		fmt.Println("This will print")
	case false:
		fmt.Println("Yeah! This will not print ")
	}
}