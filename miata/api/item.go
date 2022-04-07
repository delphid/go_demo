package api

import (
	//"errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"miata/model"
)

func (s *Server) registerItemRoutes(g *gin.RouterGroup) {
	r := g.Group("items")
	r.POST("", s.ItemCreate)
	//r.GET("/:item_id", s.ItemGetByID)
}

// ItemCreate can be tested with:
// curl -X POST localhost/items -d '{"name": "item1", "labels": {"a": 1}, "annotations": {"b": "bbb"}}'
func (s *Server) ItemCreate(c *gin.Context) {
	var req model.ItemReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		s.log.Error("bind error", zap.Error(err))
		c.JSON(
			400, map[string]error{"error": err},
		)
		return
	}
	s.log.Info("received item", zap.Any("req", req))
	item := model.Item{
		Name: req.Name,
		Labels: req.Labels,
		Annotations: req.Annotations,
	}
	err = s.store.CreateItem(&item)
	if err != nil {
		s.log.Error("create item error", zap.Error(err))
		c.JSON(
			400, map[string]error{"error": err},
		)
		return
	}
	c.JSON(
		200, map[string]interface{}{
			"msg": "ok",
			"data": req,
		},
	)
	return
}