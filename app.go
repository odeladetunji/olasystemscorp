package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){

	router := gin.Default();
	router.Static("/assets", "./assets");
	// router.LoadHTMLFiles("./templates/app.html");
	router.LoadHTMLFiles("./templates/app.html");

	router.GET("/", func(c *gin.Context){
		c.HTML(http.StatusOK, "app.html", gin.H{
			"message": "Hypersonic Inc",
		})
	})

	if err := router.Run(":8086"); err != nil {
		fmt.Println("Go server is running on port 8086")
	}

}