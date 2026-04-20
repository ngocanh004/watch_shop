package models

import "time"

type Category struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	TenDanhMuc string    `gorm:"size:100;not null" json:"ten_danh_muc"`
	MoTa       string    `gorm:"type:text" json:"mo_ta"`
	Slug       string    `gorm:"size:120;uniqueIndex" json:"slug"`
	TrangThai  bool      `gorm:"default:true" json:"trang_thai"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Category) TableName() string {
	return "categories"
}
