/*
Hands-on exercise #4
Create a for loop using this syntax
for { }
Have it print out the years you have been alive

*/

package main

import "fmt"

func main(){
	years := 1997
	for {
		fmt.Println(years)
		if(years==2019){
			break
		}
		years++
	}
}