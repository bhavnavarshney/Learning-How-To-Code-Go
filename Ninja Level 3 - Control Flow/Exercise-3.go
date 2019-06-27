/*
Hands-on exercise #3
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
*/

package main

import "fmt"

func main(){
	years := 1997
	for years < 2020 {
		fmt.Println(years)
		years++;
	}
}