package main

import "strings"

var res []string

func capitalizeFirst(str []string) []string {
	if len(str) == 0 {
		return str
	}

	res = append(res, strings.Title(str[0]))
	capitalizeFirst(str[len(str)-(len(str)-1):])
	return res
}
