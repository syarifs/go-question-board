package models


type Role struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (*Role) TableName() string {
	return "roles"
}
