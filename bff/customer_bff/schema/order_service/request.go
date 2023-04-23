package order_service

type GetOnlineOrdersStatusRequest struct {
	OrderId string
}

type GetListOrderCustomerRequest struct {
	CustomerId string
}

type GetOrderDetailCustomerRequest struct {
	OrderId string
}

type MakeOnlineOrderRequest struct {
	CustomerId      string
	PaymentMethod   string
	ListElements    []*ElementsGoods
	TotalPrice      int
	ShipFee         int
	TransactionDate string
	ExpectedDate    string
	NameReceiver    string
	PhoneReceiver   string
	EmailReceiver   string
	Address         *Address
}

type ElementsGoods struct {
	Elements *OrderGoodsRequest
}

type OrderGoodsRequest struct {
	GoodsCode  string
	GoodsColor string
	GoodsSize  string
	UnitPrice  int
	Price      int
	Name       string
	Image      string
	Quantity   int
	Discount   float32
	Tax        float32
}

type Address struct {
	Street   string
	Ward     string
	District string
	Province string
}
