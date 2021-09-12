package main

// Return nth element of fibonacci sequence

func fib(nth int) int {
	if nth <= 2 {
		return 1
	}

	return fib(nth-2) + fib(nth-1)
}
