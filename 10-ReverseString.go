// Adapted fromhttp://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html
//
//
// Author: Jamie O'Sullivan;
package main

import "fmt"

func main() {
	s := "Hello, playground"
	fmt.Println(s)

	fmt.Println(Reverse(s))
	fmt.Println(ReverseUsingRunes(s))

}

func Reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func ReverseUsingRunes(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
