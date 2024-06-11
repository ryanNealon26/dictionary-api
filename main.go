package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type word struct {
	word       string
	definition string
}

func hashFunction(word string) int {
	runeArray := []rune(word)
	hashKey := 0
	for i := 0; i < len(runeArray); i++ {
		hashKey += int(runeArray[i])
	}
	return hashKey
}
func handleCollision(hashMap map[int][]word, searchWord string) string {
	mapSize := len(hashMap[hashFunction(searchWord)])
	for i := 0; i < mapSize; i++ {
		if searchWord == hashMap[hashFunction(searchWord)][i].word {
			return hashMap[hashFunction(searchWord)][i].definition
		}
	}
	return "Search Word: " + searchWord + " was not found in the dictionary."

}
func loadWords() {
	var dictMap = make(map[int][]word)
	file, err := os.Open("Dictionary-2.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 2 && strings.Contains(text, "  ") {
			end := strings.Index(text, "  ")
			wordKey := word{text[0:end], text[end:]}
			dictMap[hashFunction(wordKey.word)] = append(dictMap[hashFunction(wordKey.word)], wordKey)
		}
	}
}
func main() {
	fmt.Println(hashFunction("shelf"))
	loadWords()
}
