package storage

import (
	"errors"
)

type Storage interface {
	Save(p *AddShopItem) error
	ShopItems() ([]ShopItem, error)
	PlusOneShopItem(string) error
	MinusOneShopItem(string) error
	RemoveShopItem(string) error
	ModifyNameShopItem(string, string) error
}

type ShopItem struct {
	ID       string `bson:"_id"`
	AddedBy  string `bson:"addedBy"`
	Count    int    `bson:"count"`
	ItemName string `bson:"itemName"`
}

type AddShopItem struct {
	AddedBy  string `bson:"addedBy"`
	Count    int    `bson:"count"`
	ItemName string `bson:"itemName"`
}

var ErrNoSavedPages = errors.New("no saved pages")
