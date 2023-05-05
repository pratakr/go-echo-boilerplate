package request

type AuthLoginRequest struct {
	Email      string `json:"email" validate:"required"`
	Password   string `json:"password"  validate:"required"`
	RememberMe bool   `json:"remember_me"`
}
