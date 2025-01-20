package database

import (
	"errors"
	"log"
)

var DbServer DBInterface

func SetDB(db DBInterface) {
	DbServer = db
}

func CheckDB() error {
	log.Println("checkDB DbServer", DbServer)
	if DbServer == nil {
		return errors.New("db is nil")
	}
	return nil
}
