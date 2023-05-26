package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"task/login/dto"

	log "github.com/sirupsen/logrus"
)

var Ctx context.Context

func Login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("INSIDE LOGIN")
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Request Body Read Error :", err)
	}

	var customer dto.Customer
	json.Unmarshal(bytes, &customer)
	fmt.Println(customer)

	if customer.Name == "admin@map.com" {
		if customer.Password == "Test@1234" {
			json.NewEncoder(w).Encode("Login Successful")
		} else {
			json.NewEncoder(w).Encode("Password is incorrect")
		}
	} else {
		json.NewEncoder(w).Encode("Username was not found, Please register now")
	}

	// query := db.M{
	// 	customer.Name :"admin@map.com",
	// 	customer.Password :"Test@1234"
	// }

	// res :=db.Collec.FindOne(Ctx, db.M{})
	//Logincheck()
}

//  func Logincheck(){

//  }
