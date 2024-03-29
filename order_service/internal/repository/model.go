package repository

import (
	"gorm.io/gorm"
	"time"
)

type orderServiceRepository struct {
	db                    *gorm.DB
	goodsTableName        string
	ordersTableName       string
	onlineOrdersTableName string
	storeOrdersTableName  string
	orderStateTableName   string
}

type GoodsModel struct {
	GoodsCode  string
	GoodsSize  string
	GoodsColor string
	GoodsName  string
	OrderCode  int
	Quantity   int
	UnitPrice  int
	TotalPrice int
	Tax        float32
	Image      string
	Promotion  float32
}

type OrdersModel struct {
	OrderCode       int `gorm:"primarykey"`
	TransactionDate string
	TotalPrice      int
	PublicOrderCode string
}

type OnlineOrdersModel struct {
	OrderCode        int
	ExpectedDelivery string
	ShippingFee      int
	CustomerId       string
	PaymentMethod    string
	Street           string
	Ward             string
	District         string
	Province         string
	CustomerName     string
	CustomerPhone    string
	CustomerEmail    string
	Status           int
}

type StoreOrdersModel struct {
	OrderCode int
	StoreCode string
	StaffId   string
}

type OrderStateModel struct {
	OrderCode int
	State     string
	StateTime time.Time
}
