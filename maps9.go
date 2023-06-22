package main

func generateFibonacciSequence(n int) map[int]int {
	fibonacci := make(map[int]int)
	fibonacci[0] = 0
	fibonacci[1] = 1

	for i := 2; fibonacci[i-1] <= n; i++ {
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	return fibonacci
}
