package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func (s *Server) registerEventRoutes(g *gin.RouterGroup) {
	r := g.Group("events")
	r.POST("", s.EventPush)
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

type EventReq struct {
	Name string `json:"name"`
	Labels map[string]interface{} `json:"labels"`
}
