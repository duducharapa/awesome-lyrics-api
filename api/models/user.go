package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Email    string `json:"email"`
	Password string `json:"-"`
}
