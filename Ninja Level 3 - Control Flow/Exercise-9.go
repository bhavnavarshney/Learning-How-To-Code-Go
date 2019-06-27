/*
Hands-on exercise #9
Create a program that uses a switch statement with the switch expression specified as a variable of 
TYPE string with the IDENTIFIER “favSport”.
*/

package main

import "fmt"

func main(){
	var favSport string = "cricket"
	switch favSport{
	case "cricket" : 
		fmt.Println("Fav Sport : Cricket")
	case "football" :
		fmt.Println("I like Football")
	case "foosball" :
		fmt.Println("Nothing beats Foosball!")
	default :
		fmt.Println("I'm lazy!!")
	}
}