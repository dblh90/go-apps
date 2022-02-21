package main

import (
	"net/http"

	"hamza.com/goapps/app2/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":4000", nil)
}
