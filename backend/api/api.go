package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	uuid "github.com/google/uuid"
)

type application struct {
	Addr string
}

var users []User

func ValidateStruct(user User) error {
	validate := validator.New()
	return validate.Struct(user)
}

func (app *application) handleAddUsers(w http.ResponseWriter, r *http.Request) {
	var user User
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}
	// validator := validator.New()
	// if err := validator.Struct(user); err != nil {
	// 	http.Error(w, "Invalid credentials", http.StatusBadRequest)
	// 	return
	// }
	if err := ValidateStruct(user); err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	for _, usr := range users {
		if usr.Email == user.Email {
			http.Error(w, "This Email Is Already Used!", http.StatusBadRequest)
			return
		}
	}

	user.UserId = uuid.New()
	users = append(users, user)
	fmt.Printf("Received: %+v\n", user)
	w.Write([]byte("Received"))
}

func (app *application) handleGetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Sending users", users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
func (app *application) handlePatch(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid Credentials", http.StatusBadRequest)
	}
	if err := ValidateStruct(user); err != nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
		return
	}

	idStr := r.PathValue("userId")
	id, err := uuid.Parse(idStr)

	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	for i, usr := range users {
		if id == usr.UserId {
			user.UserId = id
			users[i] = user
			w.WriteHeader(http.StatusOK) // Set HTTP 200 OK
			return
		}
	}

}

func (app *application) handleDelete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("userId")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	for i, usr := range users {
		if usr.UserId == id {
			users = append(users[:i], users[i+1:]...)
		}

	}
	w.WriteHeader(http.StatusOK) // Set HTTP 200 OK
}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users", app.handleGetUsers)
	mux.HandleFunc("POST /users", app.handleAddUsers)
	mux.HandleFunc("PATCH /users/{userId}", app.handlePatch)
	mux.HandleFunc("DELETE /users/{userId}", app.handleDelete)

	return mux
}

func (app *application) run(mux *http.ServeMux) error {
	srv := &http.Server{
		Addr:    app.Addr,
		Handler: mux,
	}

	return srv.ListenAndServe()
}
