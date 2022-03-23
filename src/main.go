package main

import (
	"fmt"
)

func main() {
	fmt.Println(Naive(6, 3, 10))
	fmt.Println(Smart(6, 3, 10))
	fmt.Println(Binomial(6, 3, 10))
}
