package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/url-shortener/backend/database"
	"github.com/hyperxpizza/url-shortener/backend/validator"
)

type urlRequest struct {
	URL        string
	Expiration string
}

var db database.Database

func init() {
	db = database.NewDatabase()
}

func Encode(c *gin.Context) {
	var request urlRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	//validate url
	if !validator.ValidateUrl(request.URL) {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Url is not valid",
		})
		return
	}

	var expiresAt time.Duration

	switch request.Expiration {
	case "never":
		expiresAt = 0
	case "5":
		expiresAt = time.Duration(5 * time.Minute)
	case "10":
		expiresAt = time.Duration(10 * time.Minute)
	case "15":
		expiresAt = time.Duration(15 * time.Minute)
	case "30":
		expiresAt = time.Duration(30 * time.Minute)
	case "hour":
		expiresAt = time.Duration(time.Hour)
	default:
		expiresAt = time.Duration(5 * time.Minute)
	}

	code, err := db.Insert(request.URL, expiresAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"code":    code,
	})
}

func Redirect(c *gin.Context) {
	id := c.Param("id")

}

func Info(c *gin.Context) {

}
