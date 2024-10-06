package main

import (
	"fmt"
)

// Day 1

func bandNameGenerator() string {
	var city, pet, bandName string
	fmt.Println("Welcome to the band name generator!")
	fmt.Println("What's the name of the city you grew up in?")
	fmt.Scan(&city)
	fmt.Println("What's your pet's name?")
	fmt.Scan(&pet)
	bandName = fmt.Sprintf("Your band name could be %s %s", city, pet)

	return bandName
}

func main() {
	fmt.Println(bandNameGenerator())
}
