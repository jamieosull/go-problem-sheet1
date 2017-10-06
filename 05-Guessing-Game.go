// Adapted fromhttp://golangcookbook.blogspot.ie/2012/11/guess-number-game-in-golang.html
//
//
// Author: Jamie O'Sullivan;
package main

import (
	"fmt"
	"math/rand"
	"time"
)

//this generates random number between given range
func xrand(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	myrand := xrand(1, 10)
	guessTaken := 0
	var guess int

	fmt.Printf("Hi..Take a guess of a number between 1 and 10.\n")

	//this is the while loop
	for guessTaken < 10 {
		fmt.Println("Take a guess...")
		fmt.Scanf("%d", &guess)
		guessTaken++
		if guess < myrand {
			fmt.Println("Your guess is too low.")
		}
		if guess > myrand {
			fmt.Println("Your guess is too high.")
		}
		if guess == myrand {
			break
		}
	}
	if guess == myrand {
		fmt.Printf("Good job..You guessed my number in %d tries\n", guessTaken)
	} else {
		fmt.Printf("Nope. The number I had in mind was %d\n", myrand)
	}
}
