package repository

import (
	"context"

	"gorm.io/gorm"
)

type IAccountServiceRepository interface {
	GetListAccount(ctx context.Context, username string) ([]*AccountModel, error)
	GetAccount(ctx context.Context, username string) (*AccountModel, error)
	AddAccount(ctx context.Context, data *AccountModel) error
	UpdateRole(ctx context.Context, username string, role int, password string) error
	UpdatePassword(ctx context.Context, password string) error
}

func NewRepository(db *gorm.DB) IAccountServiceRepository {
	return &accountServiceRepository{
		db:               db,
		accountTableName: "account",
	}
}

func (r *accountServiceRepository) GetListAccount(ctx context.Context, username string) ([]*AccountModel, error) {
	var result []*AccountModel
	query := r.db.WithContext(ctx).Table(r.accountTableName).Where("is_activated=1")
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
	if data.UserRole == 7 {
		return r.db.Exec("INSERT INTO `account` (`username`, `password`, `user_role`, `email`, `is_activated`) VALUES (?, ?, ?, ?, '0');",
			data.Username, data.Password, data.UserRole, data.Email).Error
	}
	return r.db.Exec("INSERT INTO `account` (`username`, `password`, `user_role`, `email`, `is_activated`) VALUES (?, ?, ?, ?, '1');",
		data.Username, data.Password, data.UserRole, data.Email).Error
}

func (r *accountServiceRepository) UpdateRole(ctx context.Context, username string, role int, password string) error {
	var err error
	if role == 7 {
		err = r.db.Exec("UPDATE `account` SET `user_role` = ?, `is_activated` = '0' WHERE `account`.`username` = ?;", role, username).Error
	} else {
		if password != "" {
			err = r.db.Exec("UPDATE `account` SET `user_role` = ?, `is_activated` = '1', `password` = ? WHERE `account`.`username` = ?", role, password, username).Error
		} else {
			err = r.db.Exec("UPDATE `account` SET `user_role` = ?, `is_activated` = '1' WHERE `account`.`username` = ?;", role, username).Error
		}
	}

	return err
}

func (r *accountServiceRepository) UpdatePassword(ctx context.Context, password string) error {

	return r.db.Exec("UPDATE `account` SET `password` = ?;", password).Error
}
