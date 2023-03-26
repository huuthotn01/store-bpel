package schema

type UpdateOnlineOrdersStatusRequest struct {
	OrderId int
	State   string
}

type MakeOnlineOrderRequest struct {
	CustomerId      string
	PaymentMethod   string
	GoodsList       []*OrderGoodsRequest
	TotalPrice      int
	ShipFee         int
	TransactionDate string
	NameReceiver    string
	PhoneReceiver   string
	EmailReceiver   string
	Address         *Address
}

type OrderGoodsRequest struct {
	GoodsId   string
	UnitPrice int
	Price     int
	Quantity  int
	Size      string
	Color     string
	Discount  float32
	Tax       float32
}

type Address struct {
	Street   string
	Ward     string
	District string
	Province string
}

type GetShipFeeRequest struct {
	Street   string
	Ward     string
	District string
	Province string
}
