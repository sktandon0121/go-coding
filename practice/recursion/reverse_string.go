package main

func reverseString(str string) string {
	if len(str) == 0 || len(str) == 1 {
		return str
	}
	return reverseString(str[1:]) + str[:1]
}
