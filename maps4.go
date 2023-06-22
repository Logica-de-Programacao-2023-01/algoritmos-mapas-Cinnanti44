package main

import (
	"sort"
	"strings"
)

func findAnagrams(wordSlice []string) map[string][]string {
	anagramMap := make(map[string][]string)

	for _, word := range wordSlice {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	return anagramMap
}

func sortString(word string) string {
	s := strings.Split(word, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
