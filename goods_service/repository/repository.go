package repository

import (
	"context"
	"gorm.io/gorm"
)

type IGoodsServiceRepository interface {
	GetGoods(ctx context.Context) ([]*GoodsModel, error)
}

func NewRepository(db *gorm.DB) IGoodsServiceRepository {
	return &goodsServiceRepository{
		db: db,
		goodsTableName: "goods",
		goodsImgTableName: "goods_img",
		goodsInWhTableName: "goods_in_wh",
	}
}

func (r *goodsServiceRepository) GetGoods(ctx context.Context) ([]*GoodsModel, error) {
	var result []*GoodsModel
	query := r.db.WithContext(ctx).Table(r.goodsTableName).Find(&result)
	return result, query.Error
}