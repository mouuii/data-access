package data

import (
	"context"
	"github/mouuii/mongo/biz"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ biz.CartRepo = (*cartRepo)(nil)

type cartRepo struct {
	data     *Data
	cartColl *mongo.Collection
}

func NewCartRepo(data *Data) biz.CartRepo {
	return &cartRepo{
		data:     data,
		cartColl: data.db.Collection("cart"),
	}
}

type Cart struct {
	UserId int64 `bson:"user_id"`
	Items  []struct {
		ItemId   int64 `bson:"item_id"`
		Quantity int64 `bson:"quantity"`
	} `bson:"items"`
}

func (r *cartRepo) GetCart(ctx context.Context, uid int64) (*biz.Cart, error) {
	result := &Cart{}
	if err := r.cartColl.FindOne(ctx, bson.M{"s": uid}).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return &biz.Cart{UserId: result.UserId}, nil
		}
		return nil, err
	}
	items := make([]biz.Item, 0)
	for _, x := range result.Items {
		items = append(items, biz.Item{
			Id:       x.ItemId,
			Quantity: x.Quantity,
		})
	}
	return &biz.Cart{UserId: result.UserId, Items: items}, nil
}

func (r *cartRepo) DeleteCart(ctx context.Context, uid int64) error {
	_, err := r.cartColl.DeleteOne(ctx, bson.M{"s": uid})
	return err
}

func (r *cartRepo) SaveCart(ctx context.Context, c *biz.Cart) error {
	items := bson.A{}
	for _, x := range c.Items {
		items = append(items, bson.M{"item_id": x.Id, "quantity": x.Quantity})
	}
	query := bson.M{"user_id": c.UserId}
	update := bson.M{"$set": bson.M{"user_id": c.UserId, "items": items}}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	result := r.cartColl.FindOneAndUpdate(ctx, query, update, opts)
	return result.Err()
}
