package repository

import (
	"context"
	"gorm.io/gorm"
)

type IAccountServiceRepository interface {
	GetListAccount(ctx context.Context, username string) ([]*AccountModel, error)
	GetAccount(ctx context.Context, username string) (*AccountModel, error)
	AddAccount(ctx context.Context, data *AccountModel) error
	UpdateRole(ctx context.Context, username string, role int) error
}

func NewRepository(db *gorm.DB) IAccountServiceRepository {
	return &accountServiceRepository{
		db:               db,
		accountTableName: "account",
	}
}

func (r *accountServiceRepository) GetListAccount(ctx context.Context, username string) ([]*AccountModel, error) {
	var result []*AccountModel
	query := r.db.WithContext(ctx).Table(r.accountTableName)
	if username != "" {
		query = query.Where("username LIKE ?", "%"+username+"%")
	}
	return result, query.Find(&result).Error
}

func (r *accountServiceRepository) GetAccount(ctx context.Context, username string) (*AccountModel, error) {
	var result *AccountModel
	query := r.db.WithContext(ctx).Table(r.accountTableName).Where("username = ?", username)
	return result, query.First(&result).Error
}

func (r *accountServiceRepository) AddAccount(ctx context.Context, data *AccountModel) error {
	return r.db.WithContext(ctx).Table(r.accountTableName).Select("username", "password", "user_role").Create(data).Error
}

func (r *accountServiceRepository) UpdateRole(ctx context.Context, username string, role int) error {
	return r.db.WithContext(ctx).Table(r.accountTableName).Where("username = ?", username).Update("user_role", role).Error
}
