package order_service

type GetOrderDetailRequest struct {
	OrderId int
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
	Quantity  int
	Size      string
	Color     string
	Discount  float32
	Tax       float32
}

type GetListOrderCustomerRequest struct {
	CustomerId string
}
