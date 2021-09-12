package main

var rev int

func reverseDigit(num int) int {
	if num == 0 {
		return 0
	}
	rev = rev*10 + num%10
	reverseDigit(num / 10)
	return rev
}

func reverseNumber(num int) int {

	return 0
}

func getDigit(num int) int {
	return 0
}
