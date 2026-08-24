package main

import (
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type UserAPI struct {
	users  map[int]User
	nextID int
}

func NewUserAPI() *UserAPI {
	return &UserAPI{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (a *UserAPI) Routes() http.Handler {
	mux := http.NewServeMux()

	return mux
}

func main() {
	api := NewUserAPI()

	log.Println("server listening on http://localhost:8080")

	if err := http.ListenAndServe(":8080", api.Routes()); err != nil {
		log.Fatal(err)
	}
}
