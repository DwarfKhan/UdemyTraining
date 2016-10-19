package main

import "fmt"

func main() {
	for i := 1; i < 11; i++ {
		fmt.Printf("%d \t %b \t %X \t %q \n", i, i, i, i)
	}
}
