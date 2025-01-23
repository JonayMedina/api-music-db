package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/JonayMedina/api-music-db/database/structs"
)

func (db DBServer) GetSongs() ([]*structs.Song, error) {
	return getSongs(db.DB)
}

func (db DBServer) GetSong(songID int) (*structs.Song, error) {
	return getSong(db.DB, songID)
}

func (db DBServer) SearchSongs(request, artist, album string, page, limit int) ([]*structs.Song, int64, error) {
	sqlQuery := "SELECT * FROM songs WHERE "

	// Slice para almacenar los argumentos
	args := make([]interface{}, 0)

	if request != "" {
		sqlQuery += " title LIKE ?"
		args = append(args, "%"+request+"%")
	}
	if artist != "" {
		sqlQuery += " AND artist_id = ?"
		args = append(args, artist)
	}
	if album != "" {
		sqlQuery += " AND album = ?"
		args = append(args, album)
	}

	limit = getLimit(limit)
	skip := getSkip(page, limit)

	sqlQuery += " LIMIT ? OFFSET ?"
	args = append(args, limit, skip)
	fmt.Println(" sqlQuery on search songs", sqlQuery, args)
	// Pasar el slice de argumentos usando args...
	res, err := db.DB.Query(sqlQuery, args...)
	if err != nil {
		log.Println("error al obtener listado de canciones", err)
		return nil, 0, err
	}
	defer res.Close()

	songs := []*structs.Song{}
	for res.Next() {
		song := &structs.Song{}
		err := res.Scan(
			&song.ID,
			&song.Title,
			&song.ArtistID,
			&song.Duration,
			&song.Album,
			&song.Genre,
			&song.ReleaseDate,
			&song.CoverImage,
			&song.CreatedAt,
		)
		if err != nil {
			log.Println("error al obtener listado de canciones", err)
			continue
		}
		songs = append(songs, song)
	}

	total, err := countSongs(db.DB, args, limit)
	if err != nil {
		log.Println("error al contar canciones", err)
		return nil, 0, err
	}

	return songs, int64(total), nil
}

func (db DBServer) GetSongByTitle(title string) (*structs.Song, error) {
	return getSongByTitle(db.DB, title)
}

func (db DBServer) GetSongByArtist(artistID int) ([]*structs.Song, error) {
	return getSongByArtist(db.DB, artistID)
}

func (db DBServer) CreateSong(song *structs.Song) (*structs.Song, error) {
	return createSong(db.DB, song)
}

func (db DBServer) UpdateSong(song *structs.Song) (*structs.Song, error) {
	return updateSong(db.DB, song)
}

func (db DBServer) DeleteSong(songID int) error {
	return deleteSong(db.DB, songID)
}

func (db DBServer) GetUserLikesSong(userID, songID int) (*structs.UserLikesSong, error) {
	return getUserLikesSong(db.DB, userID, songID)
}

