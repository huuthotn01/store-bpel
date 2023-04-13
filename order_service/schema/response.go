package schema

type UpdateResponse struct {
	StatusCode int
	Message    string
}

type GetBestSellingGoodsResponse struct {
	StatusCode int
	Message    string
	Data       []string
}

type GetOnlineOrdersStatusResponse struct {
	StatusCode int
	Message    string
	Data       []*GetOnlineOrdersStatusResponseData
}

type GetOnlineOrdersStatusResponseData struct {
	OrderId   int
	State     string
	StateTime string
}

type GetListOrderCustomerResponse struct {
	StatusCode int
	Message    string
	Data       []*GetListOrderCustomerResponseData
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

type GetListOrderStateResponse struct {
	State string
	Time  string
}

type GetOrderDetailCustomerResponse struct {
	StatusCode int
	Message    string
	Data       *GetOrderDetailCustomerResponseData
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

type GetOrderDetailAdminResponse struct {
	StatusCode int
	Message    string
	Data       *GetOrderDetailAdminResponseData
}

type GetOrderDetailAdminResponseData struct {
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

type GetShipFeeResponse struct {
	StatusCode int
	Message    string
	Data       *GetShipFeeResponseData
}

type GetShipFeeResponseData struct {
	ShipFee      int
	ExpectedDate string
}

type GetOnlineOrdersResponse struct {
	StatusCode int
	Message    string
	Data       []*GetOnlineOrdersResponseData
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

type GetOfflineOrdersResponse struct {
	StatusCode int
	Message    string
	Data       []*GetOfflineOrdersResponseData
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
