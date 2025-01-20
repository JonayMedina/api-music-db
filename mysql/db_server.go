package mysql

import (
	"database/sql"
)

type DBServer struct {
	DB *sql.DB
}
