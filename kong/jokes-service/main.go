package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type joke struct {
	Id   int    `json:"id"`
	Joke string `json:"joke"`
}

var jokes = []joke{
	{Id: 1, Joke: "What's the best thing about Switzerland? .. I don't know, but the flag is a big plus."},
	{Id: 2, Joke: "I invented a new word! .. Plagiarism!"},
	{Id: 3, Joke: "Hear about the new restaurant called Karma? .. There's no menu: You get what you deserve."},
	{Id: 4, Joke: "Why don't scientists trust atoms? .. Because they make up everything."},
	{Id: 5, Joke: "Where are average things manufactured? .. The satisfactory."},
	{Id: 5, Joke: "Why don't Calculus majors throw house parties? .. Because you should never drink and derive."},
}

func listJokes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jokes)
}

func main() {
	router := gin.Default()

	router.GET("/jokes", listJokes)

	router.Run(":3000")
}
