package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// book represents data about a record book.
type book struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Writer string  `json:"writer"`
    Price  float64 `json:"price"`
}

// books slice to seed record book data.
var books = []book{
    {ID: "1", Title: "Blue Train", Writer: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Writer: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Writer: "Sarah Vaughan", Price: 39.99},
}

func main() {
    router := gin.Default()
    router.GET("/books", getbooks)
    router.GET("/books/:id", getbookByID)
    router.POST("/books", postbooks)

    router.Run("localhost:8080")
}

// getbooks responds with the list of all books as JSON.
func getbooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, books)
}

// postbooks adds an book from JSON received in the request body.
func postbooks(c *gin.Context) {
    var newbook book

    // Call BindJSON to bind the received JSON to
    // newbook.
    if err := c.BindJSON(&newbook); err != nil {
        return
    }

    // Add the new book to the slice.
    books = append(books, newbook)
    c.IndentedJSON(http.StatusCreated, newbook)
}

// getbookByID locates the book whose ID value matches the id
// parameter sent by the client, then returns that book as a response.
func getbookByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of books, looking for
    // an book whose ID value matches the parameter.
    for _, a := range books {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
