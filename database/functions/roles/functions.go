package roles

import (
	"github.com/JonayMedina/api-music-db/database"
	"github.com/JonayMedina/api-music-db/database/structs"
)

func GetRole(id int) (*structs.Role, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetRole(id)
}

func GetRoles() ([]*structs.Role, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetRoles()
}

func CreateRole(role *structs.Role) (*structs.Role, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.CreateRole(role)
}

func UpdateRole(role *structs.Role) (*structs.Role, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.UpdateRole(role)
}

func DeleteRole(id int) (int, error) {
	if err := database.CheckDB(); err != nil {
		return 0, err
	}
	return database.DbServer.DeleteRole(id)
}
