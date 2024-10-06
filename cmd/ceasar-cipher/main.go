package main

import (
	"fmt"
    "strings"
)

const stringSet = "abcdefghijklmnopqrstuvwxyz"

var stringCharset = []rune(stringSet)
var charsetLen = len(stringCharset)

func runeIndex(charSet []rune, char rune) int {
	for i, v := range charSet {
		if v == char {
			return i
		}
	}

	return 0
}

func encrypt(text string, offset int) string {
	var encryptedText []rune
	var charIndex, offsetIndex int

	for _, char := range text {
		charIndex = runeIndex(stringCharset, char)
		offsetIndex = (charIndex + offset) % charsetLen
		encryptedText = append(encryptedText, stringCharset[offsetIndex])
	}

	return string(encryptedText)
}

func decrypt(text string, offset int) string {
	var encryptedText []rune
	var charIndex, offsetIndex int

	for _, char := range text {
		charIndex = runeIndex(stringCharset, char)
		offsetIndex = (charIndex - offset) % charsetLen
        if offsetIndex < 0 {
            offsetIndex += charsetLen
        }

		encryptedText = append(encryptedText, stringCharset[offsetIndex])
	}

	return string(encryptedText)
}

func cipher() {
	var cipherChoice, cipherText, restart string
	var cipherOffset int

	fmt.Println(`
                    ⠀⠀⠀⠀⠀⠀⠀⠀⢀⣤⣶⣶⠿⠿⠿⣶⣦⣀⠀⠀⠀
                ⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⡾⠛⠉⠀⠀⠀⠀⠀⠀⠉⠻⣧⡀⠀
                ⢠⣄⣀⣀⣀⣀⣀⣀⣀⣴⠋⠀⠀⠀⠀⠀⣴⣆⠀⠀⠀⠀⠘⣿⡀
                ⠀⠙⠻⣿⣟⠛⠛⠛⠋⠁⠀⠀⠀⠀⠀⠘⠿⠋⠀⠀⠀⠀⠀⣿⡇
                ⠀⠀⠀⠀⠙⢷⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣾⡇
                ⠀⠀⠀⠀⠀⠀⠘⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣽⠃
                ⠀⠀⠀⠀⠀⠀⢰⡿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠀
                ⠀⠀⠀⠀⠀⠀⣾⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢰⡿⠀
                ⠀⠀⠀⠀⠀⢸⡟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣸⠃⠀
                ⠀⠀⠀⠀⢀⡿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢠⡟⠀⠀
                ⠀⠀⠀⠀⣾⠇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣼⠇⠀⠀         

| '_ \| | | | '_ \  | '_ \| | | | '_ \  / __|  / __| | '_ \| '_ \ / _ \ '__|
| |_) | |_| | | | | | |_) | |_| | | | | \__ \ | (__| | |_) | | | |  __/ |   
| .__/ \__,_|_| |_| | .__/ \__,_|_| |_| |___/  \___|_| .__/|_| |_|\___|_|   
| |                 | |                              | |                    
|_|                 |_|                              |_|                    `)

	for {

		fmt.Println("Do you wish to encrypt or decrypt? Type E (encrypt) or D (decrypt)")
		fmt.Scan(&cipherChoice)

		if strings.ToUpper(cipherChoice) == "E" {
			fmt.Println("Type your message (ONLY WORDS ALLOWED, NO SYMBOLS OR NUMBERS)")
			fmt.Scan(&cipherText)
			fmt.Println("Type the offset number (ONLY NUMBERS ALLOWED)")
			fmt.Scan(&cipherOffset)
			fmt.Println()
			cipherText = encrypt(cipherText, cipherOffset)
			fmt.Printf("Your cipher text is: %v", cipherText)
			fmt.Println()
			fmt.Println("You you want to start from the beginning? (Y/N)")
			fmt.Scan(&restart)
		}

		if strings.ToUpper(cipherChoice) == "D" {
			fmt.Println("Type your message (ONLY WORDS ALLOWED, NO SYMBOLS OR NUMBERS)")
			fmt.Scan(&cipherText)
			fmt.Println("Type the offset number (ONLY NUMBERS ALLOWED)")
			fmt.Scan(&cipherOffset)
			fmt.Println()
			cipherText = decrypt(cipherText, cipherOffset)
			fmt.Printf("Your decipher text is: %v", cipherText)
			fmt.Println()
			fmt.Println("You you want to start from the beginning? (Y/N)")
			fmt.Scan(&restart)
		}

		if strings.ToUpper(restart) == "N" {
			break
		}
	}
}

func main() {
	cipher()
}
