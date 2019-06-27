/*
Hands-on exercise #2
Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
 Access each value in the map. Print out the values, ranging over the slice.

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
    
	list := map[string]person{
		"Steele" : p1,
		"Dsouza" : p2,
	}
	
	for _,value := range list{
		fmt.Println("*************************")
		fmt.Println("I'm",value.first,value.last)
		fmt.Println("I like : ")
		for _,val := range value.flavors{
			fmt.Println("  ",val)
		}
	}
}