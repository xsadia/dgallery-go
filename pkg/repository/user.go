package repository

import (
	"database/sql"
	"time"
)

type User struct {
	Id        string    `json:"id"`
	DiscordId string    `json:"discordId"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUser(m map[string]string) *User {
	return &User{
		DiscordId: m["id"],
		Email:     m["email"],
		Username:  m["username"],
	}
}

func (u *User) Create(db *sql.DB) error {
	return db.QueryRow(
		"INSERT INTO users (email, username, discord_id) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at",
		u.Email, u.Username, u.DiscordId,
	).Scan(&u.Id, &u.CreatedAt, &u.UpdatedAt)
}
