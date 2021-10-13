package main

import (
	"github.com/gin-gonic/gin"

	"pustaka-api/handler"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.MainHandler)
	v1.GET("/query/", handler.QueryHandler)
	v1.GET("/urlParam/:id", handler.UrlParamHandler)
	v1.POST("/books", handler.PostBooksHandler)
	router.Run()
}
