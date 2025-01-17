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

func CreateSong(song *structs.Song) (*structs.Song, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.CreateSong(song)
}

func UpdateSong(song *structs.Song) (*structs.Song, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.UpdateSong(song)
}

func DeleteSong(id int) error {
	if err := database.CheckDB(); err != nil {
		return err
	}
	return database.DbServer.DeleteSong(id)
}

func GetUserLikesSong(userID, songID int) (*structs.UserLikesSong, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetUserLikesSong(userID, songID)
}
func GetUserLikesSongs(userID int) ([]*structs.UserLikesSong, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetUserLikesSongs(userID)
}
func CreateUserLikesSong(userID, songID int) (*structs.UserLikesSong, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.CreateUserLikesSong(userID, songID)
}

func FindArtist(id int) (*structs.Artist, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.FindArtist(id)
}

func GetArtists() ([]*structs.Artist, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetArtists()
}

func CreateArtist(artist *structs.Artist) (*structs.Artist, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.CreateArtist(artist)
}
