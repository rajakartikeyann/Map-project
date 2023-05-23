package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	// r.Host("http://localhost:5000/")
	// r.Methods("GET", "POST")

	// cors := handlers.CORS(
	// 	handlers.AllowedOrigins([]string{"*"}),
	// 	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	// 	handlers.AllowedHeaders([]string{"Content-Type"}),
	// )
	// r.Use(cors)

	r.HandleFunc("/login", Login).Methods("POST")
	fmt.Println("=======Project starting=======")
	http.ListenAndServe(":4300", r)

}
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("INSIDE LOGIN")
	vars := mux.Vars(r)
	fmt.Println(vars["name"], vars["page"]) // the page

	byte, _ := ioutil.ReadAll(r.Body)
	fmt.Println("****************", string(byte))
}
