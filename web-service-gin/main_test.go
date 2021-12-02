package main

import (
    "bytes"
    "fmt"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
)

func TestGetAlbums(t *testing.T) {
    r := gin.New()
    r.GET("/albums", GetAlbums)
    req, _ := http.NewRequest(http.MethodGet, "/albums", nil)
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    fmt.Println("w: \n", fmt.Sprintf("%+v", w))
}

func TestPostAlbums(t *testing.T) {
    r := gin.New()
    r.POST("/albums", PostAlbums)
    albums := album{
        ID: "4",
        Title: "ttt",
        Labels: struct {
            Stars int `json:"stars" binding:"required"`
            Level int `json:"level"`
            }{
                Stars: 2,
            },
    }
    reqBody, _ := json.Marshal(albums)
    req, _ := http.NewRequest(http.MethodPost, "/albums", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    fmt.Println("w: \n", fmt.Sprintf("%+v", w))
}
