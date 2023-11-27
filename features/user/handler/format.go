package handler

type UserRequest struct {
	Nama     string `json:"nama" form:"nama"`
	Password string `json:"password" form:"password"`
}

type UserResponse struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
}

type LoginRequest struct {
	Nama     string `json:"nama" form:"nama"`
	Password string `json:"password" form:"password"`
}

type LoginResponse struct {
	ID    uint   `json:"id"`
	Nama  string `json:"nama"`
	Token string `json:"token"`
}
