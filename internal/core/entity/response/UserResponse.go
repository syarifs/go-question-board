package response

type (
	User struct {
		ID uint `json:"id"`
		Email string `json:"email"`
		FullName string `json:"full_name"`
		Password string `json:"-"`
		Gender string `json:"gender"`
		Birthdate string `json:"birthdate"`
		Role string `json:"roles"`
		Status int `json:"status"`
	}
)
