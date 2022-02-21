package main

import (
	"fmt"

	"hamza.com/goapps/app2/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "User1- FN",
		LastName:  "User1- LN",
	}
	fmt.Println(u)
}
