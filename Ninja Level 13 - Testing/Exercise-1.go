package main

import (
	"fmt"
	 "github.com/Bhavna/Learn To Code - Google's Go/Ninja Level 13 - Testing/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog.YearsTwo(20))
}