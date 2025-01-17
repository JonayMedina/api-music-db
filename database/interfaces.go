package database

import (
	"github.com/JonayMedina/api-music-db/database/structs"
)

type DBInterface interface {
	GetUser(id int) (*structs.User, error)
	GetUsers() ([]*structs.User, error)
	CreateUser(user *structs.User) (*structs.User, error)

	GetRole(id int) (*structs.Role, error)
	GetRoles() ([]*structs.Role, error)
	CreateRole(role *structs.Role) (*structs.Role, error)
	UpdateRole(role *structs.Role) (*structs.Role, error)
	DeleteRole(id int) (int, error)

	FindArtist(id int) (*structs.Artist, error)
	GetArtists() ([]*structs.Artist, error)
	CreateArtist(artist *structs.Artist) (*structs.Artist, error)

	GetSong(id int) (*structs.Song, error)
	GetSongs() ([]*structs.Song, error)
	CreateSong(song *structs.Song) (*structs.Song, error)
	UpdateSong(song *structs.Song) (*structs.Song, error)
	DeleteSong(id int) error

	GetUserLikesSong(userID, songID int) (*structs.UserLikesSong, error)
	GetUserLikesSongs(userID int) ([]*structs.UserLikesSong, error)
	CreateUserLikesSong(userID, songID int) (*structs.UserLikesSong, error)

	FindPlaylist(userID, playlistID int) (*structs.UserPlaylist, error)
	GetUserPlaylists(userID int) ([]*structs.UserPlaylist, error)
	CreateUserPlaylist(userID, playlistID int) (*structs.UserPlaylist, error)

	GetPlaylists() ([]*structs.Playlist, error)
	CreatePlaylist(playlist *structs.Playlist) (*structs.Playlist, error)
	UpdatePlaylist(playlist *structs.Playlist) (*structs.Playlist, error)
	DeletePlaylist(playlistID int) (int, error)
}
