package account_service

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

type CreateResetPasswordRequest struct {
	Username string
}

type ConfirmOTPRequest struct {
	Username string
	Otp      string
}
