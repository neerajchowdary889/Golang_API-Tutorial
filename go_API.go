package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 3},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 7},
}

func getbooks(c *gin.Context) {
	// .IndentedJSON return nicely Printed JSON, like it return JSON with nicely Indented.
	// StatusOK return http : 200
	// we need to enter "curl localhost:8080/books" to get the struct books
	c.IndentedJSON(http.StatusOK, books)
}

func updatebooks(c *gin.Context) {
	newbook := book{ID: "4", Title: "Atomic Habits", Author: "Dont know", Quantity: 3}
	books = append(books, newbook)
	//If the newbook was created then we will get http : 201 that means Req was successfull and resource has been created.
	//we need to enter " curl localhost:8080/books --request "POST" " command to get newbook details and add to the books struct.
	c.IndentedJSON(http.StatusCreated, newbook)

}
func main() {
	router := gin.Default()
	//.GET is the get request we creating to get the data.
	//.RUN is the request to get server started.
	router.GET("/books", getbooks)
	router.POST("/books", updatebooks)
	router.Run("localhost:8080")
}
