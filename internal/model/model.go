package model

import "time"

type URL struct {
	ID        uint      `json:"id"`
	ShortLink string    `json:"short_link"`
	LongURL   string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type URLRequest struct {
	LongURL string `json:"long_url"`
}

type URLResponse struct {
	ShortLink string `json:"short_link"`
}
