package models

type User struct {
	ID           uint    `gorm:"primaryKey;autoIncrement"`
	UserCode     string  `gorm:"not null;unique;size:50"`
	UserName     string  `gorm:"not null;size:100"`
	Position     *string `gorm:"size:100"`
	Telephone    *string `gorm:"size:20"`
	Handphone    *string `gorm:"size:20"`
	Email        *string `gorm:"size:100"`
	Password     string  `gorm:"not null;size:255"`
	SecurityCode string  `gorm:"not null;size:255"`
	GroupID      uint    `gorm:"not null;default:0"`
	Status       int32   `gorm:"not null;default:0"`
	UserID       uint    `gorm:"not null;default:0"`
	LogIn        int32   `gorm:"not null;default:0"`
}

func (User) TableName() string {
	return "myuser"
}
