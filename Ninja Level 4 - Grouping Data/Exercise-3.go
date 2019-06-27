/*
Hands-on exercise #3
Using the code from the previous example, use SLICING to create the following new slices which are then printed:
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]

*/

package main

import "fmt"

func main(){
    a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x1 := a[:5]
	x2 := a[5:]
	x3 := a[2:7]
	x4 := a[1:6]

	fmt.Println(x1)
	fmt.Println(x2)
	
	fmt.Println(x3)
	fmt.Println(x4)

}