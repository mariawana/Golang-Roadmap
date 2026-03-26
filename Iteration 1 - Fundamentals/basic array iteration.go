package main

import "fmt"

func main() {

	number := [...]int{
		1, 2, 3, 4, 5}

	for i, v := range number {

		fmt.Printf("Index: %d \n", i)
		fmt.Printf("Value: %v \n", v)
	}

}
