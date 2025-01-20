package mysql

import (
	"database/sql"
	"log"

	"github.com/JonayMedina/api-music-db/database/structs"
)

func (db DBServer) GetUser(userID int) (*structs.User, error) {
	return getUser(db.DB, userID)
}

func getUser(db *sql.DB, userID int) (*structs.User, error) {
	user := &structs.User{}

	rows, err := db.Query(`SELECT 
		u.id,
		u.username,
		u.email,
		u.password,
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
			&user.Password,
			&user.CreatedAt,
		)
		if err != nil {
			log.Println("error al obtener el usuario", err)
			return nil, err
		}
	}
	return user, nil
}
