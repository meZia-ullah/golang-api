package models

import (
	"time"

	"gorm.io/gorm"
)

type Staff struct {
	ID        uint    `gorm:"primaryKey"`
	ShopID    uint    `gorm:"not null;index:idx_shop_email,unique"`
	RoleID    uint    `gorm:"not null"`
	FirstName string  `gorm:"size:50;not null"`
	LastName  string  `gorm:"size:50;not null"`
	Email     string  `gorm:"size:100;not null;index:idx_shop_email,unique"`
	Phone     *string `gorm:"size:20"`
	Password  string  `gorm:"size:255;not null"`
	IsActive  bool    `gorm:"default:true"`
	CreatedBy *uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
