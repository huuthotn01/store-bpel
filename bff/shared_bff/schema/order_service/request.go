package order_service

type UpdateOnlineOrdersStatusRequest struct {
	OrderId      string
	State        string
	StatusNumber int
}

type Address struct {
	Street   string
	Ward     string
	District string
	Province string
}
