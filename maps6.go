package main

func sumWordCounts(wordCounts []map[string]int) map[string]int {
	totalWordCount := make(map[string]int)

	for _, wordCount := range wordCounts {
		for word, count := range wordCount {
			totalWordCount[word] += count
		}
	}

	return totalWordCount
}
