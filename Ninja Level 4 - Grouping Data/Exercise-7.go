/*
Hands-on exercise #7
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.
*/

package main

import "fmt"

func main(){
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	ab := [][]string{a,b}

	for i,x := range ab{
		fmt.Println("Record ",i)
		for j,v := range x{
			fmt.Println(" ",j,v)
		}
	}
	
}