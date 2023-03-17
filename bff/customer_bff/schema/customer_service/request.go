package customer_service

type GetCustomerInfoRequest struct {
	Username string
}

type UpdateCustomerInfoRequest struct {
	Username string
	Email    string
	Name     string
	Phone    string
	Gender   string
	Age      int32
	Street   string
	Ward     string
	District string
	Province string
}

type AddCustomerRequest struct {
	Email    string
	Username string
	Name     string
	Phone    string
	Gender   string
	Age      int32
	Street   string
	Ward     string
	District string
	Province string
}
