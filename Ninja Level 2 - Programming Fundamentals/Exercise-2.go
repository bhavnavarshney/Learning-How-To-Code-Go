package main

import "fmt"

/*
Hands-on exercise #2
Using the following operators, write expressions and assign their values to variables:
==
<=
>=
!=
<
>
Now print each of the variables
*/

func main(){
	x := 10 
	y := 20

	a := (x == y)
	b := (x <= y)
	c := (x >= y)
	d := (x!=y)
	e := (x<y)
	f := (x>y)
	fmt.Println(a,b,c,d,e,f)
}