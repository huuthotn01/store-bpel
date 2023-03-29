package repository

import (
	"context"

	"gorm.io/gorm"
)

type ICartServiceRepository interface {
	AddCart(ctx context.Context, customerId string) error
	GetCart(ctx context.Context, customerId string) (*GetCartModel, error)
}

func NewRepository(db *gorm.DB) ICartServiceRepository {
	return &cartServiceRepository{
		db:             db,
		cartTableName:  "cart",
		goodsTableName: "goods",
	}
}

func (r *cartServiceRepository) AddCart(ctx context.Context, customerId string) error {
	cartModel := &CartModel{
		CustomerId: customerId,
	}
	err := r.db.WithContext(ctx).Table(r.cartTableName).Create(cartModel).Error

	return err
}

func (r *cartServiceRepository) GetCart(ctx context.Context, customerId string) (*GetCartModel, error) {
	var cartModel *CartModel
	err := r.db.WithContext(ctx).Table(r.cartTableName).Where("customer_id = ?", customerId).First(&cartModel).Error
	if err != nil {
		return nil, err
	}

	var goodsModel []*GoodsModel
	err = r.db.WithContext(ctx).Table(r.goodsTableName).Where("cart_id = ?", cartModel.CartId).Find(&goodsModel).Error
	if err != nil {
		return nil, err
	}

	result := &GetCartModel{
		CartId: cartModel.CartId,
		Goods:  goodsModel,
	}

	return result, nil
}
