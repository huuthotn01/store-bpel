package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type IGoodsServiceRepository interface {
	FilterGoods(ctx context.Context, name string, newAdded bool) ([]string, error)
	GetGoods(ctx context.Context) ([]*GoodsModel, error)
	GetGoodsDefault(ctx context.Context, limit, offset int) ([]string, error)
	GetImages(ctx context.Context) ([]*GoodsImg, error)
	GetGoodsImages(ctx context.Context, goodsId string) ([]*GoodsImg, error)
	GetGoodsImageUrls(ctx context.Context, goodsId string) ([]string, error)
	GetDetailGoods(ctx context.Context, goodsId string) ([]*GoodsModel, error)
	AddGoods(ctx context.Context, data []*GoodsModel) error
	UpdateGoods(ctx context.Context, data []*GoodsModel) error
	UpdateGoodsIsForSaleToNo(ctx context.Context, goodsId string) error
	GetGoodsInWHData(ctx context.Context, data *GoodsInWh) ([]*GoodsInWh, error)
	UpdateGoodsInWHInOut(ctx context.Context, data *GoodsInWh) error
	UpdateGoodsInWHTransfer(ctx context.Context, data *GoodsInWh, fromWH, toWH string) error
	GetWarehouseByGoods(ctx context.Context, goodsId string) ([]*GoodsInWh, error)
	AddGoodsImage(ctx context.Context, data *GoodsImg) error
	DeleteGoodsImage(ctx context.Context, url string) error
}

func NewRepository(db *gorm.DB) IGoodsServiceRepository {
	return &goodsServiceRepository{
		db:                 db,
		goodsTableName:     "goods",
		goodsImgTableName:  "goods_img",
		goodsInWhTableName: "goods_in_wh",
	}
}

type FilterGoodsData struct {
	GoodsCode string
	CreatedAt string
}

func (r *goodsServiceRepository) FilterGoods(ctx context.Context, name string, newAdded bool) ([]string, error) {
	var (
		filteredData []*FilterGoodsData
		result       []string
	)
	query := r.db.WithContext(ctx).Table(r.goodsTableName).Where("goods_name like '%" + name + "%'").Select("distinct goods_code, created_at")
	if newAdded {
		query = query.Where("created_at >= date_sub(curdate(), interval 30 day)").Order("created_at desc")
	}
	err := query.Find(&filteredData).Error
	if err != nil {
		return nil, err
	}

	for _, data := range filteredData {
		result = append(result, data.GoodsCode)
	}

	return result, nil
}

func (r *goodsServiceRepository) GetGoodsDefault(ctx context.Context, pageSize, pageNum int) ([]string, error) {
	var (
		limit         = pageSize
		offset        = (pageNum - 1) * pageSize
		goodsCodeList []string
	)
	query := r.db.WithContext(ctx).Table(r.goodsTableName).Limit(limit).Offset(offset).Distinct("goods_code").Find(&goodsCodeList)
	return goodsCodeList, query.Error
}

func (r *goodsServiceRepository) GetImages(ctx context.Context) ([]*GoodsImg, error) {
	var result []*GoodsImg
	query := r.db.WithContext(ctx).Table(r.goodsImgTableName).Find(&result)
	return result, query.Error
}

func (r *goodsServiceRepository) GetGoodsImages(ctx context.Context, goodsId string) ([]*GoodsImg, error) {
	var result []*GoodsImg
	query := r.db.WithContext(ctx).Table(r.goodsImgTableName).Where("goods_code = ?", goodsId).Find(&result)
	return result, query.Error
}

func (r *goodsServiceRepository) GetGoodsImageUrls(ctx context.Context, goodsId string) ([]string, error) {
	var result []string
	query := r.db.WithContext(ctx).Table(r.goodsImgTableName).Where("goods_code = ?", goodsId).Select("goods_img").Order("is_default DESC").Find(&result)
	return result, query.Error
}

func (r *goodsServiceRepository) GetGoods(ctx context.Context) ([]*GoodsModel, error) {
	var result []*GoodsModel
	query := r.db.WithContext(ctx).Table(r.goodsTableName)
	return result, query.Find(&result).Error
}

func (r *goodsServiceRepository) GetDetailGoods(ctx context.Context, goodsId string) ([]*GoodsModel, error) {
	var result []*GoodsModel
	query := r.db.WithContext(ctx).Table(r.goodsTableName).Where("goods_code = ?", goodsId).Find(&result)
	return result, query.Error
}

