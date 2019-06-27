/*
Hands-on exercise #4
Fix the race condition you created in the previous exercise by using a mutex
it makes sense to remove runtime.Gosched()
*/

package main

import (
	"fmt"
	"sync"
	
)

func main(){
	incrementer :=0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mutex sync.Mutex
	for i:=0;i<gs;i++{
		go func(){
			mutex.Lock()
			x := incrementer
			x++
			incrementer = x
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter ",incrementer)
}
