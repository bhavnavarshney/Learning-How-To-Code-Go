/*
Hands-on exercise #1
Create a value and assign it to a variable. 
Print the address of that value.

*/

package main

import "fmt"

func main(){
    x:=100
    fmt.Println(&x)
}