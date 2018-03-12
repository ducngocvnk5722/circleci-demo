package models

import "time"

type User struct {
	Id        uint64    `gorm:"primary_key"`
	Name      string    `gorm:"not null;type:varchar(255)"`
	Email     string    `gorm:"not null;type:varchar(255)"`
	CompanyId uint64    `gorm:"not null;type:int"`
	CreatedAt time.Time `gorm:"not null;Default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"null"`
}
