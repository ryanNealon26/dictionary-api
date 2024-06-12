package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Word struct {
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
func handleCollision(hashMap map[int][]Word, searchWord string) Word {
	mapSize := len(hashMap[hashFunction(searchWord)])
	for i := 0; i < mapSize; i++ {
		if searchWord == hashMap[hashFunction(searchWord)][i].word {
			return hashMap[hashFunction(searchWord)][i]
		}
	}
	err := Word{searchWord, "Search Word: " + searchWord + " was not found in the dictionary."}
	return err

}
func loadDictionary() map[int][]Word {
	var dictMap = make(map[int][]Word)
	file, err := os.Open("Dictionary-2.txt")
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 2 && strings.Contains(text, "  ") {
			end := strings.Index(text, "  ")
			wordKey := Word{text[0:end], text[2+end:]}
			dictMap[hashFunction(wordKey.word)] = append(dictMap[hashFunction(wordKey.word)], wordKey)
		}
	}
	return dictMap
}
