/*
Hands-on exercise #3
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

*/

package main

import "fmt"

func main(){
	fmt.Println("I'm hungry!")
	eat()
	defer drink() 
	fmt.Println("Back to Work!")
}

func eat(){
	fmt.Println("Time to Eat")
}

func drink(){
	fmt.Println("But I want to drink Juice!")
}