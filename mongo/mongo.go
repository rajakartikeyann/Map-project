package mongo

import (
	"context"
	"fmt"
	dto "task/login/dto"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client
var C map[string]*mongo.Collection
var Db *mongo.Database
var Collec *mongo.Collection

// Mongo DB server Connect
func MongoConnect() error {

	Ctx := context.Background()
	defer Ctx.Done()

	var err error
	Client, err = mongo.Connect(Ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}

	err = Client.Ping(Ctx, readpref.Primary())
	if err != nil {
		return err
	}
	fmt.Println("MongoDB Connected!")
	LoginInit()
	return nil
}

// Initial UserLogin insert in DB for initial access

func LoginInit() {
	Db = Client.Database("Customers")
	Collec = Db.Collection("LoginData")
	log.Println("Initialising Default login...")

	person := dto.Customer{}
	person.Name = "admin@map.com"
	person.Password = "Test@1234"

	_, err := Collec.InsertOne(context.Background(), person)
	if err != nil {
		fmt.Println("Default Login insert Error :", err)
	}

}
