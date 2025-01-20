package structs

type Song struct {
	ID          int     `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	ArtistID    int     `json:"artist_id,omitempty"`
	Duration    int     `json:"duration,omitempty"`
	Album       string  `json:"album,omitempty"`
	Genre       string  `json:"genre,omitempty"`
	ReleaseDate string  `json:"release_date,omitempty"`
	CoverImage  string  `json:"cover_image,omitempty"`
	CreatedAt   string  `json:"created_at,omitempty"`
	PlaylistID  int     `json:"playlist_id,omitempty"`
	Artist      *Artist `json:"artist,omitempty"`
}

type Artist struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}
