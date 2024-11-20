package domain

type User struct {
	id            int    `json:"id",db:"id"`
	name          string `json:"name",db:"name"`
	email         string `json:"email",db:"email"`
	password_hash string `json:"password_hash",db:"password_hash"`
}
