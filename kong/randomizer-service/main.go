package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getRandom(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, rand.Intn(100))
}

func main() {
	router := gin.Default()

	router.GET("/", getRandom)

	router.Run(":3000")
}
