package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/url-shortener/backend/validator"
)

type urlRequest struct {
	URL string
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
}

func Redirect(c *gin.Context) {

}

func Info(c *gin.Context) {

}
