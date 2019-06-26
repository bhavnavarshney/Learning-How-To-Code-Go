/*
Hands-on exercise #6
Create a program that shows the “if statement” in action.

*/

package main

import "fmt"

func main(){
    for x:= 0 ;x<10;x++{
		if x%2==0{
			fmt.Println("Even Number - ",x)
		}
	}
}