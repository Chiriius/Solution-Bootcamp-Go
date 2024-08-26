package models

type User struct {
	ID          string `json:"id" bson:"id"`
	Password    string `json:"password" bson:"password"`
	Age         string `json:"age" bson:"age"`
	Information string `json:"information" bson:"information"`
	Parents     string `json:"parents" bson:"parents"`
	Email       string `json:"email" bson:"email"`
	Name        string `json:"name" bson:"name"`
}
