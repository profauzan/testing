package handler

import (
	"fmt"
	"net/http"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func MainHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name":    "Fauzan Pratama Putra",
		"role":    "Mobile Apps Developer",
		"company": "PT IAM EDU NETWORKS",
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
	})
}

func UrlParamHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func PostBooksHandler(c *gin.Context) {
	var booksInput book.Books
	err := c.ShouldBindJSON(&booksInput)

	if err != nil {
		errorMessages := []string{}
		for _, error := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, please condition %s", error.Field(), error.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": booksInput.Title,
		"price": booksInput.Price,
	})
}
