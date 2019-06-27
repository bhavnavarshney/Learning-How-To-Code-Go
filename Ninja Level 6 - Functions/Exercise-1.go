/*
Hands on exercise
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results

*/

package main

import "fmt"

func main(){
	x := foo()
	y, brand := bar()
	fmt.Println(x)
	fmt.Println(y,brand) 
}

func foo() int{
	return 100
}
func bar() (int,string){
	return 200,"Pantaloons"
}