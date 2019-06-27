/*
Hands-on exercise #2
Starting with this code, unmarshal the JSON into a Go data structure. Hint: https://mholt.github.io/json-to-go/ 
*/

package main

import (
	"fmt"; "encoding/json"
)

//Part of Actual Solution
type person struct {
	First string        
	Last string			
	Age int				
	Sayings []string	
} 

func main() {
	//Given Code
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	// Actual Solution here 
	var Person []person
	bs := []byte(s)
	err := json.Unmarshal(bs,&Person)
	if err!=nil{
		fmt.Println("ERROR ",err)
	}
	fmt.Println("********************************************\n\n")
	fmt.Println(Person)
	fmt.Println("\n\n Displaying GO ")
	for no,val := range Person{
		fmt.Println("Person - ",no)
		fmt.Println(val.First,val.Last,val.Age)
		for _,value := range val.Sayings {
			fmt.Println(value)
		}
	}
    
}
