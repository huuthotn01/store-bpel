package repository

import (
	"gorm.io/gorm"
)

type customerServiceRepository struct {
	db                *gorm.DB
	customerTableName string
}

type CustomerModel struct {
	Username       string
	CustomerName   string
	CustomerEmail  string
	CustomerPhone  string
	CustomerAge    int32
	CustomerGender string
	Street         string
	Ward           string
	District       string
	Province       string
}
