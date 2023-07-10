package model

import "miata/model/common"

type ItemReq struct {
	ID uint `uri:"id"`
	Name string `json:"name"`
	Labels map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	Price common.Price `json:"price"`
}

func (i ItemReq) ToModel() *Item {
	item := Item{
		Name: i.Name,
		Labels: i.Labels,
		Annotations: i.Annotations,
		Price: i.Price,
	}
	return &item
}

type ItemResp struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Labels map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	Price common.Price `json:"price"`
}

type ItemIDReq struct {
	ItemID uint `uri:"item_id" binding:"required"`
}
