package model

type User struct {
	Id   int
	Name string
}

var users = []User{
	{Id: 1, Name: "Alice"},
	{Id: 2, Name: "Bob"},
}

func GetName() []User {
	return users
}
