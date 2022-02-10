package main

import (
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/books/:id/:title", handler.BooksHandler)
	v1.POST("/books", handler.PostBooksHandler)
	v1.GET("/query", handler.QueryHandler)

	router.Run(":8888")

}
