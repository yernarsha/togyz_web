package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"togyz_web/models"
)

func main() {
	err := models.ConnectDatabase()
	checkErr(err)

	router := gin.Default()
	//ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Modern TogyzKumalak",
			"games": models.GetGames(1000),
		})
	})

	router.GET("/game/:id", func(c *gin.Context) {
		c.HTML(http.StatusOK, "game.tmpl", gin.H{
			"game": models.GetGameById(c.Param("id")),
		})
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	//    router.Run(":8080")

}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
