// Adapted fromhttps://gist.github.com/pyk/10519339
//
//
// Author: Jamie O'Sullivan;
package main

import "fmt"

func main() {
	var smallest, biggest int
	x := []int{
		42, 93, 88, 77,
		57, -1, 0, 70,
		34, 31, 83, 29,
		20, 98, 7, 1,
	}

	if len(x) == 0 {
		fmt.Println("Array must be non empty")
	} else {
		smallest = x[0]
		biggest = x[0]
		for _, v := range x {
			if v > biggest {
				biggest = v
			} else if v < smallest {
				smallest = v
			}
		}

		fmt.Println("The biggest number is ", biggest)
		fmt.Println("The smallest number is ", smallest)
	}

}
