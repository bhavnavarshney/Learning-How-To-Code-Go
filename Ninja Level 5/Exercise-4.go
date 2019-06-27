/*
Hands-on exercise #4
Create and use an anonymous struct.

*/

package main

import "fmt"

func main(){
    s1 := struct{
        name string
        subject map[string]int
        companies []string
    }{
        "James",
        map[string]int{
            "ML":80,
            "AI":65,
            "CC":90,
        },
        []string{
            "Amazon","Facebook","Google",
		},
		
    }

	fmt.Println(s1)
	for key,val := range s1.subject{
		fmt.Println(key,val)
	}

	for _,val := range s1.companies{
		fmt.Println(val)
	}
}