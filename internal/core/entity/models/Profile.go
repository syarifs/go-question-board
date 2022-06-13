package models

type Profile struct {
	ID uint `json:"id" gorm:"primaryKey"`
	FullName string `json:"full_name"`
	Gender string `json:"gender"`
	BirthDate string `json:"birthdate"`
	UserID uint `json:"user_id"`
	User User `json:"user"`
}
