package request

type CreateUserRequest struct {
	Name     string `json:"name" validate:"nonnil"`
	Email    string `json:"email" validate:"nonnil"`
	Password string `json:"password" validate:"nonnil,min=18"`
	RoleID   int32  `json:"role_id"`
}
