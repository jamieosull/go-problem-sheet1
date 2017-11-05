// Adapted from https://play.golang.org/p/J4gxs46HET
//
//
// Author: Jamie O'Sullivan;
package main

import (
	"fmt"
	"math"
)

func znext(z float64, x float64) float64 {

	return z - (((z * z) - x) / (2 * z))
}

func main() {
	//The number to find the SQUARE ROOT
	x := 64.0

	//The Guess
	z := float64(1)

	for z = 1.0; z != znext(z, x); z = znext(z, x) {
		fmt.Println("Current Guess:&f \n", z)

	}

	fmt.Println("The square root of &f is &f", x, z)

	fmt.Println("Math square gives the value as &f.\n", math.Sqrt(x))

}
