package storage

type Storage interface {
	Save(p *AddShopItem) error
	ShopItems(...int64) ([]ShopItem, error)
	ChangeShopItemCount(string, int) error
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
