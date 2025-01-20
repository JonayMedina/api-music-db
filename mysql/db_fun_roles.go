package mysql

import (
	"database/sql"
	"log"

	"github.com/JonayMedina/api-music-db/database/structs"
)

func (db DBServer) GetRoles() ([]*structs.Role, error) {
	return getRoles(db.DB)
}

func (db DBServer) GetRole(roleID int) (*structs.Role, error) {
	return getRole(db.DB, roleID)
}

func (db DBServer) CreateRole(role *structs.Role) (*structs.Role, error) {
	return createRole(db.DB, role)
}

func (db DBServer) UpdateRole(role *structs.Role) (*structs.Role, error) {
	return updateRole(db.DB, role)
}

func (db DBServer) DeleteRole(roleID int) (int, error) {
	return deleteRole(db.DB, roleID)
}

func updateRole(db *sql.DB, role *structs.Role) (*structs.Role, error) {
	query := `UPDATE roles SET name = ? WHERE id = ?`
	_, err := db.Exec(query, role.Name, role.ID)
	if err != nil {
		log.Println("error al actualizar el rol", err)
		return nil, err
	}
	return role, nil
}

func getRoles(db *sql.DB) ([]*structs.Role, error) {
	roles := []*structs.Role{}

	rows, err := db.Query(`SELECT id, name FROM roles`)
	if err != nil {
		log.Println("error al obtener los roles", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		role := &structs.Role{}
		err := rows.Scan(&role.ID, &role.Name)
		if err != nil {
			log.Println("error al obtener los roles", err)
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func getRole(db *sql.DB, roleID int) (*structs.Role, error) {
	role := &structs.Role{}

	row := db.QueryRow(`SELECT id, name FROM roles WHERE id = ?`, roleID)
	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		log.Println("error al obtener el rol", err)
		return nil, err
	}
	return role, nil
}

func createRole(db *sql.DB, role *structs.Role) (*structs.Role, error) {
	query := `INSERT INTO roles (name) VALUES (?)`
	res, err := db.Exec(query, role.Name)
	if err != nil {
		log.Println("error al crear el rol", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("error al obtener el id del rol", err)
		return nil, err
	}

	role.ID = int(id)
	return role, nil
}

func deleteRole(db *sql.DB, roleID int) (int, error) {
	query := `DELETE FROM roles WHERE id = ?`
	_, err := db.Exec(query, roleID)
	if err != nil {
		log.Println("error al eliminar el rol", err)
		return 0, err
	}
	return roleID, nil
}
