/*
Hands-on exercise #1
Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise 
- remember to ask yourself what you need to do to EXPORT a value from a package.
*/

package main

import (
	"fmt"; "encoding/json"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// Actual Solution below
	json_data,err := json.Marshal(users)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(json_data))
}
