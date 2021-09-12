package main

import "fmt"

var memo = make(map[int]int)
var i = 0

func fibDP(nth int) int {

	if nth <= 2 {
		return 1
	}

	if memo[nth-1] == 0 && memo[nth-2] == 0 {
		fmt.Println("count of calls ", i)
		i++
		memo[nth-1] = fibDP(nth - 1)
		memo[nth-2] = fibDP(nth - 2)
	}

	return memo[nth-1] + memo[nth-2]
}
