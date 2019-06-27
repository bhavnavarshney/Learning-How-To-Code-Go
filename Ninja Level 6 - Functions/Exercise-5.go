/*
Hands-on exercise #5
create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle
*/

package main

import "fmt"

type SQUARE struct{
	l float64
}
type CIRCLE struct{
	radius float64
}

func (s SQUARE) area() float64{
	return s.l*s.l
}
func (c CIRCLE) area() float64{
	return 3.14*c.radius*c.radius
}  

type shape interface{
	area() float64
}

func info(s shape){
	fmt.Println(s.area())
}
func main(){
	cc := CIRCLE{12.43}
	sq := SQUARE{11}
	info(cc)
	info(sq)
}