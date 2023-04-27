package model

import "time"

// Car represents the model for an car
type Car struct {
	ID        		uint   `gorm:"primaryKey"`
	Pemilik      	string `gorm:"not null ;unique;type:varchar(191)"`
	Merk     		string `gorm:"not null ;unique;type:varchar(191)"`
	Type     		string `gorm:"not null ;unique;type:varchar(191)"`
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}