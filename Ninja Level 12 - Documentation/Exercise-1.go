package main

import (
	"fmt"
	"dog"
)

type canine struct{
    name string
    age int
}

func main(){
    c := canine{"Fido",dog.Years(10)}
    fmt.Println(c)
}