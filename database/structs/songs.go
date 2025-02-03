package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type Song struct {
	ID          int                `json:"id,omitempty"`
	MongoID     primitive.ObjectID `json:"mongoid,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty"`
	ArtistID    int                `json:"artist_id,omitempty"`
	Duration    int                `json:"duration,omitempty"`
	Album       string             `json:"album,omitempty"`
	Genre       string             `json:"genre,omitempty"`
	ReleaseDate string             `json:"release_date,omitempty"`
	CoverImage  string             `json:"cover_image,omitempty"`
	CreatedAt   string             `json:"created_at,omitempty"`
	PlaylistID  int                `json:"playlist_id,omitempty"`
	Price       string             `json:"price,omitempty"`
	Origin      string             `json:"origin,omitempty"`
	Artist      *Artist            `json:"artist,omitempty"`
}

type Artist struct {
	ID        int                `json:"id,omitempty"`
	MongoID   primitive.ObjectID `json:"mongoid,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty"`
	CreatedAt string             `json:"created_at,omitempty"`
}

type Playlist struct {
	ID        int                `json:"id"`
	MongoID   primitive.ObjectID `json:"mongoid,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name"`
	CreatedAt string             `json:"created_at"`
	Songs     []*Song            `json:"songs,omitempty"`
}

type UserLikesSong struct {
	UserID    int                `json:"user_id"`
	MongoID   primitive.ObjectID `json:"mongoid,omitempty" bson:"_id,omitempty"`
	SongID    int                `json:"song_id"`
	CreatedAt string             `json:"created_at"`
	Song      *Song              `json:"song,omitempty"`
}

type UserPlaylist struct {
	UserID     int                `json:"user_id"`
	MongoID    primitive.ObjectID `json:"mongoid,omitempty" bson:"_id,omitempty"`
	PlaylistID int                `json:"playlist_id"`
	CreatedAt  string             `json:"created_at"`
	Playlist   *Playlist          `json:"playlist,omitempty"`
}
