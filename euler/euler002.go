package main

import "fmt"

func genrateFibonaciiSeries(limit int, ch chan int) {
	a, b, c := 0, 1, 0
	for {
		c = b + a
		if c > limit { break }
		ch <- c
		a, b = b, c
	}
	ch <- -1
}

func main () {
	c := make(chan int, 100)
	sum := 0
	go genrateFibonaciiSeries(4000000, c)
	for {
		num := <- c
		if num == -1 { break; }
		if (num % 2) == 0 { sum += num }
	}
	fmt.Println(sum)
}