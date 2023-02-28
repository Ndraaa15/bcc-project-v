package entities

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          uint    `json:"ID" gorm:"primaryKey" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	AdminID     uint    `json:"admin_id" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Stock       uint    `json:"stock" binding:"required"`
}
