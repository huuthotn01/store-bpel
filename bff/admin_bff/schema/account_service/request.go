package account_service

type GetListAccountRequest struct {
	Username string // for filtering
}

type SignUpRequest struct {
	Username string
	Password string
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

type SignInRequest struct {
	Username string
	Password string
}

type UpdateRoleRequest struct {
	Username string
	Role     int
}

type AddAccountRequest struct {
	Username string
	Password string
	Role     int
}
