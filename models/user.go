package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// albums slice to seed record album data.
var Users = []User{
	{ID: "1", Name: "John", Email: "", Password: ""},
	{ID: "1", Name: "John", Email: "", Password: ""},
	{ID: "1", Name: "John", Email: "", Password: ""}}
