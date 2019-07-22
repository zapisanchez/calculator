package main

import (
	"fmt"
	"math"
)

type meal struct {
	ticketCount int
	ticketValue float64
	totalCount  float64
	deltaMoney  float64
}

func roundUp(input float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * input
	round = math.Ceil(digit)
	newVal = round / pow
	return
}

/*
	getRoundTicketNumber(price float32, ticketValue float32) int
	Price       : Cost of the meal
	ticketValue : Value of each ticket

	return ROunded ticket number
	bool flag to know if deltamoney must be taked into account
*/
func getRoundTicketNumber(price float64, ticketValue float64) (result int, flag bool) {

	var thrs = 0.80
	preciseDiv := price / ticketValue

	//round 2 decimals
	preciseDiv = roundUp(preciseDiv, 2)

	integerPart, div := math.Modf(preciseDiv)

	if div < thrs {
		result = int(integerPart)
		flag = true
		fmt.Println("IF: ", result)
	} else {
		result = int(math.Round(preciseDiv))
		flag = false
		fmt.Println("ELSE: ", result)
	}

	fmt.Println("Resultado preciseDiv: ", preciseDiv)
	fmt.Println("Resultado div: ", div)
	fmt.Println("Resultado integerPart: ", integerPart)
	fmt.Println("Resultado a devolver: ", result)
	return
	//return int(price / ticketValue)
}

func getDeltaMoney(theMeal *meal) {
	temp1 := float64(theMeal.ticketCount) * theMeal.ticketValue

	temp2 := theMeal.totalCount - temp1

	theMeal.deltaMoney = temp2

}

func calcDinner(ticketVal float64, totalCount float64) meal {
	var result meal
	var needMoney bool

	result.ticketValue = ticketVal
	result.totalCount = totalCount
	result.ticketCount, needMoney = getRoundTicketNumber(totalCount, ticketVal)

	if needMoney {
		getDeltaMoney(&result)
	}

	return result
}
