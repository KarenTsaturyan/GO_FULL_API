package auth

// DTO's
type LoginRequest struct {
	Email string `json:"email" validate:"required,email"`
	// if I send passwords or any different key or no key
	// in body, I will get zero value
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	// for json it will map it like token:...
	Token string `json:"token"`
}

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
