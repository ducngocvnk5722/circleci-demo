package models

import "time"

type JobType struct {
	Id        uint64    `gorm:"primary_key"`
	Title     string    `gorm:"not null;type:varchar(255)"`
	CompanyId uint64    `gorm:"not null;type:int"`
	CreatedAt time.Time `gorm:"not null;Default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"null"`
}
