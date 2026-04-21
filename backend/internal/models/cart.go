package models

import "time"

type CartItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"index" json:"user_id"`
	ProductID uint      `gorm:"index" json:"product_id"`
	SoLuong   int       `gorm:"default:1" json:"so_luong"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (CartItem) TableName() string {
	return "cart_items"
}

type WishlistItem struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"index" json:"user_id"`
	ProductID uint      `gorm:"index" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func (WishlistItem) TableName() string {
	return "wishlist_items"
}
