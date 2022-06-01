package model

// Representation of user data
type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

// Populate user data
var Users = []User{
	{
		Id:       1,
		Name:     "Administrator",
		Email:    "admin@mail.com",
		Password: "$2y$14$3ubGx5XApI3tmnbedjdiFOxKUqc8nULV6M/DNZOhuY8lYcCRMjlkC", // secret
	},
	{
		Id:       2,
		Name:     "Muhammad Rizki Maulidan",
		Email:    "mrizkimaulidan@mail.com",
		Password: "$2y$14$3ubGx5XApI3tmnbedjdiFOxKUqc8nULV6M/DNZOhuY8lYcCRMjlkC", // secret
	},
}
