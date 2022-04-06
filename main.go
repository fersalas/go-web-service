package main

import (
	"exercises/go-web-service/internal/usecases/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	var err error
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)

	err = router.Run("localhost:8083")
	if err != nil {
		fmt.Println("Something went wrong with the router")
	}

	fmt.Println("App up and running.")

}
