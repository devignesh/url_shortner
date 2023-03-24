package service

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"url_shortener/internal/model"
	"url_shortener/internal/repository"

	"github.com/google/uuid"
)

// URLShortenerService defines methods to interact with the URL shortener service
type URLShortenerService interface {
	ShortenURL(longURL string) (*model.URLResponse, error)
	Redirect(shortLink string) (string, error)
}

type urlShortenerService struct {
	repo repository.URLShortenerRepository
}

const (
	shortLinkPrefix = "my-short-link/"
	shortLinkLength = 8
	expiryTime      = time.Hour * 24
)

// NewURLShortenerService creates a new instance of URLShortenerService
func NewURLShortenerService(repo repository.URLShortenerRepository) URLShortenerService {

	return &urlShortenerService{
		repo: repo,
	}
}

func (s *urlShortenerService) ShortenURL(longURL string) (*model.URLResponse, error) {

	url, count, err := s.repo.FindByLongURL(longURL)
	if err == nil {
		return &model.URLResponse{ShortLink: url.ShortLink}, nil
	}

	if count >= 20000 {
		return &model.URLResponse{}, err
	}

	log.Println("\n longurl service", longURL)

	shortLink, err := s.generateShortLink()
	if err != nil {
		return &model.URLResponse{}, err
	}

	log.Println("\n shotlink in service", shortLink)

	url = &repository.URL{
		ShortLink: shortLink,
		LongURL:   longURL,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(expiryTime),
	}

	err = s.repo.CreateShortURL(url)
	log.Println("\n err in create", err)
	if err == nil {
		return &model.URLResponse{ShortLink: url.ShortLink}, nil
	}

	return &model.URLResponse{ShortLink: url.ShortLink}, nil
}

func (s *urlShortenerService) Redirect(shortLink string) (string, error) {

	log.Println("\n redirect")

	url, err := s.repo.FindByShortLink(shortLink)
	if err != nil {
		return "", err
	}

	if url.ExpiresAt.Before(time.Now()) {
		return "", errors.New("expired link")
	}

	return url.LongURL, nil
}

func (s *urlShortenerService) generateShortLink() (string, error) {

	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	hash := md5.Sum([]byte(uuid.String()))
	hexHash := hex.EncodeToString(hash[:])

	// return shortLinkPrefix + hexHash[:shortLinkLength], nil
	return hexHash[:shortLinkLength], nil
}
