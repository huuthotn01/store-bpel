package account_service

type ChangePasswordRequest struct {
	Username    string
	OldPassword string
	NewPassword string
}
