package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JonayMedina/api-music-db/database/structs"
)

func (db DBServer) GetUser(userID int) (*structs.User, error) {
	return getUser(db.DB, userID)
}

func (db DBServer) GetUsers() ([]*structs.User, error) {
	return getUsers(db.DB)
}

func (db DBServer) CreateUser(user *structs.User) (*structs.User, error) {
	return createUser(db.DB, user)
}

func getUsers(db *sql.DB) ([]*structs.User, error) {
	users := []*structs.User{}

	rows, err := db.Query(`SELECT * FROM users`)
	if err != nil {
		log.Println("error al obtener los usuarios", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		user := &structs.User{}
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			log.Println("error al obtener el usuario", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func getUser(db *sql.DB, userID int) (*structs.User, error) {
	user := &structs.User{}

	rows, err := db.Query(`SELECT 
		u.id,
		u.username,
		u.email,
		u.created_at
	FROM users u WHERE u.id = ?`, userID)

	if err != nil {
		log.Println("error al obtener el usuario", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		)
		if err != nil {
			log.Println("error al obtener el usuario", err)
			return nil, err
		}
	}
	return user, nil
}

func createUser(db *sql.DB, user *structs.User) (*structs.User, error) {
	fmt.Printf("mysql createUser %+v\n", user)

	query := `INSERT INTO users (username, email, password, created_at) VALUES (?, ?, ?, ?)`
	user.HashPassword()

	result, err := db.Exec(query, user.Username, user.Email, user.Password, getNowDateTime())
	if err != nil {
		log.Println("error al crear el usuario", err)
		return nil, err
	}
	fmt.Println("result", result)
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("error al obtener el id del usuario", err)
		return nil, err
	}
	user.ID = int(id)
	return user, nil
}

func (db DBServer) GetUserByEmail(email string) (*structs.User, error) {
	user := &structs.User{}

	rows, err := db.DB.Query(`SELECT
		u.id,
		u.username,
		u.email,
		u.created_at
	FROM users u WHERE u.email = ? LIMIT 1`, email)
	if err != nil {
		log.Println("error al obtener el usuario por email", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Println("error al obtener el usuario", err)
			return nil, err
		}
	}
	return user, nil
}
