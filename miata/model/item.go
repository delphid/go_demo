package model

import "miata/model/common"

type Item struct {
	Name string `json:"name"`
	Labels common.JSONMap `json:"labels"`
	Annotations common.JSONMap `json:"annotations"`
}

func (s *Store) CreateItem(item *Item) error {
	return s.db.Create(item).Error
}