package account_service

type GetListAccountRequest struct {
	Username string // for filtering
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
