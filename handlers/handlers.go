package handlers

import (
	"github.com/gin-gonic/gin"
	"microservices/database"
	"microservices/shortener"
	"net/http"
)

const (
	Host = "http://localhost"
	Port = ":5000"
)

// request struct, just contians URL and ID
type ShortenRequest struct {
	URL   string `json:"original_url" binding:"required"`
	ReqID string `json:"request_id" binding:"required"`
}

func DefaultHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "default response for URL shortening API"})
}

func CreateShortenedURL(c *gin.Context) {
	// checking if request contains everything needed
	var req ShortenRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	shortened := shortener.ShortenURL(req.URL, req.ReqID)
	database.StoreURL(shortened, req.URL)

	c.JSON(http.StatusOK, gin.H{
		"message":   "successfully shortened URL",
		"short_url": Host + Port + "/" + shortened,
	})
}

func HandleShortenedRedirect(c *gin.Context) {
	shortened := c.Param("shortenedURL")
	original := database.RetrieveURL(shortened)
	c.Redirect(http.StatusFound, original)
}
