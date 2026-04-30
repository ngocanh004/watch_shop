package handlers

import (
	"strconv"

	"watch-shop/backend/internal/models"
	"watch-shop/backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CartHandler struct {
	db *gorm.DB
}

func NewCartHandler(db *gorm.DB) *CartHandler {
	return &CartHandler{db: db}
}

type cartItemInput struct {
	ProductID uint `json:"product_id"`
	SoLuong   int  `json:"so_luong"`
}

func (h *CartHandler) List(c *gin.Context) {
	uid, _ := c.Get("user_id")
	var items []models.CartItem
	h.db.Preload("Product").Preload("Product.Brand").
		Where("user_id = ?", uid).Order("id DESC").Find(&items)

	var tong float64
	for _, it := range items {
		tong += it.Product.Gia * float64(it.SoLuong)
	}
	utils.OK(c, "OK", gin.H{
		"items":     items,
		"tong_tien": tong,
		"tong_sl":   len(items),
	})
}

func (h *CartHandler) Add(c *gin.Context) {
	uid, _ := c.Get("user_id")
	var in cartItemInput
	if err := c.ShouldBindJSON(&in); err != nil || in.ProductID == 0 {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}
	if in.SoLuong < 1 {
		in.SoLuong = 1
	}

	var p models.Product
	if err := h.db.First(&p, in.ProductID).Error; err != nil {
		utils.NotFound(c, "Khong tim thay san pham")
		return
	}

	var existing models.CartItem
	err := h.db.Where("user_id = ? AND product_id = ?", uid, in.ProductID).First(&existing).Error
	if err == nil {
		existing.SoLuong += in.SoLuong
		if existing.SoLuong > p.SoLuong {
			existing.SoLuong = p.SoLuong
		}
		h.db.Save(&existing)
		utils.OK(c, "Da cap nhat gio hang", existing)
		return
	}

	item := models.CartItem{
		UserID:    uid.(uint),
		ProductID: in.ProductID,
		SoLuong:   in.SoLuong,
	}
	if err := h.db.Create(&item).Error; err != nil {
		utils.ServerError(c, "Khong the them vao gio hang")
		return
	}
	utils.Created(c, "Da them vao gio hang", item)
}

func (h *CartHandler) Update(c *gin.Context) {
	uid, _ := c.Get("user_id")
	id, _ := strconv.Atoi(c.Param("id"))
	var in cartItemInput
	if err := c.ShouldBindJSON(&in); err != nil || in.SoLuong < 1 {
		utils.BadRequest(c, "So luong khong hop le")
		return
	}
	var item models.CartItem
	if err := h.db.Where("id = ? AND user_id = ?", id, uid).First(&item).Error; err != nil {
		utils.NotFound(c, "Khong tim thay")
		return
	}
	item.SoLuong = in.SoLuong
	h.db.Save(&item)
	utils.OK(c, "Cap nhat thanh cong", item)
}

func (h *CartHandler) Remove(c *gin.Context) {
	uid, _ := c.Get("user_id")
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.db.Where("id = ? AND user_id = ?", id, uid).Delete(&models.CartItem{}).Error; err != nil {
		utils.ServerError(c, "Xoa that bai")
		return
	}
	utils.OK(c, "Da xoa khoi gio hang", nil)
}

func (h *CartHandler) Clear(c *gin.Context) {
	uid, _ := c.Get("user_id")
	h.db.Where("user_id = ?", uid).Delete(&models.CartItem{})
	utils.OK(c, "Da xoa gio hang", nil)
}
