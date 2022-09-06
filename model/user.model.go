package model

type User struct {
	Id       int    `json:"id" bson:"id"`
	Username string `json:"username" bson:"username"`
	Age      int    `json:"age" bson:"age"`
}
