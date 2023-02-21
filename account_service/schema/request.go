package schema

type SignInRequest struct {
	Username string
	Password string
}

type UpdateRoleRequest struct {
	Role int
}

type AddAccountRequest struct {
	Username string
	Password string
	Role     int
}
