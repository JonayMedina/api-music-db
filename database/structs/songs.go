package structs

import "time"

type Song struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	ArtistID    int       `json:"artist_id,omitempty"`
	Duration    int       `json:"duration,omitempty"`
	Album       string    `json:"album,omitempty"`
	Genre       string    `json:"genre,omitempty"`
	ReleaseDate time.Time `json:"release_date,omitempty"`
	CoverImage  string    `json:"cover_image,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	PlaylistID  int       `json:"playlist_id,omitempty"`
	Price       string    `json:"price,omitempty"`
	Origin      string    `json:"origin,omitempty"`
	Artist      *Artist   `json:"artist,omitempty"`
}

type Artist struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type Playlist struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	CreatedAt string  `json:"created_at"`
	Songs     []*Song `json:"songs,omitempty"`
}

type UserLikesSong struct {
	UserID    int    `json:"user_id"`
	SongID    int    `json:"song_id"`
	CreatedAt string `json:"created_at"`
	Song      *Song  `json:"song,omitempty"`
}

type UserPlaylist struct {
	UserID     int       `json:"user_id"`
	PlaylistID int       `json:"playlist_id"`
	CreatedAt  string    `json:"created_at"`
	Playlist   *Playlist `json:"playlist,omitempty"`
}
