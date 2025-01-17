package structs

type User struct {
	ID        int             `json:"id"`
	Username  string          `json:"username"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	CreatedAt string          `json:"created_at"`
	Roles     []Role          `json:"roles"`
	Playlists []Playlist      `json:"playlists"`
	Likes     []UserLikesSong `json:"likes"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Playlist struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	Songs     []Song `json:"songs"`
}

type UserLikesSong struct {
	UserID    int    `json:"user_id"`
	SongID    int    `json:"song_id"`
	CreatedAt string `json:"created_at"`
	Song      Song   `json:"song"`
}
