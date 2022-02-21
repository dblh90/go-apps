package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

	//slice := []int{1, 2, 3}
	wellKnownPorts := map[string]int{"http": 80, "https": 81}

	for _, v := range wellKnownPorts {
		println(v)
	}
}
