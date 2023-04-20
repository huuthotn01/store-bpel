package schema

type UpdateCustomerInfoRequest struct {
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

type UploadImageRequest struct {
	Username string
	ImageUrl string
}
