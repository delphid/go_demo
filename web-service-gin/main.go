package main

import (
	"errors"
	"fmt"
	"reflect"

	//"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)


type albumer interface {
    GetTitle() string
    GetID() string
}

type specialAlbum struct {
    Name    string  `json:"name"`
}

type event struct {
    Title   string `json:"title" binding:"required"`
    Content string `json:"content"`
}

type album struct {
    ID      string  `json:"id"`
    Title   string  `json:"title" binding:"required"`
    Artist  string  `json:"artist"`
    Price   float64 `json:"price"`
    Events  []event   `json:"events"`
    Labels  struct {
        Stars   int `json:"stars" binding:"required"`
        Level   int `json:"level"`
    } `json:"labels" binding:"required"`
}

func (a album) GetTitle() string {
    return a.Title
}

func (a album) GetID() string {
    return a.ID
}

var albums = []albumer{
    album{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    album{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    album{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99, Events: []event{{"a", "b"}, {"c", "d"}}},
}

func main() {
    router := gin.Default()
    router.GET("/albums", GetAlbums)
    router.GET("/albums/:id", GetAlbumByID)
    router.POST("/albums", PostAlbums)

    router.Run("0.0.0.0:8080")
}

func GetAlbums(c *gin.Context) {
    fmt.Println(albums)
    c.IndentedJSON(http.StatusOK, albums)
}

func AlbumBind(c *gin.Context) (interface{}, error) {
	bindSuccess := false
	for _, a := range []interface{}{
		album{},
		specialAlbum{},
	} {
		b := binding.Default(c.Request.Method, c.ContentType())
		err := c.ShouldBindWith(a, b)
		if err == nil {
			bindSuccess = true
			return a, nil
		}
		// TODO: should recover iobuf after each trial bind
	}
	if !bindSuccess {
		return nil, errors.New("binding album error")
	}
}

func PostAlbums(c *gin.Context) {
    fmt.Println("inside post")
    newAlbum, err := AlbumBind(c)
    if err != nil {
        fmt.Println("can't resolve album data", err)
        return
    }
    fmt.Printf("newAlbum: %+v\n", newAlbum)
    if reflect.TypeOf(newAlbum) == reflect.TypeOf(album{}){
        albums = append(albums, newAlbum)
    }
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range albums {
        if a.GetID() == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
