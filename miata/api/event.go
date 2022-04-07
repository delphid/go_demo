package api

import (
	"fmt"
	//"errors"

	"github.com/gin-gonic/gin"

	"miata/model"
)

func (s *Server) registerEventRoutes(g *gin.RouterGroup) {
	r := g.Group("events")
	r.POST("", s.EventPush)
    r.POST("/tf", s.TFPush)
    r.GET("/tf", s.TFPush)
    r.POST("nest", s.NestEventPush)
}

// EventPush can be tested with:
// curl -X POST localhost/events -d '{"name": "a", "labels": {"a": 1, "b": 2, "c": true, "d": null, "e": "eee"}}'
func (s *Server) EventPush(c *gin.Context) {
	var req model.EventReq
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
	var req model.TFReq
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
	fmt.Println("req.C==nil: ", req.C==nil)
	fmt.Println("*req.C==false: ", *req.C==false)
	c.JSON(
		200, map[string]interface{}{
			"msg": "ok",
			"data": req,
		},
	)
	return
}

// NestEventPush can be tested with:
// curl -X POST localhost/events/nest -d '{"name": "aaa", "annotations": {"summary": "sss"}}'
func (s *Server) NestEventPush(c *gin.Context) {
	var req model.NestEventReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println("this is err: ", err)
		fmt.Println("this is err.Error(): ", err.Error())
		fmt.Println(`this is map[string]error{"error": err}: `, map[string]error{"error": err})
		// c.JSON(
		// 	400, map[string]interface{}{"error": err, "msg": err.Error()},
		// )
        c.JSON(
        	400, struct {
            	Msg  interface{} `json:"msg"`
            	Data interface{} `json:"data"`
            }{err, err.Error()},
        )
		//c.JSON(
		//	400, map[string]interface{}{"error": errors.New("aaa"), "msg": errors.New("aaa").Error()},
		//)
        fmt.Println("err resp is sent")
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
