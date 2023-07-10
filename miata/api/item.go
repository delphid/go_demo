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
	r.GET("/:item_id", s.ItemGetByID)
	r.GET("", s.ItemSearch)
	r.PATCH("/:item_id", s.ItemPatch)
}

// ItemCreate can be tested with:
// curl -X POST localhost/items -d '{"name": "item1", "labels": {"a": "aaa"}, "annotations": {"b": "bbb"}}'
func (s *Server) ItemCreate(c *gin.Context) {
	var req model.ItemReq
	var item *model.Item
	var err error
	err = c.ShouldBindJSON(&req)
	if err != nil {
		s.log.Error("bind error", zap.Error(err))
		c.JSON(
			400, map[string]error{"error": err},
		)
		return
	}
	s.log.Info("received item", zap.Any("req", req))
	item = req.ToModel()
	err = s.store.CreateItem(item)
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

// ItemGetByID can be tested with:
// curl localhost/items/1
func (s *Server) ItemGetByID(c *gin.Context) {
	var req model.ItemIDReq
	var resp model.ItemResp
	var err error
	var item *model.Item
	if err = c.ShouldBindUri(&req); err != nil {
		s.log.Error("ShouldBindUri error", zap.Error(err))
		c.JSON(
			400, map[string]error{"error": err},
		)
		return
	}
	if item, err = s.store.GetItemByID(&req.ItemID); err != nil {
		s.log.Error("ItemGetByID error", zap.Error(err))
		c.JSON(
			400, map[string]error{"error": err},
		)
		return
	}
	if resp, err = item.ToResp(); err != nil {
		s.log.Error("ToResp error", zap.Error(err))
		c.JSON(
			400, map[string]error{"error": err},
		)
		return
	}
	c.JSON(
		200, map[string]interface{}{
			"msg": "ok",
			"data": resp,
		},
	)
	return
}

// ItemSearch can be tested with:
// curl localhost/items
func (s *Server) ItemSearch(c *gin.Context) {
	var items model.Items
	var resps []model.ItemResp
	var err error
	if items, err = s.store.SearchItem(); err != nil {
		s.log.Error("SearchItem error", zap.Error(err))
		c.JSON(
			400, map[string]error{"error": err},
		)
		return
	}
	if resps, err = items.ToResps(); err != nil {
		s.log.Error("ToResps error", zap.Error(err))
		c.JSON(
			400, map[string]error{"error": err},
		)
		return
	}
	c.JSON(
		200, map[string]interface{}{
			"msg": "ok",
			"data": resps,
		},
	)
	return
}

// ItemPatch can be tested with:
// curl -X PATCH localhost/items/2 -d '{"name": "item2"}'
func (s *Server) ItemPatch(c *gin.Context) {
	var req model.ItemReq
	var idReq model.ItemIDReq
	var err error
	if err = c.ShouldBindJSON(&req); err != nil {
		s.log.Error("ShouldBindJSON error", zap.Error(err))
		c.JSON(
			400, map[string]error{"error": err},
		)
		return
	}
	if err = c.ShouldBindUri(&idReq); err != nil {
		s.log.Error("ShouldBindUri error", zap.Error(err))
		c.JSON(
			400, map[string]error{"error": err},
		)
		return
	}
	if err = s.store.PatchItem(idReq.ItemID, req); err != nil {
		s.log.Error("PatchItem error", zap.Error(err))
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
