package structs

type User struct {
	ID        int              `json:"id"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	CreatedAt string           `json:"created_at"`
	Roles     []*Role          `json:"roles,omitempty"`
	Playlists []*UserPlaylist  `json:"playlists,omitempty"`
	Likes     []*UserLikesSong `json:"likes,omitempty"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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
