package main

import "fmt"

type Number uint64

func (p Number) isPrime() bool {
	for i:=3; i <= (p / 2); i+=3 {
		if p % i == 0 {
			return false
		}
	}
	return true
}

func main() {
	v := Number{ 12 }
	fmt.Print(v.isPrime)
}