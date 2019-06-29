/*
Hands-on exercise #7
write a program that
launches 10 goroutines
each goroutine adds 10 numbers to a channel
pull the numbers off the channel and print them
*/

package main

import "fmt"

func main(){
    c := make(chan int)
	go func(){
	for i:=1;i<=10;i++{
        send(c,i)
	}
	close(c)
}()

    for v := range c{
        fmt.Println(v)
    }
}

func send(c chan int,index int){
    for i:=0;i<10;i++{
        c <- i
	}

}