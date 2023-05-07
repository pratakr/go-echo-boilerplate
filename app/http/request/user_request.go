package request

type CreateUserRequest struct {
	Name     string `json:"name" validate:"nonnil"`
	Email    string `json:"email" validate:"nonnil"`
	Password string `json:"password" validate:"nonnil,min=8"`
	RoleID   int32  `json:"role_id"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"nonnil"`
	Email    string `json:"email" validate:"nonnil"`
	Password string `json:"password" validate:"nonnil,min=8"`
	RoleID   int32  `json:"role_id"`
}
