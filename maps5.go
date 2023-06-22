package main

func characterFrequency(str string) map[rune]int {
	frequencyMap := make(map[rune]int)

	for _, char := range str {
		frequencyMap[char]++
	}

	return frequencyMap
}
