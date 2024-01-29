package main

import (
	"fmt"
)

func Cats(start, finish int) int {
	result := (finish - start) / 3
	result += (finish - start) % 3
	return result
}

func main() {
	fmt.Println(Cats(1, 5))
	fmt.Println(Cats(1, 1))
	fmt.Println(Cats(2, 5))
	fmt.Println(Cats(2, 4))
}
