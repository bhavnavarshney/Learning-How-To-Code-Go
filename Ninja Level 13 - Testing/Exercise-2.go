package main

import (
	"fmt"
    "github.com/Bhavna/Learn To Code - Google's Go/Ninja Level 13 - Testing/quote"
    "github.com/Bhavna/Learn To Code - Google's Go/Ninja Level 13 - Testing/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}