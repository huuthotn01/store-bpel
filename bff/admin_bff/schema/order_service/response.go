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

type GetOrderDetailResponseData struct {
	OrderId          int
	OrderCode        string
	ListGoods        []*OrderGoodsResponse
	TotalPrice       int
	TotalGoods       int
	TotalDiscount    int
	TotalOrder       int
	TransactionDate  string
	IsOnline         bool
	OnlineOrderData  *OnlineOrderData
	OfflineOrderData *OfflineOrderData
}

type OrderGoodsResponse struct {
	GoodsId   string
	Image     string
	Name      string
	UnitPrice int
	Price     int
	Tax       float32
	Quantity  int
	Size      string
	Color     string
	Discount  float32
}

type GetOnlineOrdersResponseData struct {
	OrderId         int
	OrderCode       string
	ListGoods       []*OrderGoodsResponse
	TotalPrice      int
	TotalGoods      int
	TotalDiscount   int
	TotalOrder      int
	TransactionDate string
	OnlineOrderData *OnlineOrderData
}

type GetOfflineOrdersResponseData struct {
	OrderId          int
	OrderCode        string
	ListGoods        []*OrderGoodsResponse
	TotalPrice       int
	TotalGoods       int
	TotalDiscount    int
	TotalOrder       int
	TransactionDate  string
	OfflineOrderData *OfflineOrderData
}

type OnlineOrderData struct {
	PaymentMethod string
	CustomerId    string
	IsCompleted   bool
	ShipFee       int
	ExpectDate    string
	Status        int
	NameReceiver  string
	PhoneReceiver string
	EmailReceiver string
	Address       *Address
}

type OfflineOrderData struct {
	StaffId  string
	BranchId string
}

type Address struct {
	Street   string
	Ward     string
	District string
	Province string
}
