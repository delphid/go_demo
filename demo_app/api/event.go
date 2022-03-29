package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerEventRoutes(g *gin.RouterGroup) {
	r := g.Group("events")
	r.POST("", s.EventPush)
    r.POST("/tf", s.TFPush)
    r.GET("/tf", s.TFPush)
}

// EventPush can be tested with:
// curl -X POST localhost/events -d '{"name": "a", "labels": {"a": 1, "b": 2, "c": true, "d": null, "e": "eee"}}'
func (s *Server) EventPush(c *gin.Context) {
	var req EventReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err)
		c.JSON(
			400, map[string]error{"error": err},
		)
		return
	}
	fmt.Printf("req: %+v\n", req)
	c.JSON(
		200, map[string]interface{}{
			"msg": "ok",
			"data": req,
		},
	)
	return
}

// TFPush can be tested with:
// curl -X POST localhost/events/tf -d '{"a": true, "b": true}'
func (s *Server) TFPush(c *gin.Context) {
	var req TFReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println(err)
		c.JSON(
			400, map[string]error{"error": err},
		)
		return
	}
	fmt.Printf("req: %+v\n", req)
    fmt.Println("req.A: ", req.A)
    fmt.Println("req.B: ", *req.B)
    fmt.Println("req.C: ", req.C)
	c.JSON(
		200, map[string]interface{}{
			"msg": "ok",
			"data": req,
		},
	)
	return
}

type EventReq struct {
	Name string `json:"name"`
	Labels map[string]interface{} `json:"labels"`
}

type TFReq struct {
    A bool `json:"a" binding:"required"`
    B *bool `json:"b" binding:"required"`
    C *bool `json:"c" binding:"omitempty"`

}

type TFFormReq struct {
    A bool `form:"a" binding:"required"`
    B *bool `form:"b" binding:"required"`
    C bool `form:"c"`
    D *bool `form:"d"`
}
