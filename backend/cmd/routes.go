package main

import (
	"net/http"

	"github.com/Rha02/carpool_app/handlers"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

func routes(key string) http.Handler {
	router := mux.NewRouter()

	// Middleware
	csrfMiddleware := csrf.Protect([]byte(key), csrf.Secure(false))
	router.Use(csrfMiddleware)

	// CSRF Token
	router.HandleFunc("/get-token", handlers.Repo.GetToken).Methods("GET")

	// Authentication
	router.HandleFunc("/login", handlers.Repo.Login).Methods("POST")
	router.HandleFunc("/logout", handlers.Repo.Logout).Methods("POST")
	router.HandleFunc("/register", handlers.Repo.Register).Methods("POST")
	router.HandleFunc("/checkauth", handlers.Repo.CheckAuth).Methods("GET")

	// Users
	router.HandleFunc("/users", handlers.Repo.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.Repo.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.Repo.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{id}", handlers.Repo.UpdateUser).Methods("PUT", "PATCH")

	// Threads
	router.HandleFunc("/threads", handlers.Repo.GetAllThreads).Methods("GET")
	router.HandleFunc("/users/{u_id}/threads", handlers.Repo.GetUserThreads).Methods("GET")
	router.HandleFunc("/threads/{id}", handlers.Repo.GetThread).Methods("GET")
	router.HandleFunc("/threads", handlers.Repo.PostThread).Methods("POST")
	router.HandleFunc("/threads/{id}", handlers.Repo.DeleteThread).Methods("DELETE")
	router.HandleFunc("/threads/{id}", handlers.Repo.UpdateThread).Methods("PUT", "PATCH")

	return router
}
