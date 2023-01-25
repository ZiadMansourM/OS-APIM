package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type quote struct {
	Id     int    `json:"id"`
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var quotes = []quote{
	{Id: 1, Author: "Abraham Lincoln", Quote: "Whatever you are, be a good one."},
	{Id: 2, Author: "Winston Churchill", Quote: "If you're going through hell, keep going."},
	{Id: 3, Author: "Will Rogers", Quote: "Don't let yesterday take up too much of today."},
	{Id: 4, Author: "William Wallace", Quote: "Every man dies. Not every man lives."},
	{Id: 5, Author: "Maya Angelou", Quote: "We need much less than we think we need."},
	{Id: 5, Author: "Henry Ford", Quote: "The only real mistake is the one from which we learn nothing."},
}

func listQuotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, quotes)
}

func main() {
	router := gin.Default()

	router.GET("/quotes", listQuotes)

	router.Run(":3000")
}