func (db DBServer) GetAllUserSongs(userID int) ([]*structs.UserLikesSong, error) {
	UsersLikesSongs := []*structs.UserLikesSong{}

	query := `SELECT * 
			uls.user_id,
			uls.song_id,
			uls.created_at,
			s.id,
			s.title,
			s.artist_id,
			s.album,
			s.duration,
			s.genre,
			s.release_date,
			s.cover_image,
			s.created_at
	FROM users_likes_songs uls 
	INNER JOIN songs s ON s.id = uls.song_id
	WHERE uls.user_id = ?
	`

	var err error
	rows, err := db.DB.Query(query, userID)

	if err != nil {
		log.Println("error al obtener los likes de la cancion", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		likedSong := &structs.UserLikesSong{Song: &structs.Song{}}
		err := rows.Scan(
			&likedSong.UserID,
			&likedSong.SongID,
			&likedSong.CreatedAt,
			&likedSong.Song.ID,
			&likedSong.Song.Title,
			&likedSong.Song.ArtistID,
			&likedSong.Song.Album,
			&likedSong.Song.Duration,
			&likedSong.Song.Genre,
			&likedSong.Song.ReleaseDate,
			&likedSong.Song.CoverImage,
			&likedSong.Song.CreatedAt,
		)
		if err != nil {
			log.Println("error al obtener los likes de la cancion", err)
			return nil, err
		}
		UsersLikesSongs = append(UsersLikesSongs, likedSong)
	}
	return UsersLikesSongs, nil

}

func (db DBServer) CreateUserLikesSong(userID, songID int) (*structs.UserLikesSong, error) {
	return createUserLikesSong(db.DB, userID, songID)
}

func (db DBServer) FindArtist(id int) (*structs.Artist, error) {
	return findArtist(db.DB, id)
}

func (db DBServer) GetArtists() ([]*structs.Artist, error) {
	return getArtists(db.DB)
}

func (db DBServer) CreateArtist(artist *structs.Artist) (*structs.Artist, error) {
	return createArtist(db.DB, artist)
}

func getSongs(db *sql.DB) ([]*structs.Song, error) {
	songs := []*structs.Song{}

	rows, err := db.Query(`SELECT id, title, artist_id, duration, album, genre, release_date, cover_image, created_at FROM songs`)
	if err != nil {
		log.Println("error al obtener listado de canciones", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		song := &structs.Song{}
		err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.ArtistID,
			&song.Duration,
			&song.Album,
			&song.Genre,
			&song.ReleaseDate,
			&song.CoverImage,
			&song.CreatedAt,
		)
		if err != nil {
			log.Println("error al obtener listado de canciones", err)
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func getSong(db *sql.DB, songID int) (*structs.Song, error) {
	song := &structs.Song{}

	rows, err := db.Query(`SELECT * FROM songs WHERE id = ?`, songID)
	if err != nil {
		log.Println("error al obtener la cancion", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.ArtistID,
			&song.Duration,
			&song.Album,
			&song.Genre,
			&song.ReleaseDate,
			&song.CoverImage,
			&song.CreatedAt,
		)
		if err != nil {
			log.Println("error al obtener la cancion por id", err)
			return nil, err
		}
	}
	return song, nil

}

func getSongByTitle(db *sql.DB, title string) (*structs.Song, error) {
	song := &structs.Song{}

	rows, err := db.Query(`SELECT * FROM songs WHERE title LIKE(?)`, title)
	if err != nil {
		log.Println("error al obtener la cancion", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.ArtistID,
			&song.Duration,
			&song.Album,
			&song.Genre,
			&song.ReleaseDate,
			&song.CoverImage,
			&song.CreatedAt,
		)
		if err != nil {
			log.Println("error al obtener la cancion por titulo", err)
			return nil, err
		}
	}
	return song, nil
}

func getSongByArtist(db *sql.DB, artistID int) ([]*structs.Song, error) {
	songs := []*structs.Song{}

	rows, err := db.Query(`SELECT * FROM songs WHERE artist_id = ?`, artistID)
	if err != nil {
		log.Println("error al obtener la cancion por artista", err)
		return nil, err
	}
	defer rows.Close()

	if rows == nil {
		return nil, nil
	}

	for rows.Next() {
		song := &structs.Song{}
		err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.ArtistID,
			&song.Duration,
			&song.Album,
			&song.Genre,
			&song.ReleaseDate,
			&song.CoverImage,
			&song.CreatedAt,
		)
		if err != nil {
			log.Println("error al obtener la cancion por titulo", err)
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil

}

func createSong(db *sql.DB, song *structs.Song) (*structs.Song, error) {
	query := `INSERT INTO songs (
		title,
		artist_id,
		duration,
		album,
		genre,
		release_date,
		cover_image,
		created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	res, err := db.Exec(query, song.Title, song.ArtistID, song.Duration, song.Album, song.Genre, song.ReleaseDate, song.CoverImage, song.CreatedAt)

	if err != nil {
		log.Println("error al crear la cancion", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("error al obtener el id de la cancion", err)
		return nil, err
	}

	song.ID = int(id)
	return song, nil

}

func updateSong(db *sql.DB, song *structs.Song) (*structs.Song, error) {
	query := `UPDATE songs
		SET title = ?,
		artist_id = ?,
		duration = ?,
		album = ?,
		genre = ?,
		release_date = ?,
		cover_image = ?,
		created_at = ?
	WHERE id = ?`

	res, err := db.Exec(query, song.Title, song.ArtistID, song.Duration, song.Album, song.Genre, song.ReleaseDate, song.CoverImage, song.CreatedAt, song.ID)

	if err != nil {
		log.Println("error al actualizar la cancion", err)
		return nil, err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Println("error al obtener el numero de filas afectadas", err)
		return nil, err
	}

	if rows == 0 {
		return nil, errors.New("no se encontró la cancion")
	}

	return song, nil
}

func deleteSong(db *sql.DB, songID int) error {
	query := `DELETE FROM songs WHERE id = ?`
	res, err := db.Exec(query, songID)
	if err != nil {
		log.Println("error al eliminar la cancion", err)
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Println("error al obtener el numero de filas afectadas", err)
		return err
	}

	if rows == 0 {
		return errors.New("no se encontró la cancion")
	}

	return nil
}

func getUserLikesSong(db *sql.DB, userID, songID int) (*structs.UserLikesSong, error) {
	query := `SELECT 
			uls.user_id,
			uls.song_id,
			uls.created_at,
			s.id,
			s.title,
			s.artist_id,
			s.album,
			s.duration,
			s.genre,
			s.release_date,
			s.cover_image,
			s.created_at
	FROM users_likes_songs uls 
	INNER JOIN songs s ON s.id = uls.song_id
	WHERE uls.user_id = ? 
	AND uls.song_id = ?
	LIMIT 1`

	row := db.QueryRow(query, userID, songID)

	Song := &structs.UserLikesSong{}
	err := row.Scan(
		&Song.UserID,
		&Song.SongID,
		&Song.CreatedAt,
		&Song.Song.ID,
		&Song.Song.Title,
		&Song.Song.ArtistID,
		&Song.Song.Album,
		&Song.Song.Duration,
		&Song.Song.Genre,
		&Song.Song.ReleaseDate,
		&Song.Song.CoverImage,
		&Song.Song.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		log.Println("error al obtener los likes de la cancion", err)
		return nil, err
	}

	return Song, nil
}

func createUserLikesSong(db *sql.DB, userID, songID int) (*structs.UserLikesSong, error) {
	query := `INSERT INTO users_likes_songs (user_id, song_id, created_at) VALUES (?, ?, ?)`
	_, err := db.Exec(query, userID, songID, getNowDateTime())
	if err != nil {
		log.Println("error al crear el like de la cancion", err)
		return nil, err
	}
	userLikesSong := &structs.UserLikesSong{
		UserID:    userID,
		SongID:    songID,
		CreatedAt: getNowDateTime(),
	}
	userLikesSong.Song, err = getSong(db, songID)
	if err != nil {
		log.Println("error al obtener la cancion", err)
		return nil, err
	}
	return userLikesSong, nil
}

func getPlaylistsSongs(db *sql.DB, playlistIDs []int) ([]*structs.Song, error) {
	query := `SELECT
		s.id,
		s.title,
		s.artist_id,
		s.duration,
		s.album,
		s.genre,
		s.release_date,
		s.cover_image,
		s.created_at
		ps.playlist_id
	FROM songs s
	INNER JOIN playlists_songs ps ON ps.song_id = s.id
	WHERE ps.playlist_id IN (?)`
	rows, err := db.Query(query, playlistIDs)
	if err != nil {
		log.Println("error al obtener listado de canciones de la playlist", err)
		return nil, err
	}

	defer rows.Close()

	songs := []*structs.Song{}
	for rows.Next() {
		song := &structs.Song{}
		err := rows.Scan(
			&song.ID,
			&song.Title,
			&song.ArtistID,
			&song.Duration,
			&song.Album,
			&song.Genre,
			&song.ReleaseDate,
			&song.CoverImage,
			&song.CreatedAt,
			&song.PlaylistID,
		)
		if err != nil {
			log.Println("error al obtener listado de canciones de la playlist", err)
			continue
		}
		songs = append(songs, song)
	}
	return songs, nil
}

func findArtist(db *sql.DB, id int) (*structs.Artist, error) {
	query := `SELECT * FROM artists WHERE id = ?`
	row := db.QueryRow(query, id)
	artist := &structs.Artist{}
	err := row.Scan(&artist.ID, &artist.Name, &artist.CreatedAt)
	return artist, err
}

func getArtists(db *sql.DB) ([]*structs.Artist, error) {
	query := `SELECT * FROM artists`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("error al obtener listado de artistas", err)
		return nil, err
	}
	defer rows.Close()

	artists := []*structs.Artist{}
	for rows.Next() {
		artist := &structs.Artist{}
		err := rows.Scan(&artist.ID, &artist.Name, &artist.CreatedAt)
		if err != nil {
			log.Println("error al obtener listado de artistas", err)
			continue
		}
		artists = append(artists, artist)
	}
	return artists, nil
}

func createArtist(db *sql.DB, artist *structs.Artist) (*structs.Artist, error) {
	query := `INSERT INTO artists (name, created_at) VALUES (?, ?)`
	res, err := db.Exec(query, artist.Name, getNowDateTime())
	if err != nil {
		log.Println("error al crear el artista", err)
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("error al obtener el id del artista", err)
		return nil, err
	}
	artist.ID = int(id)
	return artist, nil
}

func countSongs(db *sql.DB, args []interface{}, limit int) (int, error) {
	query := `SELECT COUNT(*) FROM songs WHERE `

	row := db.QueryRow(query, args...)
	var total int
	err := row.Scan(&total)

	total = getTotalPages(total, limit)
	return total, err
}
