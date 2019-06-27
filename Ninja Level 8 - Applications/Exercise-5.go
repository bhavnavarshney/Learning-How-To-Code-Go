/*
Hands-on exercise #5
Starting with this code, sort the []user by 
age
last
Also sort each []string “Sayings” for each user
print everything in a way that is pleasant
*/

package main

import (
	"fmt" ; "sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLastName []user

func (a ByLastName) Len() int           { return len(a) }
func (a ByLastName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLastName) Less(i, j int) bool { return a[i].Last < a[j].Last }

func main(){
    u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

    // Actual Solution here
	
	// Custom sort by age
	sort.Sort(ByAge(users))
	fmt.Println("\n -------Sort by Age------- \n",users)

	//Custom sort by last name
	sort.Sort(ByLastName(users))
	fmt.Println("\n -------- Sort by Last Name ------------\n",users)

	// Sort sayings
	fmt.Println("\n ---------- Sort by Sayings ----------\n")
	for p,value := range users{
		fmt.Println(" Person ",p)
		fmt.Println(value.First,value.Last,value.Age)
		sort.Strings(value.Sayings)
		for _,val := range value.Sayings {
			fmt.Println("  ",val)
		}
	}
}