package main

import (
	"fmt"
	"net/http"
	srv "task/login/service"
	db "task/mongo"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("=======Project On Live=======")
	r := mux.NewRouter()
	err := db.MongoConnect()
	if err != nil {
		fmt.Println("Mongo server not Connected :", err)
	}
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)
	r.Use(cors)
	r.HandleFunc("/login", srv.Login)
	http.ListenAndServe(":4300", r)
}
