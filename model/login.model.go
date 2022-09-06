package model

type Login struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type LoginForm struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
