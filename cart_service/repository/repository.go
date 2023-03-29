package repository

import (
	"context"

	"gorm.io/gorm"
)

type ICartServiceRepository interface {
	AddCart(ctx context.Context, customerId string) error
	GetCart(ctx context.Context, customerId string) (*GetCartModel, error)
	AddGoods(ctx context.Context, cartId int, data []*AddGoodsData) error
	DeleteGoods(ctx context.Context, cartId int, data []*DeleteGoodsData) error
	UpdateGoods(ctx context.Context, cartId int, data []*AddGoodsData) error
	DeleteAllGoods(ctx context.Context, cartId int) error
}

type (
	AddGoodsData struct {
		GoodsId    string
		GoodsSize  string
		GoodsColor string
		Quantity   int
	}
	DeleteGoodsData struct {
		GoodsId    string
		GoodsSize  string
		GoodsColor string
	}
)

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

func (r *cartServiceRepository) AddGoods(ctx context.Context, cartId int, data []*AddGoodsData) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		for _, goods := range data {
			err = r.db.WithContext(ctx).Table(r.goodsTableName).Create(&GoodsModel{
				CartId:     cartId,
				GoodsId:    goods.GoodsId,
				GoodsSize:  goods.GoodsSize,
				GoodsColor: goods.GoodsColor,
				Quantity:   goods.Quantity,
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *cartServiceRepository) DeleteGoods(ctx context.Context, cartId int, data []*DeleteGoodsData) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		for _, goods := range data {

			err = r.db.WithContext(ctx).Table(r.goodsTableName).
				Where("cart_id = ? and goods_id= ? and goods_color = ? and goods_size = ?", cartId, goods.GoodsId,
					goods.GoodsColor, goods.GoodsSize).Delete(cartId).Error

			if err != nil {
				return err
			}
		}
		return nil
	})

}

func (r *cartServiceRepository) UpdateGoods(ctx context.Context, cartId int, data []*AddGoodsData) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := r.db.WithContext(ctx).Table(r.goodsTableName).Where("cart_id = ? ", cartId).Delete(cartId).Error
		if err != nil {
			return err
		}

		for _, goods := range data {
			err = r.db.WithContext(ctx).Table(r.goodsTableName).Create(&GoodsModel{
				CartId:     cartId,
				GoodsId:    goods.GoodsId,
				GoodsSize:  goods.GoodsSize,
				GoodsColor: goods.GoodsColor,
				Quantity:   goods.Quantity,
			}).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *cartServiceRepository) DeleteAllGoods(ctx context.Context, cartId int) error {
	return r.db.WithContext(ctx).Table(r.goodsTableName).Where("cart_id = ? ", cartId).Delete(cartId).Error
}
