// Adapted fromhttps://adriann.github.io/programming_problems.html
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
	var myname string
	myrand := xrand(1, 6)
	guessTaken := 0
	var guess int

	fmt.Println("Hello! What is your name?")
	fmt.Scanf("%s", &myname)
	fmt.Printf("Well, %s, I am thinking of a number between 1 and 6.\n", myname)

	//this is the while loop
	for guessTaken < 6 {
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
		fmt.Printf("Good job %s! You guessed my number in %d tries\n", myname, guessTaken)
	} else {
		fmt.Printf("Nope. The number I had in mind was %d\n", myrand)
	}
}
