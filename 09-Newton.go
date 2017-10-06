// Adapted from https://tour.golang.org/welcome/1s
//
//
// Author: Jamie O'Sullivan;
package main

import (
	"fmt"
	"math"
)

func z_next(z float64, x float64) float64 {

	return z - (((z * z) - x) / (2 * z))
}

func main() {
	//The number to find the SQUARE ROOT
	x := 64.0

	//The Guess
	z := float64(1)

	for z = 1.0; z != z_next(z, x); z = z_next(z, x) {
		fmt.Println("Current Guess:&f \n", z)

	}

	fmt.Println("The square root of &f is &f", x, z)

	fmt.Println("Math square gives the value as &f.\n", math.Sqrt(x))

}
