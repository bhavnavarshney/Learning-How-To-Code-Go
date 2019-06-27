/*
Hands-on exercise #6
Build and use an anonymous func 
*/

package main

import "fmt"

func main(){
    func (){
		fmt.Println("Anonymous Function")
	}()
	fmt.Println("Implementation Done!")
}