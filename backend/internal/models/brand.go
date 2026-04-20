package models

import "time"

type Brand struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	TenThuongHieu string    `gorm:"size:100;not null" json:"ten_thuong_hieu"`
	MoTa          string    `gorm:"type:text" json:"mo_ta"`
	Logo          string    `gorm:"size:255" json:"logo"`
	QuocGia       string    `gorm:"size:100" json:"quoc_gia"`
	TrangThai     bool      `gorm:"default:true" json:"trang_thai"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (Brand) TableName() string {
	return "brands"
}
