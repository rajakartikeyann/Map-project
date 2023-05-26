package dto

type Customer struct {
	Name     string `json:"name,omitempty" bson:"name,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}

type Request struct {
}

type Response struct {
}
