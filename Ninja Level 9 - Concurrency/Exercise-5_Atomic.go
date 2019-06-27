/*
Hands-on exercise #5
Fix the race condition you created in exercise #4 by using package atomic
*/

package main

import (
    "fmt"
	"sync/atomic"
	"sync"
)
func main(){
	 var incrementer int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i:=0;i<gs;i++{
		go func(){
			atomic.AddInt64(&incrementer,1)
			fmt.Println("Incrementer\t ",atomic.LoadInt64(&incrementer))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter ",incrementer)
}