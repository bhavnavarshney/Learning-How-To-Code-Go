/*
Hands-on exercise #7
Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
*/

package main

import "fmt"

func main(){
	x := "heavy"
	if (x=="light"){
		fmt.Println("Light Rain Shower")
	}else if(x=="medium"){
		fmt.Println("Medium Rain Shower")
	}else{
		fmt.Println("Heavy Rains!")
	}
}