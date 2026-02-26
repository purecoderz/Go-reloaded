package main

import (
	"fmt"
	"os"
	"reload/packs"
	"strings"
)

func main() {
	// Check if user pass 2 params
	if len(os.Args) != 3 {
		fmt.Println("Please make sure you provide 2 .txt file")
		os.Exit(1)
	}

	// confirm that the file exits
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// check empty file
	if len(file) < 1 {
		fmt.Println("Your sample file is empty")
		os.Exit(1)
	}
	// convert byte to string and split it
	testFile := string(file)

	words := strings.Fields(testFile)
	start := -1
	end := -1
	for index := 0; index < len(words); index++ {
		switch words[index] {
		case "'s", "'ve", "'t", "'re":
			words, index = packs.Apostrophe(words, index)
		case "'":
			words, index, start, end = packs.Quotes(words, index, start, end)
		case "(cap)":
			words, index = packs.ChangeCase(index, words, packs.ToTitleCase)
		case "(cap,":
			words, index = packs.ChangeMultipleCase(index, words, packs.ToTitleCase)
		case "(up)":
			words, index = packs.ChangeCase(index, words, strings.ToUpper)
		case "(up,":
			words, index = packs.ChangeMultipleCase(index, words, strings.ToUpper)
		case "(low)":
			words, index = packs.ChangeCase(index, words, strings.ToLower)
		case "(low,":
			words, index = packs.ChangeMultipleCase(index, words, strings.ToLower)
		case "(hex)":
			words, index = packs.NumConversion(16, index, words)
		case "(bin)":
			words, index = packs.NumConversion(2, index, words)
		default:
			words, index = packs.CheckPunctuation(index, words)
		}

	}

	// fixed vowels 
	packs.FixVowels(words)

	formattedText := strings.Join(words, " ")

	err = os.WriteFile(os.Args[2],[]byte(formattedText), 0644) 
	if err != nil {
		fmt.Println("Error writing to File:",err)
		os.Exit(1)
	}

	fmt.Println("Completed sucessful")

}
