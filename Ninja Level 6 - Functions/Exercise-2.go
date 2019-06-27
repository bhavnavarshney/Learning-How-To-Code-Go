/*
Hands-on exercise #2
create a func with the identifier foo that 
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in
create a func with the identifier bar that 
takes in a parameter of type []int
returns the sum of all values of type int passed in

*/

package main

import "fmt"

func main(){
	xx := []int{1,2,3,4,5}
	sum := foo(xx...)
	fmt.Println("Sum of the numbers are = ",sum)

	xy := []int{10,20,30,40,40}
	sum2 := bar(xy)
	fmt.Println("Sum of the numbers are = ",sum2)
}
func foo(x ...int) int{
	total :=0
	for _,value := range x {
		total+=value
	}
	return total
}
func bar(x []int) int{
	total :=0
	for _,value := range x {
		total+=value
	}
	return total
}

