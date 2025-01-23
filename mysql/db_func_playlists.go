package mysql

import (
	"database/sql"
	"log"

	"github.com/JonayMedina/api-music-db/database/structs"
)

func (db DBServer) CreateUserPlaylist(userID, playlistID int) error {
	return createUserPlaylist(db.DB, userID, playlistID)
}

func (db DBServer) GetUserPlaylists(userID int) ([]*structs.Playlist, error) {
	return getUserPlaylists(db.DB, userID)
}

func (db DBServer) FindPlaylist(userID, playlistID int) (*structs.UserPlaylist, error) {
	return findPlaylist(db.DB, userID, playlistID)
}

func (db DBServer) UpdatePlaylist(playlist *structs.Playlist) (*structs.Playlist, error) {
	return updatePlaylist(db.DB, playlist)
}

func (db DBServer) DeletePlaylist(playlistID int) (int, error) {
	return deletePlaylist(db.DB, playlistID)
}

func (db DBServer) CreatePlaylist(playlist *structs.Playlist) (*structs.Playlist, error) {
	return createPlaylist(db.DB, playlist)
}

func createUserPlaylist(db *sql.DB, userID, playlistID int) error {
	query := `INSERT INTO users_playlists (user_id, playlist_id, created_at) VALUES (?, ?, ?)`
	_, err := db.Exec(query, userID, playlistID, getNowDateTime())
	if err != nil {
		log.Println("error al crear el playlist del usuario", err)
		return err
	}

	return nil
}

func createPlaylist(db *sql.DB, playlist *structs.Playlist) (*structs.Playlist, error) {
	query := `INSERT INTO playlists (name, created_at) VALUES (?, ?)`
	_, err := db.Exec(query, playlist.Name, getNowDateTime())
	if err != nil {
		log.Println("error al crear el playlist", err)
		return nil, err
	}
	return playlist, nil
}

func findPlaylist(db *sql.DB, userID, playlistID int) (*structs.UserPlaylist, error) {

	query := `SELECT
		up.user_id,
		up.created_at
		p.id,
		p.name
	FROM user_playlists up
	INNER JOIN playlists p ON p.id = up.playlist_id
	WHERE up.user_id = ?
	AND up.playlist_id = ?`
	rows, err := db.Query(query, userID, playlistID)
	if err != nil {
		log.Println("error al obtener el playlist del usuario", err)
		return nil, err
	}
	defer rows.Close()

	userPlaylist := &structs.UserPlaylist{}
	playlistIDs := []int{playlistID}
	for rows.Next() {
		err := rows.Scan(
			&userPlaylist.UserID,
			&userPlaylist.CreatedAt,
			&userPlaylist.PlaylistID,
			&userPlaylist.Playlist.Name,
		)
		if err != nil {
			log.Println("error al obtener el playlist del usuario", err)
			continue
		}
		songs, err := getPlaylistsSongs(db, playlistIDs)
		if err != nil {
			log.Println("error al obtener listado de canciones de las playlists", err)
			continue
		}
		userPlaylist.Playlist.Songs = songs
	}

	return userPlaylist, nil
}

func getUserPlaylists(db *sql.DB, userID int) ([]*structs.Playlist, error) {
	playlists := []*structs.Playlist{}

	query := `SELECT
		p.id,
		p.name,
		up.user_id,
		up.created_at,
	FROM playlists p
	WITH (NOLOCK)
	INNER JOIN users_playlists up WITH (NOLOCK) ON up.playlist_id = p.id
	WHERE up.user_id = ?`

	rows, err := db.Query(query, userID)
	if err != nil {
		log.Println("error al obtener listado de playlists del usuario", err)
		return nil, err
	}

	defer rows.Close()
	playlistIDs := []int{}

	for rows.Next() {
		playlist := &structs.Playlist{}
		err := rows.Scan(
			&playlist.ID,
			&playlist.Name,
		)
		if err != nil {
			log.Println("error al obtener listado de playlists del usuario", err)
			continue
		}
		playlists = append(playlists, playlist)
		playlistIDs = append(playlistIDs, playlist.ID)
	}

	songs, err := getPlaylistsSongs(db, playlistIDs)
	if err != nil {
		log.Println("error al obtener listado de canciones de las playlists", err)
		return nil, err
	}

	for _, playlist := range playlists {
		for _, song := range songs {
			if playlist.ID == song.PlaylistID {
				playlist.Songs = append(playlist.Songs, song)
			}
		}
	}
	return playlists, nil
}

func updatePlaylist(db *sql.DB, playlist *structs.Playlist) (*structs.Playlist, error) {
	query := `UPDATE playlists SET name = ? WHERE id = ?`
	_, err := db.Exec(query, playlist.Name, playlist.ID)
	if err != nil {
		log.Println("error al actualizar el playlist", err)
		return nil, err
	}
	return playlist, nil
}

func deletePlaylist(db *sql.DB, playlistID int) (int, error) {
	query := `DELETE FROM playlists WHERE id = ?`
	_, err := db.Exec(query, playlistID)
	if err != nil {
		log.Println("error al eliminar el playlist", err)
		return 0, err
	}
	return playlistID, nil
}
