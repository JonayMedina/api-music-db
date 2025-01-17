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
	users, err := database.DbServer.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func CreateUser(user *structs.User) (int, error) {
	if err := database.CheckDB(); err != nil {
		return 0, err
	}
	return database.DbServer.CreateUser(user)
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
func CreateUserLikesSong(userID, songID int) (int, error) {
	if err := database.CheckDB(); err != nil {
		return 0, err
	}
	return database.DbServer.CreateUserLikesSong(userID, songID)
}

func GetUserPlaylist(userID, playlistID int) (*structs.UserPlaylist, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetUserPlaylist(userID, playlistID)
}
func GetUserPlaylists(userID int) ([]*structs.UserPlaylist, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetUserPlaylists(userID)
}
func CreateUserPlaylist(userID, playlistID int) (int, error) {
	if err := database.CheckDB(); err != nil {
		return 0, err
	}
	return database.DbServer.CreateUserPlaylist(userID, playlistID)
}

func GetPlaylists() ([]*structs.Playlist, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetPlaylists()
}

func CreatePlaylist(playlist *structs.Playlist) (int, error) {
	if err := database.CheckDB(); err != nil {
		return 0, err
	}
	return database.DbServer.CreatePlaylist(playlist)
}
func UpdatePlaylist(playlist *structs.Playlist) (int, error) {
	if err := database.CheckDB(); err != nil {
		return 0, err
	}
	return database.DbServer.UpdatePlaylist(playlist)
}
func DeletePlaylist(playlistID int) (int, error) {
	if err := database.CheckDB(); err != nil {
		return 0, err
	}
	return database.DbServer.DeletePlaylist(playlistID)
}
