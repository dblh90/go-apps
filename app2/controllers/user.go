package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPatter *regexp.Regexp
}

func newUserController() *userController {
	return &userController{
		userIDPatter: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}

func (uc userController) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("Hello from User Controller!"))
}
