package main

import "fmt"

/*
Hands-on exercise #6
Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
*/

const(
    a = 2019 + iota
    b = 2019 +iota
    c = 2019 + iota
    d = 2019 +iota
)
func main(){
    fmt.Println(a,b,c,d)
}