package schema

type UpdateOnlineOrdersStatusRequest struct {
	OrderId      string
	State        string
	StatusNumber int
}

type MakeOnlineOrderRequest struct {
	CustomerId      string
	PaymentMethod   string
	GoodsList       []*OrderGoodsRequest
	TotalPrice      int
	ShipFee         int
	TransactionDate string
	ExpectedDate    string
	NameReceiver    string
	PhoneReceiver   string
	EmailReceiver   string
	Address         *Address
}

type MakeOfflineOrderRequest struct {
	GoodsList       []*OrderGoodsRequest
	TotalPrice      int
	TransactionDate string
	StaffId         string
	BranchId        string
}

type OrderGoodsRequest struct {
	GoodsId   string
	UnitPrice int
	Price     int
	Name      string
	Image     string
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
