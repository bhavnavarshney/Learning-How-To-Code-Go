package main

import "fmt"

/*
Hands-on exercise #1
Write a program that prints a number in decimal, binary, and hex
*/

func main(){
    x := 10
    fmt.Printf("%d\t%b\t%#x",x,x,x)
}