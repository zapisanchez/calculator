package main

import (
	"fmt"
)

func main() {
	var ticketPrice float64
	var totalCost float64

	fmt.Println("************** TICKET CALCULATOR **************")
	fmt.Print("Ticket value: ")
	fmt.Scanln(&ticketPrice)

	fmt.Print("Total Count: ")
	fmt.Scanln(&totalCost)

	result := calcDinner(ticketPrice, totalCost)

	fmt.Println("")
	fmt.Println("-------------- RESULTS --------------")

	fmt.Println("Total Meal cost: ", result.totalCount, "€")
	fmt.Println("Ticket Value: ", result.ticketValue, "€")
	fmt.Println("Number of tickets: ", result.ticketCount)
	fmt.Println("Money to Pay: ", result.deltaMoney, "€")
}
