package app

type LoginRequest struct {
	Email    string `json:"email"`
	Passowrd string `json:"password"`
}

const _USER_NOT_FOUND_ = "User not found"
