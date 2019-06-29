/*
Hands-on exercise #6
write a program that
puts 100 numbers to a channel
pull the numbers off the channel and print them
*/

package main

import "fmt"

func main(){
    c := make(chan int)
    go func(){
        for i:=0;i<100;i++{
            c <- i
        }
        close(c)
    }()

    receive(c)
    fmt.Println("Done")
}

func receive(c chan int){
    for i := range c{
        fmt.Println(i)
    }
}