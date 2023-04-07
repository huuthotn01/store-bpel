package order_service

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
}

type GetListOrderCustomerResponseData struct {
	OrderId         int    // private code
	OrderCode       string // public code
	PaymentMethod   string
	ListGoods       []*OrderGoodsResponse
	TotalPrice      int
	TotalGoods      int
	TotalDiscount   int
	TotalOrder      int
	IsCompleted     bool
	ShipFee         int
	StatusShip      []*GetListOrderStateResponse
	TransactionDate string
	ExpectDate      string
}

type GetOrderDetailCustomerResponseData struct {
	OrderId         int    // private code
	OrderCode       string // public code
	PaymentMethod   string
	ListGoods       []*OrderGoodsResponse
	TotalPrice      int
	TotalGoods      int
	TotalDiscount   int
	TotalOrder      int
	IsCompleted     bool
	ShipFee         int
	StatusShip      []*GetListOrderStateResponse
	TransactionDate string
	Status          int
	NameReceiver    string
	PhoneReceiver   string
	EmailReceiver   string
	Address         *Address
	ExpectDate      string
}

type OrderGoodsResponse struct {
	GoodsId   string
	Image     string
	Name      string
	UnitPrice int
	Price     int
	Quantity  int
	Size      string
	Color     string
	Discount  float32
}

type GetListOrderStateResponse struct {
	State string
	Time  string
}

type GetShipFeeResponseData struct {
	ShipFee      int
	ExpectedDate string
}

type GetOnlineOrdersStatusResponseData struct {
	OrderId   int
	State     string
	StateTime string
}
