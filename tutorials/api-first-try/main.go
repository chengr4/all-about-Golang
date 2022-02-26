package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	// uppercase means public field
	Id       string `json: "id"`
	Title    string `json: "title"`
	Author   string `json: "author"`
	Quantity int    `json: "quantity"`
}

var books = []book{
	{Id: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{Id: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{Id: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 6},
}

func createBook(context *gin.Context) {
	var newBook book

	// if err:= c.BindJSON
}

func getBook(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, books)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBook)
	router.Run("localhost:8080")
}
