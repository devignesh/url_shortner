package model

import "time"

// URL represents a shortened URL
type URL struct {
	ID        uint      `json:"id"`
	ShortLink string    `json:"short_link"`
	LongURL   string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// URLRequest represents a request to shorten a URL
type URLRequest struct {
	LongURL string `json:"long_url"`
}

// URLResponse represents a response containing a shortened URL
type URLResponse struct {
	ShortLink string `json:"short_link"`
}
