package model

import (
	"errors"
	"fmt"
	"miata/model/common"
)

type Item struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Labels common.JSONMap `json:"labels"`
	Annotations common.JSONMap `json:"annotations"`
	Price common.Price `json:"price"`
}

func (i Item) ToResp() (ItemResp, error) {
	resp := ItemResp{
		ID: i.ID,
		Name: i.Name,
		Labels: i.Labels,
		Annotations: i.Annotations,
		Price: i.Price,
	}
	return resp, nil
}

func (s *Store) CreateItem(item *Item) error {
	return s.db.Create(item).Error
}

type Items []Item

func (items Items) ToResps() ([]ItemResp, error) {
	var resp ItemResp
	var err error
	resps := make([]ItemResp, 0, len(items))
	for _, item := range items {
		if resp, err = item.ToResp(); err != nil {
			return resps, err
		}
		resps = append(resps, resp)
	}
	return resps, nil
}

func (s *Store) GetItemByID(i *uint) (*Item, error) {
	if i == nil {
		return nil, fmt.Errorf("GetItemByID error caused by: %w", errors.New("nil item id"))
	}
	var item Item
	var err error
	if err = s.db.First(&item, *i).Error; err != nil {
		return nil, fmt.Errorf("GetItemByID error caused by: %w", err)
	}
	return &item, nil
}

func (s *Store) SearchItem() ([]Item, error) {
	var items []Item
	var err error
	//if err = s.db.Model(&Item{}).Where("id IN ?", []int{1, 2}).Find(&items).Error; err != nil {
	//	return items, fmt.Errorf("SearchItem caused by: %w", err)
	//}
	if err = s.db.Model(&Item{}).Where(`labels->>"$.a" = ? AND annotations->>"$.b" = ?`, "a2", "b3").Find(&items).Error; err != nil {
		return items, fmt.Errorf("SearchItem caused by: %w", err)
	}
	return items, nil
}

func (s *Store) PatchItem(id uint, req ItemReq) error {
	var item *Item
	item = req.ToModel()
	if err := s.db.Model(&item).Where("id = ?", id).Updates(item).Error; err != nil {
		return fmt.Errorf("PatchItem error caused by: %w", err)
	}
	return nil
}
