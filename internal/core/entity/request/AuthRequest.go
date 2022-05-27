package request

type LoginRequest struct {
	Email string `json:"email" example:"admin@web.io"`
	Password string `json:"password" example:"admin"`
}
