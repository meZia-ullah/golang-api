package models

type Education struct {
	ID          uint   `gorm:"primaryKey"`
	StudentName string `gorm:"size:100;not null"`
	Course      string `gorm:"size:100;not null"`
	Grade       string `gorm:"size:10;not null"`
}
