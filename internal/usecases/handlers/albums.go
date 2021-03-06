package handlers

import (
	"exercises/go-web-service/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

var albums = []domain.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}
