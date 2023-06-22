package main

func countNumberPairs(numbers []int) map[[2]int]int {
	pairCount := make(map[[2]int]int)

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			pair := [2]int{numbers[i], numbers[j]}
			pairCount[pair]++
		}
	}

	return pairCount
}
