package main

import (
    "fmt"
    //"io/ioutil"
    "net/http"

    "github.com/gin-gonic/gin"
)


type event struct {
    Title   string `json:"title"`
    Content string `json:"content"`
}

type album struct {
    ID      string  `json:"id"`
    Title   string  `json:"title"`
    Artist  string  `json:"artist"`
    Price   float64 `json:"price"`
    Events  []event   `json:"events"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99, Events: []event{{"a", "b"}, {"c", "d"}}},
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)

    router.Run("0.0.0.0:8080")
}

func getAlbums(c *gin.Context) {
    fmt.Println(albums)
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
    var newAlbum album
    fmt.Println("inside post")
    /*
    buf, _ := ioutil.ReadAll(c.Request.Body)
    fmt.Println(buf)
    */
    if err := c.BindJSON(&newAlbum); err != nil {
        fmt.Println("can't resolve album data")
        return
    }
    fmt.Println(newAlbum)
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
