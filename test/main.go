package main

import (
	"fmt"
	"time"
)

type User struct {
	Name      string
	Age       int
	StartDate *time.Time
}

func main() {
	d := time.Date(2021, 3, 23, 0, 0, 0, 0, time.UTC)
	user := User{StartDate: &d}
	currentMonth :=3
	fmt.Println(user.StartDate.Month() == 3)
	fmt.Println(user.StartDate.Month() == time.Month(currentMonth))
}
