/*
Hands-on exercise #2
Using a COMPOSITE LITERAL: 
create a SLICE of TYPE int
assign 10 VALUES 
Range over the slice and print the values out. 
Using format printing
print out the TYPE of the slice

*/

package main

import "fmt"

func main(){
	a := []int{1,2,3,4,5,6,7,8,9,0}
	for i,v := range a{
		fmt.Println("Index ",i," = ",v)
	}
	fmt.Printf("%T",a)

}