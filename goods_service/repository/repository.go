package repository

import (
	"context"
	"gorm.io/gorm"
)

type IGoodsServiceRepository interface {
	GetGoods(ctx context.Context) ([]*GoodsModel, error)
	GetDetailGoods(ctx context.Context, goodsId string) (*GoodsModel, error)
	AddGoods(ctx context.Context, data *GoodsModel) error
	UpdateGoods(ctx context.Context, data *GoodsModel) error
	UpdateGoodsIsForSaleToNo(ctx context.Context, goodsId string) error
	GetGoodsInWHData(ctx context.Context, data *GoodsInWh) ([]*GoodsInWh, error)
	UpdateGoodsInWHInOut(ctx context.Context, data *GoodsInWh) error
	UpdateGoodsInWHTransfer(ctx context.Context, data *GoodsInWh, fromWH, toWH string) error
}

func NewRepository(db *gorm.DB) IGoodsServiceRepository {
	return &goodsServiceRepository{
		db:                 db,
		goodsTableName:     "goods",
		goodsImgTableName:  "goods_img",
		goodsInWhTableName: "goods_in_wh",
	}
}

func (r *goodsServiceRepository) GetGoods(ctx context.Context) ([]*GoodsModel, error) {
	var result []*GoodsModel
	query := r.db.WithContext(ctx).Table(r.goodsTableName).Find(&result)
	return result, query.Error
}

func (r *goodsServiceRepository) GetDetailGoods(ctx context.Context, goodsId string) (*GoodsModel, error) {
	var result *GoodsModel
	query := r.db.WithContext(ctx).Table(r.goodsTableName).Where("goods_code = ?", goodsId).First(&result)
	return result, query.Error
}

func (r *goodsServiceRepository) AddGoods(ctx context.Context, data *GoodsModel) error {
	return r.db.WithContext(ctx).Table(r.goodsTableName).Create(&data).Error
}

func (r *goodsServiceRepository) UpdateGoods(ctx context.Context, data *GoodsModel) error {
	return r.db.WithContext(ctx).Table(r.goodsTableName).Where("goods_code = ?", data.GoodsCode).Updates(&data).Error
}

func (r *goodsServiceRepository) UpdateGoodsIsForSaleToNo(ctx context.Context, goodsId string) error {
	return r.db.WithContext(ctx).Table(r.goodsTableName).Where("goods_code = ?", goodsId).Update("is_for_sale", 0).Error
}

func (r *goodsServiceRepository) GetGoodsInWHData(ctx context.Context, data *GoodsInWh) ([]*GoodsInWh, error) {
	var result []*GoodsInWh
	query := r.db.WithContext(ctx).Table(r.goodsInWhTableName)
	if data.GoodsCode != "" {
		query = query.Where("goods_code = ?", data.GoodsCode)
	}
	if data.GoodsSize != "" {
		query = query.Where("goods_size = ?", data.GoodsSize)
	}
	if data.GoodsColor != "" {
		query = query.Where("goods_color = ?", data.GoodsColor)
	}
	if data.WhCode != "" {
		query = query.Where("wh_code = ?", data.WhCode)
	}
	return result, query.Find(&result).Error
}

func (r *goodsServiceRepository) UpdateGoodsInWHInOut(ctx context.Context, data *GoodsInWh) error {
	return r.db.WithContext(ctx).
		Exec("UPDATE `goods_in_wh` SET `quantity` = `quantity` + ? WHERE `goods_code` = ? AND `goods_size` = ? AND `goods_color` = ? AND `wh_code` = ?",
			data.Quantity, data.GoodsCode, data.GoodsSize, data.GoodsColor, data.WhCode).Error
}

func (r *goodsServiceRepository) UpdateGoodsInWHTransfer(ctx context.Context, data *GoodsInWh, fromWH, toWH string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// decrease from WH
		err := tx.Exec("UPDATE `goods_in_wh` SET `quantity` = `quantity` - ? WHERE `goods_code` = ? AND `goods_size` = ? AND `goods_color` = ? AND `wh_code` = ?",
			data.Quantity, data.GoodsCode, data.GoodsSize, data.GoodsColor, fromWH).Error
		if err != nil {
			return err
		}
		// increase to WH
		return tx.Exec("UPDATE `goods_in_wh` SET `quantity` = `quantity` + ? WHERE `goods_code` = ? AND `goods_size` = ? AND `goods_color` = ? AND `wh_code` = ?",
			data.Quantity, data.GoodsCode, data.GoodsSize, data.GoodsColor, toWH).Error
	})
}
