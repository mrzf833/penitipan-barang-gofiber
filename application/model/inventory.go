package model

import (
	"database/sql"
	"time"
)

type Inventory struct {
	Id          int            `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryId  sql.NullInt64  `gorm:"null" json:"category_id"`
	Category    Category       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:CategoryId;foreignKey:Id;"`
	DepositName sql.NullString `gorm:"null" json:"deposit_name"`

	DepositStudentId sql.NullInt64 `gorm:"not null" json:"deposit_student_id"`
	DepositStudent   Student       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:DepositStudentId;foreignKey:Id;"`

	DepositAdmin     int  `gorm:"not null" json:"deposit_admin"`
	DepositUserAdmin User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:DepositAdmin;foreignKey:Id;"`

	DepositTime time.Time      `gorm:"not null" json:"deposit_time"`
	ItemName    string         `gorm:"not null" json:"item_name"`
	Description sql.NullString `gorm:"null" json:"description"`
	Status      string         `gorm:"type:enum('deposit','take');not null" json:"status"`
	TakeName    sql.NullString `gorm:"null" json:"take_name"`

	TakeStudentId sql.NullInt64 `gorm:"null" json:"take_student_id"`
	TakeStudent   Student       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;references:TakeStudentId;foreignKey:Id;"`

	TakeTime      sql.NullTime  `gorm:"null" json:"take_time"`
	TakeAdmin     sql.NullInt64 `gorm:"null" json:"take_admin"`
	TakeUserAdmin User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;references:TakeAdmin;foreignKey:Id;"`
}
