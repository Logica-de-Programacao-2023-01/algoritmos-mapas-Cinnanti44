package main

import "fmt"

func sumValues(m map[string]int) int {
	sum := 0
	for _, value := range m {
		sum += value
	}
	return sum
}

func main() {
	values := map[string]int{
		"a": 10,
		"b": 20,
		"c": 30,
	}

	total := sumValues(values)
	fmt.Println("Soma dos valores:", total)
}
