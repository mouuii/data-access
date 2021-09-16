package biz

import "context"

type Item struct {
	Id       int64
	Quantity int64
}

type Cart struct {
	UserId int64
	Items  []Item
}

type CartRepo interface {
	GetCart(ctx context.Context, uid int64) (*Cart, error)
	SaveCart(ctx context.Context, c *Cart) error
	DeleteCart(ctx context.Context, uid int64) error
}