func (r *goodsServiceRepository) AddGoods(ctx context.Context, data []*GoodsModel) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		for _, good := range data {
			err := r.db.WithContext(ctx).Table(r.goodsTableName).Create(&good).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *goodsServiceRepository) UpdateGoods(ctx context.Context, data []*GoodsModel) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		goods := data[0]
		err := r.db.WithContext(ctx).Table(r.goodsTableName).Where("goods_code = ?", goods.GoodsCode).Delete(&goods).Error
		if err != nil {
			return err
		}
		for _, item := range data {
			err := r.db.WithContext(ctx).Table(r.goodsTableName).Create(&item).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
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
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var currGoods []*GoodsInWh
		err := tx.Table(r.goodsInWhTableName).
			Where("goods_code = ? and goods_size = ? and goods_color = ? and wh_code = ?", data.GoodsCode, data.GoodsSize, data.GoodsColor, data.WhCode).
			Find(&currGoods).Error
		if err != nil {
			return err
		}

		if len(currGoods) == 0 {
			if data.Quantity < 0 { // export but not have item => return error
				return errors.New("not have goods to export")
			} else { // import but currently not have goods in destination warehouse
				newData := &GoodsInWh{
					GoodsCode:  data.GoodsCode,
					GoodsColor: data.GoodsColor,
					GoodsSize:  data.GoodsSize,
					WhCode:     data.WhCode,
					Quantity:   data.Quantity,
				}
				return tx.Table(r.goodsInWhTableName).Select("goods_code", "goods_size", "goods_color", "wh_code", "quantity").Create(newData).Error
			}
		}

		if currGoods[0].Quantity+data.Quantity < 0 {
			// not have enough items to export => return error
			return errors.New("not have enough goods number to export")
		}

		err = tx.Exec("UPDATE `goods_in_wh` SET `quantity` = `quantity` + ? WHERE `goods_code` = ? AND `goods_size` = ? AND `goods_color` = ? AND `wh_code` = ?",
			data.Quantity, data.GoodsCode, data.GoodsSize, data.GoodsColor, data.WhCode).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func (r *goodsServiceRepository) UpdateGoodsInWHTransfer(ctx context.Context, data *GoodsInWh, fromWH, toWH string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// decrease from WH
		err := tx.Exec("UPDATE `goods_in_wh` SET `quantity` = `quantity` - ? WHERE `goods_code` = ? AND `goods_size` = ? AND `goods_color` = ? AND `wh_code` = ?",
			data.Quantity, data.GoodsCode, data.GoodsSize, data.GoodsColor, fromWH).Error
		if err != nil {
			return err
		}
		var goodsModel *GoodsModel
		err = tx.Table(r.goodsInWhTableName).Where("goods_code = ? and goods_size = ? and goods_color = ? and wh_code = ?", data.GoodsCode, data.GoodsSize, data.GoodsColor, toWH).First(&goodsModel).Error

		if err != nil {
			// create in WH
			return tx.Exec("INSERT `goods_in_wh` (`goods_code`, `goods_size`, `goods_color`, `wh_code`, `quantity`) VALUE (?, ?, ?, ?, ?)",
				data.GoodsCode, data.GoodsSize, data.GoodsColor, toWH, data.Quantity).Error
		}
		// increase to WH
		return tx.Exec("UPDATE `goods_in_wh` SET `quantity` = `quantity` + ? WHERE `goods_code` = ? AND `goods_size` = ? AND `goods_color` = ? AND `wh_code` = ?",
			data.Quantity, data.GoodsCode, data.GoodsSize, data.GoodsColor, toWH).Error

	})
}

func (r *goodsServiceRepository) GetWarehouseByGoods(ctx context.Context, goodsId string) ([]*GoodsInWh, error) {
	var result []*GoodsInWh
	query := r.db.WithContext(ctx).Table(r.goodsInWhTableName).Where("goods_code = ?", goodsId).Find(&result)
	return result, query.Error
}

func (r *goodsServiceRepository) AddGoodsImage(ctx context.Context, data *GoodsImg) error {
	return r.db.WithContext(ctx).Table(r.goodsImgTableName).Create(&data).Error
}

func (r *goodsServiceRepository) DeleteGoodsImage(ctx context.Context, url string) error {
	return r.db.WithContext(ctx).Table(r.goodsImgTableName).Where("goods_img = ?", url).Delete(url).Error
}
