package main

import "fmt"

func main() {

}

func MaxInt(a, b int) int {
	if a >= b {
		return a
	}
	fmt.Println("max int", a, b)
	return b
}
