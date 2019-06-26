package main

import "fmt"

/*
Hands-on exercise #3
Create TYPED and UNTYPED constants. Print the values of the constants.
*/

const (
	// Untyped Constants
	number = 2
    name = "Apple"
	present = true
	
	//Typed Constants
	stock int =15
	stock_present bool = true
	stock_name string = "Clothes"
)
func main(){
	fmt.Println(number,name,present)
	fmt.Println(stock,stock_name,stock_present)
}