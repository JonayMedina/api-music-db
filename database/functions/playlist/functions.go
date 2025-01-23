package playlist

import (
	"github.com/JonayMedina/api-music-db/database"
	"github.com/JonayMedina/api-music-db/database/structs"
)

func FindPlaylist(userID, playlistID int) (*structs.UserPlaylist, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.FindPlaylist(userID, playlistID)
}
func GetUserPlaylists(userID int) ([]*structs.UserPlaylist, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetUserPlaylists(userID)
}
func CreateUserPlaylist(userID, playlistID int) error {
	if err := database.CheckDB(); err != nil {
		return err
	}
	return database.DbServer.CreateUserPlaylist(userID, playlistID)
}

func GetPlaylists() ([]*structs.Playlist, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.GetPlaylists()
}

func CreatePlaylist(playlist *structs.Playlist) (*structs.Playlist, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.CreatePlaylist(playlist)
}
func UpdatePlaylist(playlist *structs.Playlist) (*structs.Playlist, error) {
	if err := database.CheckDB(); err != nil {
		return nil, err
	}
	return database.DbServer.UpdatePlaylist(playlist)
}
func DeletePlaylist(playlistID int) (int, error) {
	if err := database.CheckDB(); err != nil {
		return 0, err
	}
	return database.DbServer.DeletePlaylist(playlistID)
}
