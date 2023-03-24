package handler

import (
	"log"
	"net/http"

	"url_shortener/internal/model"
	"url_shortener/internal/service"

	"github.com/gin-gonic/gin"
)

// URLShortenerHandler defines methods to handle URL shortener requests
type URLShortenerHandler interface {
	ShortenURL(c *gin.Context)
	Redirect(c *gin.Context)
}

type urlShortenerHandler struct {
	service service.URLShortenerService
}

// NewURLShortenerHandler creates a new instance of URLShortenerHandler
func NewURLShortenerHandler(service service.URLShortenerService) URLShortenerHandler {
	return &urlShortenerHandler{
		service: service,
	}
}

func (h *urlShortenerHandler) ShortenURL(c *gin.Context) {

	var req model.URLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if req.LongURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "long_url cannot be empty"})
		return
	}

	url, err := h.service.ShortenURL(req.LongURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to shorten URL"})
		return
	}

	// url.ShortLink = c.Request.Host + "/" + url.ShortLink

	c.JSON(http.StatusOK, url)
}

func (h *urlShortenerHandler) Redirect(c *gin.Context) {

	shortLink := c.Param("shortlink")
	log.Println("\n redurected", shortLink)
	if shortLink == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "shortLink cannot be empty"})
		return
	}

	longURL, err := h.service.Redirect(shortLink)
	log.Println("\n longurl from redirect", longURL)
	if err != nil {
		if err.Error() == "expired link" {
			c.JSON(http.StatusGone, gin.H{"error": "expired link"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "shortLink not found"})
		}
		return
	}

	c.Redirect(http.StatusMovedPermanently, longURL)
}

// SetupRouter setups the gin router for the URL shortener service
func SetupRouter(service service.URLShortenerService) *gin.Engine {

	log.Println("\n came to setup routes")
	handler := NewURLShortenerHandler(service)
	router := gin.Default()

	router.POST("/api/shorten", handler.ShortenURL)
	router.GET("/api/shorten/:shortlink", handler.Redirect)

	return router
}
