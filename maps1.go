package main

import (
	"fmt"
	"strings"
)

func countWords(text string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	text := "Olá, isso é um exemplo. Exemplo de contagem de palavras em uma string."

	wordCount := countWords(text)

	for word, count := range wordCount {
		fmt.Println(word, count)
	}
}
