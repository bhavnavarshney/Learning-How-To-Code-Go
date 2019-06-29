/*
Hands-on exercise #3
Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that has a value
 of type error as a parameter. Create a value of type “customErr” and pass it into “foo”. 
 */

 package main

 import "fmt"

 type customErr struct{
	 msg string
 }

 func (c customErr) Error() string{
	 return fmt.Sprintf("Error %v",c.msg)
 }

 func foo(e error){
	fmt.Println("Error in foo: ",e)
 }
 func main(){
	x := customErr{"Some Error here!"}
	foo(x)
 }