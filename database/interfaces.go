package database

import (
	"github.com/JonayMedina/api-music-db/database/structs"
)

type DBInterface interface {
	GetUser(id int) (*structs.User, error)
	GetUsers() ([]*structs.User, error)
	CreateUser(user *structs.User) (int, error)

	GetRole(id int) (*structs.Role, error)
	GetRoles() ([]*structs.Role, error)
	CreateRole(role *structs.Role) (int, error)

	GetArtist(id int) (*structs.Artist, error)
	GetArtists() ([]*structs.Artist, error)
	CreateArtist(artist *structs.Artist) (int, error)

	GetSong(id int) (*structs.Song, error)
	GetSongs() ([]*structs.Song, error)
	CreateSong(song *structs.Song) (int, error)
	UpdateSong(song *structs.Song) (int, error)
	DeleteSong(id int) (int, error)

	GetUserLikesSong(userID, songID int) (*structs.UserLikesSong, error)
	GetUserLikesSongs(userID int) ([]*structs.UserLikesSong, error)
	CreateUserLikesSong(userID, songID int) (int, error)

	GetUserPlaylist(userID, playlistID int) (*structs.UserPlaylist, error)
	GetUserPlaylists(userID int) ([]*structs.UserPlaylist, error)
	CreateUserPlaylist(userID, playlistID int) (int, error)

	GetPlaylist(playlistID int) (*structs.Playlist, error)
	GetPlaylists() ([]*structs.Playlist, error)
	CreatePlaylist(playlist *structs.Playlist) (int, error)
	UpdatePlaylist(playlist *structs.Playlist) (int, error)
	DeletePlaylist(playlistID int) (int, error)
}
