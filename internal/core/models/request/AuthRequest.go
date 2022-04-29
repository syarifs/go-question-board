package request

type LoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type LevelRequest struct {
	Name string `json:"name"`
}
