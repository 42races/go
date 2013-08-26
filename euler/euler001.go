package main

import "fmt"

func sumOfMultiplesOf(x, limit int, c chan int) int {
	last_term := limit - (limit % x)
	if last_term == 1000 { last_term -= x }
	n := last_term / x
	sum := (n * (n + 1) * x ) / 2
	c <- sum
	return sum
}

func main() {
	sum, items := 0, [...]int{ 3, 5, 15 }
	c := make(chan int)
	for _, i := range items {
		go sumOfMultiplesOf(i, 1000, c)
	}

	sum = (<- c) + (<- c) - (<- c)
	fmt.Println("Ans: %d", sum)
}