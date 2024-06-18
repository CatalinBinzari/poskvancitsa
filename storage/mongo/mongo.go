package mongo

import (
	"context"
	"fmt"
	"log"
	"poskvancitsa/storage"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	shopList ShopList
}

type ShopList struct {
	*mongo.Collection
}

func New(connectString string, connectTimeout time.Duration) Storage {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectString))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	shopList := ShopList{
		Collection: client.Database("poskvancitsa").Collection("shoplist"),
	}

	return Storage{
		shopList: shopList,
	}
}

func (s Storage) Save(i *storage.AddShopItem) error {

	_, err := s.shopList.InsertOne(context.Background(), storage.AddShopItem{
		AddedBy:  i.AddedBy,
		Count:    i.Count,
		ItemName: i.ItemName,
	})
	if err != nil {
		return fmt.Errorf("can't save shop item %s", err)
	}

	return nil
}

func (s Storage) ShopItems() ([]storage.ShopItem, error) {

	filter := bson.M{}

	cur, err := s.shopList.Find(context.Background(), filter)
	if err != nil {
		return []storage.ShopItem{}, err
	}
	defer cur.Close(context.Background())

	var items []storage.ShopItem
	for cur.Next(context.Background()) {
		var item storage.ShopItem
		if err := cur.Decode(&item); err != nil {
			return nil, err
		}

		items = append(items, item)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (s Storage) PlusOneShopItem(_id string) error {
	mongoID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": mongoID}
	update := bson.D{{"$inc", bson.D{{"count", 1}}}}
	_, err = s.shopList.UpdateOne(context.Background(), filter, update,
		options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("can't update shop item %s", err)
	}

	return nil
}

func (s Storage) MinusOneShopItem(_id string) error {
	mongoID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": mongoID}
	update := bson.D{{"$inc", bson.D{{"count", -1}}}}
	_, err = s.shopList.UpdateOne(context.Background(), filter, update,
		options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("can't update shop item %s", err)
	}

	return nil
}

func (s Storage) RemoveShopItem(_id string) error {
	mongoID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": mongoID}
	_, err = s.shopList.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("can't delete shop item %s", err)
	}

	return nil
}

func (s Storage) ModifyNameShopItem(_id string, itemName string) error {
	mongoID, err := primitive.ObjectIDFromHex(_id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": mongoID}
	update := bson.D{{"$set", bson.D{{"itemName", itemName}}}}
	_, err = s.shopList.UpdateOne(context.Background(), filter, update,
		options.Update().SetUpsert(true))
	if err != nil {
		return fmt.Errorf("can't update shop item %s", err)
	}

	return nil
}

// func (s Storage) IsExists(ctx context.Context, storagePage *storage.Page) (bool, error) {
// 	count, err := s.pages.CountDocuments(ctx, toPage(storagePage).Filter())
// 	if err != nil {
// 		return false, e.Wrap("can't check if page exists", err)
// 	}

// 	return count > 0, nil
// }

// func toPage(p *storage.Page) Page {
// 	return Page{
// 		URL:      p.URL,
// 		UserName: p.UserName,
// 	}
// }

// func (p Page) Filter() bson.M {
// 	return bson.M{
// 		"url":      p.URL,
// 		"username": p.UserName,
// 	}
// }
