/*
Hands-on exercise #1
in addition to the main goroutine, launch two additional goroutines
each additional goroutine should print something out
use waitgroups to make sure each goroutine finishes before your program exists

*/

package main
import (
	"fmt"
	"sync"
	"runtime"
)

var wg sync.WaitGroup

func main(){
	fmt.Println("No of CPU ",runtime.NumCPU())
	fmt.Println("No of Go Routines ",runtime.NumGoroutine())
	wg.Add(2)
    go Odd()
	go Even()
	fmt.Println("No of Go Routines ",runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Time to EXIT")
	fmt.Println("No of CPU ",runtime.NumCPU())
	fmt.Println("No of Go Routines ",runtime.NumGoroutine())
}

func Odd(){
    for i:=1;i<=10;i++{
        if i%2!=0{
        fmt.Println("Odd : ",i)
    }
	}
	fmt.Println("No of Go Routines ",runtime.NumGoroutine())
	wg.Done()
}
func Even(){
    for i:=1;i<=10;i++{
        if i%2==0{
        fmt.Println("Even : ",i)
    }
	}
	fmt.Println("No of Go Routines ",runtime.NumGoroutine())
	wg.Done()
}