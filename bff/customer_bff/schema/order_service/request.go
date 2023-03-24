package order_service

type GetOnlineOrdersStatusRequest struct {
	OrderId int
}

type UpdateOnlineOrdersStatusRequest struct {
	OrderId int
	State   string
}
