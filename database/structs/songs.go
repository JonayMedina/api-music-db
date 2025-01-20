package structs

type Song struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	ArtistID    int    `json:"artist_id"`
	Duration    int    `json:"duration"`
	Album       string `json:"album"`
	Genre       string `json:"genre"`
	ReleaseDate string `json:"release_date"`
	CoverImage  string `json:"cover_image"`
	CreatedAt   string `json:"created_at"`
	PlaylistID  int    `json:"playlist_id,omitempty"`
}

type Artist struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
