package main

import (
	"encoding/json"
	"fmt"
)

type Worker struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type Data struct {
	UserName string    `json:"user_name"`
	Worker   []*Worker `json:"workers,omitempty"`
}

func main1() {
	w := Data{
		UserName: "subodhqss",
	}

	list := make([]*Worker, 0, 5)

	fmt.Println(list)

	b, _ := json.MarshalIndent(w, " ", " ")

	fmt.Println(string(b))
}
