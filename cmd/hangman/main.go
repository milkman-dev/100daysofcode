package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func getWord(index int) string {
	words := []string{
		"table", "chair", "lamp", "sofa", "bed", "pillow", "blanket", "desk", "pen", "notebook",
		"book", "cup", "plate", "fork", "spoon", "knife", "mirror", "window", "door", "curtain",
		"towel", "phone", "charger", "key", "wallet", "backpack", "watch", "glasses", "shoes", "socks",
		"toothbrush", "toothpaste", "soap", "shampoo", "conditioner", "comb", "hairbrush", "razor", "tissue", "napkin",
		"bottle", "remote", "television", "computer", "mouse", "keyboard", "headphones", "speaker", "clock", "calendar",
		"pencil", "eraser", "scissors", "tape", "stapler", "envelope", "paperclip", "ruler", "glue", "highlighter",
		"paintbrush", "canvas", "broom", "dustpan", "vacuum", "mop", "bucket", "sponge", "soap", "laundry",
		"hanger", "iron", "lightbulb", "fan", "heater", "blender", "toaster", "microwave", "oven", "stove",
		"refrigerator", "freezer", "dishwasher", "kettle", "pot", "spatula", "whisk", "measuring", "cutting", "trash",
		"recycling", "fan", "curtain", "shoe", "coat", "doormat", "shelf", "basket", "drawer", "candle",
	}

	return words[index]
}

func maskWord(word string) []rune {
	maskedWord := make([]rune, len(word))
	for i := range maskedWord {
		maskedWord[i] = '_'
	}

	return maskedWord
}

// Day 6

func hangman() {
	var letter string
	fmt.Println("Welcome to Hangman Game!")
	word := getWord(rand.IntN(100))
	maskWord := maskWord(word)
	lives := 6

	for {
		switch lives {
		case 6:
			fmt.Printf(`You have %d lives! ( • ᴗ - )✧`, lives)
		case 5:
			fmt.Printf(`You have %d lives! (˶ᵔ ᵕ ᵔ˶)`, lives)
		case 4:
			fmt.Printf(`You have %d lives! (・_・)`, lives)
		case 3:
			fmt.Printf(`You have %d lives! (￢_￢)`, lives)
		case 2:
			fmt.Printf(`You have %d lives! (｡>﹏<)`, lives)
		case 1:
			fmt.Printf(`You have %d lives! (｡Ó﹏Ò｡)`, lives)
		case 0:
			fmt.Printf(`You have %d lives! 💀`, lives)
		}

		if lives < 1 {
			fmt.Println()
			fmt.Printf("Game Over, the word was: %s", word)
			break
		}

		fmt.Println()
		fmt.Printf("Word to guess: %v", string(maskWord))
		fmt.Println()
		fmt.Print("Guess a letter in the word: ")
		fmt.Scan(&letter)
		fmt.Println()

		if !strings.Contains(word, letter) {
			lives--
		}

		for i, char := range word {
			if letter == string(char) {
				maskWord[i] = rune(letter[0])
			}
		}

		if string(maskWord) == word {
			fmt.Printf("-------> %v \n", string(maskWord))
			fmt.Println(`You won! (⸝⸝ᵕᴗᵕ⸝⸝)`)
			break
		}
	}
}

func main() {
	hangman()
}
