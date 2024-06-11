package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type word struct {
	definition string
	speechType string
}

func hashFunction(word string) int {
	runeArray := []rune(word)
	hashKey := 0
	for i := 0; i < len(runeArray); i++ {
		hashKey += int(runeArray[i])
	}
	return hashKey
}
func loadWords() {
	file, err := os.Open("Dictionary-2.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 2 && strings.Contains(text, "  ") {
			end := strings.Index(text, "  ")
			word := text[0:end]
			definition := text[end:]
			fmt.Println(word, definition)
		}
	}
}
func main() {
	fmt.Println(hashFunction("shelf"))
	loadWords()
}
