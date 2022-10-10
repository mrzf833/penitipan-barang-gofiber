package model

type User struct {
	Id       int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"not null;size:100" json:"name"`
	Username string `gorm:"uniqueIndex;not null;size:100" json:"username"`
	Password string `gorm:"size:240" json:"password"`
	Role     string `gorm:"type:enum('super_admin','admin');not null" json:"role"`
}
