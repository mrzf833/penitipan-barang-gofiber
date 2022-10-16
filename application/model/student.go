package model

type Student struct {
	Id          int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Fullname    string `gorm:"not null;size:225" json:"fullname"`
	Email       string `gorm:"not null;size:225|uniqueIndex" json:"email"`
	PhoneNumber string `grom:"not null;size:15" json:"phone_number"`
	Address     string `grom:"null;" json:"address"`
}
