/*
Hands-on exercise #9
Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

*/

package main

import "fmt"

func main(){
    a := map[string][]string{
		`bond_james` : []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss` : []string{ `James Bond`, `Literature`, `Computer Science`},
		`no_dr` : []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	
	a[`Ana_Panna`] = []string{"Men","Chocolate ice cream","Perry-Perry"}
	for key,value := range a{
		fmt.Println("Record ",key,value)
		for index,val := range value{
			fmt.Println("   ",index,val)
		}
	}
}