package repository

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type URLShortenerRepository interface {
	FindByShortLink(shortLink string) (*URL, error)
	FindByLongURL(longURL string) (*URL, int, error)
	CreateShortURL(url *URL) error
}

// URL is a data structure for storing URLs in the database
type URL struct {
	ID        uint      `gorm:"primary_key"`
	ShortLink string    `gorm:"unique_index"`
	LongURL   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	ExpiresAt time.Time `gorm:"not null"`
}

type urlShortenerRepository struct {
	db *gorm.DB
}

// NewURLShortenerRepository creates a new instance of URLShortenerRepository
func NewURLShortenerRepository(db *gorm.DB) URLShortenerRepository {
	return &urlShortenerRepository{db: db}
}

func (r *urlShortenerRepository) FindByShortLink(shortLink string) (*URL, error) {
	var url URL
	if err := r.db.Where("short_link = ?", shortLink).First(&url).Error; err != nil {
		return nil, err
	}
	return &url, nil
}

func (r *urlShortenerRepository) FindByLongURL(longURL string) (*URL, int, error) {
	var url URL
	var count int

	r.db.Count(&count)

	if err := r.db.Where("long_url = ?", longURL).First(&url).Error; err != nil {
		return nil, 0, err
	}
	return &url, count, nil
}

func (r *urlShortenerRepository) CreateShortURL(url *URL) error {
	log.Println("\n url in repo", url)
	log.Println("\n db in repo", r.db.Create(url).Error)
	return r.db.Create(url).Error
}
