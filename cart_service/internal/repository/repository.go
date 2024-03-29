package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type ICartServiceRepository interface {
	AddCart(ctx context.Context, customerId string) error
	GetCart(ctx context.Context, customerId string) (*GetCartModel, error)
	AddGoods(ctx context.Context, cartId string, data []*AddGoodsData) error
	DeleteGoods(ctx context.Context, cartId string, data []*DeleteGoodsData) error
	UpdateGoods(ctx context.Context, cartId string, data []*AddGoodsData) error
	DeleteAllGoods(ctx context.Context, cartId string) error
}

type (
	AddGoodsData struct {
		GoodsId     string
		GoodsSize   string
		GoodsColor  string
		Quantity    int
		MaxQuantity int
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
		CartId:     customerId,
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

func (r *cartServiceRepository) AddGoods(ctx context.Context, cartId string, data []*AddGoodsData) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		for _, goods := range data {
			var goodsModel *GoodsModel
			err = tx.WithContext(ctx).Table(r.goodsTableName).Where("goods_id = ? AND goods_size = ? AND goods_color = ?",
				goods.GoodsId, goods.GoodsSize, goods.GoodsColor).First(&goodsModel).Error

			if err != nil {
				err = tx.WithContext(ctx).Table(r.goodsTableName).Create(&GoodsModel{
					CartId:     cartId,
					GoodsId:    goods.GoodsId,
					GoodsSize:  goods.GoodsSize,
					GoodsColor: goods.GoodsColor,
					Quantity:   goods.Quantity,
				}).Error
				if err != nil {
					return err
				}
			} else {
				if (goodsModel.Quantity + goods.Quantity) > goods.MaxQuantity {
					return errors.New("quantity limit exceeded")
				}
				err = tx.Exec("UPDATE `goods` SET `quantity` = `quantity` + ? WHERE `cart_id` =? AND `goods_id` = ? AND `goods_size` = ? AND `goods_color` = ?",
					goods.Quantity, cartId, goods.GoodsId, goods.GoodsSize, goods.GoodsColor).Error
				if err != nil {
					return err
				}
			}

		}
		return nil
	})
}

func (r *cartServiceRepository) DeleteGoods(ctx context.Context, cartId string, data []*DeleteGoodsData) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var err error
		for _, goods := range data {

			err = tx.WithContext(ctx).Table(r.goodsTableName).
				Where("cart_id = ? and goods_id= ? and goods_color = ? and goods_size = ?", cartId, goods.GoodsId,
					goods.GoodsColor, goods.GoodsSize).Delete(cartId).Error

			if err != nil {
				return err
			}
		}
		return nil
	})

}

func (r *cartServiceRepository) UpdateGoods(ctx context.Context, cartId string, data []*AddGoodsData) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.WithContext(ctx).Table(r.goodsTableName).Where("cart_id = ? ", cartId).Delete(cartId).Error
		if err != nil {
			return err
		}

		for _, goods := range data {
			err = tx.WithContext(ctx).Table(r.goodsTableName).Create(&GoodsModel{
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

func (r *cartServiceRepository) DeleteAllGoods(ctx context.Context, cartId string) error {
	return r.db.WithContext(ctx).Table(r.goodsTableName).Where("cart_id = ? ", cartId).Delete(cartId).Error
}
