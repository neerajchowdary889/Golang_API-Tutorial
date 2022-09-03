package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type book struct {
	//These are the fields that we created and field names in json file will in all small.
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
	//We can enter book details in newbook and we make it a POST request to add to books.
	newbook := book{ID: "4", Title: "Atomic Habits", Author: "Dont know", Quantity: 3}
	books = append(books, newbook)
	//If the newbook was created then we will get http : 201 that means Req was successfull and resource has been created.
	//we need to enter " curl localhost:8080/books --request "POST" " command to get newbook details and add to the books struct.
	c.IndentedJSON(http.StatusCreated, newbook)
}

func getBookbyID(id string) (*book, error) {
	//Here getBookbyID take input int and search that id exists or not in the books array.
	for i, bookk := range books {
		if bookk.ID == id {
			return &books[i], nil
			//here if id exists then it returns the position of book. we used pointers, one of the best feature of GO.
		}
	}
	return nil, errors.New("No book bud...")
}

func bookbyID(c *gin.Context) {
	// command is "curl localhost:8080/books/2" or 1 anything.
	// If we enter 2 we get output of book of ID 2.
	id := c.Param("id")
	book, err := getBookbyID(id)

	if err != nil {
		log.Panic("panicking...")
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	//.GET is the get request we creating to get the data.
	//.RUN is the request to get server started.
	router.GET("/books", getbooks) // To get the all books.

	router.GET("/books/:id", bookbyID) // To get the book with any ID.

	router.POST("/books", updatebooks) // This is post req. UpdateBook

	router.Run("localhost:8080") // for initilizing the server.
}
