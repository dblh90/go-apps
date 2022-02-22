package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"hamza.com/goapps/app2/models"
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

	if req.URL.Path == "/users" {
		switch req.Method {

		case http.MethodGet:
			uc.getAll(writer, req)
		case http.MethodPost:
			uc.post(writer, req)

		default:
			writer.WriteHeader(http.StatusNotImplemented)
		}

	} else {

	}

	writer.Write([]byte("Hello from User Controller!"))
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	// Encode Response as JSON
}

func (uc *userController) get(id int, w http.ResponseWriter) {

	_, err := models.GetUserById(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	//	encodeReponseAsJSON()
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	userRequest, err := uc.parseRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}

	addedUser, err := models.AddUser(userRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	encodeReponseAsJSON(addedUser, w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	userRequest, err := uc.parseRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}

	if id != userRequest.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}

	updatedUser, err := models.UpdateUser(userRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error))
		return
	}
	encodeReponseAsJSON(updatedUser, w)
}

func (uc *userController) parseRequest(request *http.Request) (models.User, error) {

	dec := json.NewDecoder(request.Body)
	var user models.User
	err := dec.Decode(&user)

	if err != nil {
		return models.User{}, err
	}

	return user, nil

}
