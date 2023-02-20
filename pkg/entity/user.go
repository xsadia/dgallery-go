package entity

import "time"

type User struct {
	Id          int       `json:"id"`
	DiscordId   string    `json:"discordId"`
	Email       string    `json:"email"`
	AccessToken string    `json:"accessToken"`
	Username    string    `json:"username"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
