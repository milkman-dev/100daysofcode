package main

import "fmt"

func tipCalculator(bill float64) []string {
	var tipRates []float64
	var tip float64
	var tipOptions []string

	tipRates = []float64{0.10, 0.15, 0.20}

	for _, v := range tipRates {
		tip = bill * v
		tipOptions = append(tipOptions, fmt.Sprintf("%.2f,", tip))
	}

	return tipOptions
}

func splitBill() {
	var bill, tip float64
	var people int
	fmt.Println("Welcome to the tip calculator!")
	fmt.Print("What was the total bill? (Include decimal) $")
	fmt.Scan(&bill)
	fmt.Printf("How much tip would you like to give? %v: $", tipCalculator(bill))
	fmt.Scan(&tip)
	fmt.Print("How many people will split the bill? ")
	fmt.Scan(&people)

	dividedBill := (bill + tip) / float64(people)

	fmt.Printf("Each person should pay: $%.2f", dividedBill)
}

func main() {
	splitBill()
}
