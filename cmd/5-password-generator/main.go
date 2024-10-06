package main

import (
	"fmt"
	"math/rand/v2"
)

// Day 5

func passwordGenerator(letters, symbols, numbers int) string {
	var charArray []string
	var password string
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbolsChars := `!#$%&'()*+,-./:;<=>?@^_{|}~`
	numbersChars := "1234567890"
	numbersLen := len(numbersChars)
	symbolsLen := len(symbolsChars)
	alphabetLen := len(alphabet)

	for i := 0; i < letters; i++ {
		charArray = append(charArray, string(alphabet[rand.IntN(alphabetLen)]))
	}

	for i := 0; i < symbols; i++ {
		charArray = append(charArray, string(symbolsChars[rand.IntN(symbolsLen)]))
	}

	for i := 0; i < numbers; i++ {
		charArray = append(charArray, string(numbersChars[rand.IntN(numbersLen)]))
	}

	fmt.Println(charArray)
	for index := range charArray {
		rIndex := rand.IntN(index + 1)
		charArray[index], charArray[rIndex] = charArray[rIndex], charArray[index]
	}
	fmt.Println(charArray)

	for _, char := range charArray {
		password += char
	}

	return password
}

func main() {
	var letters, symbols, numbers int
	fmt.Println("Welcome to the GoPassword Generator!")
	fmt.Println("How many letters would you like in your password?")
	fmt.Scan(&letters)
	fmt.Println("How many symbols would you like?")
	fmt.Scan(&symbols)
	fmt.Println("How many numbers would you like?")
	fmt.Scan(&numbers)

	password := passwordGenerator(letters, symbols, numbers)
	fmt.Printf("Your password is: %s", password)
}
