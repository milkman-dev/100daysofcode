package main

import (
	"fmt"
	"math/rand/v2"
)

// Day 4

func rockPaperScissors() {

	rock := `
    _______
---'   ____)
      (_____)
      (_____)
      (____)
---.__(___)
	`
	paper := `
     _______
---'    ____)____
           ______)
          _______)
         _______)
---.__________)
	`
	scissors := `
    _______
---'   ____)____
          ______)
       __________)
      (____)
---.__(___)
	`

	var playerChoice, randomChoice int
	var computerPick, playerPick string

	randomChoice = rand.IntN(3)

	fmt.Println(`
Welcome to rock, paper, scissors!
- To play rock, press 0
- To play paper, press 1
- To play scissors, press 2
	`)
	fmt.Scan(&playerChoice)

	switch randomChoice {
	case 0:
		computerPick = rock
	case 1:
		computerPick = paper
	case 2:
		computerPick = scissors
	}

	switch playerChoice {
	case 0:
		playerPick = rock
	case 1:
		playerPick = paper
	case 2:
		playerPick = scissors
	}

	if playerPick == computerPick {
		fmt.Println("Player plays:")
		fmt.Println(playerPick)
		fmt.Println("Computer plays:")
		fmt.Println(computerPick)
		fmt.Println("It's a draw!")
	}

	if (playerChoice == 0 && randomChoice == 2) || (playerChoice == 2 && randomChoice == 1) || (playerChoice == 1 && randomChoice == 0) {
		fmt.Println("Player plays:")
		fmt.Println(playerPick)
		fmt.Println("Computer plays:")
		fmt.Println(computerPick)
		fmt.Println("You win!")
	}

	if (playerChoice == 0 && randomChoice == 1) || (playerChoice == 1 && randomChoice == 2) || (playerChoice == 2 && randomChoice == 0) {
		fmt.Println("Player plays:")
		fmt.Println(playerPick)
		fmt.Println("Computer plays:")
		fmt.Println(computerPick)
		fmt.Println("You lose!")
	}

	fmt.Println("Game Over!")

}

func main() {

	for {
		rockPaperScissors()

		var playAgain string
		fmt.Print("Play again? y/n: ")
		fmt.Scan(&playAgain)

		if playAgain == "n" {
			break
		}
	}
}
