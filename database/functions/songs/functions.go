package songs

import (
	"github.com/JonayMedina/api-music-db/database"
	"github.com/JonayMedina/api-music-db/database/structs"
)

func GetSong(id int) (*structs.Song, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetSong(id)
}

func GetSongs() ([]*structs.Song, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetSongs()
}

func CreateSong(song *structs.Song) (int, error) {
	if err := database.CheckDB(); err != nil {
		return 0, err
	}
	return database.DbServer.CreateSong(song)
}

func UpdateSong(song *structs.Song) (int, error) {
	if err := database.CheckDB(); err != nil {
		return 0, err
	}
	return database.DbServer.UpdateSong(song)
}

func DeleteSong(id int) (int, error) {
	if err := database.CheckDB(); err != nil {
		return 0, err
	}
	return database.DbServer.DeleteSong(id)
}
