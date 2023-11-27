package handler

type CouponRequest struct {
	NamaProgram string `json:"nama_program" form:"nama_program"`
	Link        string `json:"link" form:"link"`
}

type CouponResponse struct {
	ID          uint   `json:"id"`
	NamaProgram string `json:"nama_program" form:"nama_program"`
	Link        string `json:"link" form:"link"`
	Gambar      string `json:"gambar" form:"gambar"`
	UserID      uint   `json:"user_id" form:"user_id"`
}
