package users

import (
	"github.com/JonayMedina/api-music-db/database"
	"github.com/JonayMedina/api-music-db/database/structs"
)

func GetUser(id int) (*structs.User, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetUser(id)
}

func GetUsers() ([]*structs.User, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetUsers()
}

func CreateUser(user *structs.User) (*structs.User, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.CreateUser(user)
}
