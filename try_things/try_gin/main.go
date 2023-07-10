package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println("method: ", method)
		fmt.Println("url: ", c.Request.URL)
		origin := c.Request.Header.Get("Origin")
		fmt.Println("origin: ", origin)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	})

	r.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		log.Println("long sync")
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(time.Second * 60 * 4)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
		c.JSON(http.StatusOK, nil)
	})

	r.GET("/view", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
		<!DOCTYPE html>
		<html>
			<head>
				<meta http-equiv="refresh" content="7; url='https://www.w3docs.com'" />
			</head>
			<body>
				<p>Please follow <a href="https://www.w3docs.com">this link</a>.</p>
				<script>
					console.log("fetch in script");
					fetch("https://www.w3docs.com");
				</script>
			</body>
		</html>
	`))
		return
	})
	r.GET("/redirect", func(c *gin.Context) {
		//c.Redirect(http.StatusFound, "http://127.0.0.1:8083/view")
		c.Redirect(http.StatusFound, "http://127.0.0.1:8080/resources")
		return
	})

	r.Run(":8083")
}
