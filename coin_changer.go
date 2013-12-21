package main

import "fmt"

var Denominations = []int{500, 100, 50, 20, 10, 5, 2, 1}

func main() {
	fmt.Print("Enter the amount: ")
	amt := 0

	fmt.Scanf("%d", &amt)

	if amt > 1000 {
		fmt.Println("Can't get change for amount > 1000")
	} else if amt <= 1 {
		fmt.Println("Sorry No changes")
	} else {
		changes := getChange(amt)
		fmt.Println(changes)
	}
}

func getChange(amt int) []int {
	tmp := amt
	changes := make([]int, 0, 12)

	for _, dm := range Denominations {
		for (amt-dm) >= 0 && tmp != dm {
			changes = append(changes, dm)
			amt -= dm
		}

		if amt == 0 {
			break
		}
	}

	return (changes)
}
