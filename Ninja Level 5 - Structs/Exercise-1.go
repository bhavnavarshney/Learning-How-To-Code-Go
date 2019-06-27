/*
Hands-on exercise #1
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the 
favorite flavors. 

*/

package main

import "fmt"

type person struct{
    first string
    last string
    flavors []string
}
func main(){
    p1 := person{"Ana","Steele",[]string{"Vanilla","Choco","Orange"}}
    p2 := person{"Robin","Dsouza",[]string{"Mix-fruit","Chocolate","Rasberry"}}
    
	fmt.Println(p1.first,p1.last," likes : ")
	for _,v := range p1.flavors{
		fmt.Println(" -- ",v)
	}

	fmt.Println(p2.first,p2.last," likes : ")
	for _,v := range p2.flavors{
		fmt.Println(" -- ",v)
	}
}