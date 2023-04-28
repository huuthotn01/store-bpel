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
	Email    string
	Role     int
}

type ChangePasswordRequest struct {
	Username    string
	OldPassword string
	NewPassword string
}
