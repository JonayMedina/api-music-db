package structs

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int                `json:"id,omitempty"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	MongoID   primitive.ObjectID `json:"mongoid,omitempty" bson:"_id,omitempty"`
	CreatedAt string             `json:"created_at"`
	Roles     []*Role            `json:"roles,omitempty"`
	Playlists []*UserPlaylist    `json:"playlists,omitempty"`
	Likes     []*UserLikesSong   `json:"likes,omitempty"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}
