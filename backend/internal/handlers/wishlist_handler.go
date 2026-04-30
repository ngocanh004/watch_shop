package handlers

import (
	"strconv"

	"watch-shop/backend/internal/models"
	"watch-shop/backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WishlistHandler struct {
	db *gorm.DB
}

func NewWishlistHandler(db *gorm.DB) *WishlistHandler {
	return &WishlistHandler{db: db}
}

type wishlistInput struct {
	ProductID uint `json:"product_id"`
}

func (h *WishlistHandler) List(c *gin.Context) {
	uid, _ := c.Get("user_id")
	var items []models.WishlistItem
	h.db.Preload("Product").Preload("Product.Brand").
		Where("user_id = ?", uid).Order("id DESC").Find(&items)
	utils.OK(c, "OK", items)
}

func (h *WishlistHandler) Add(c *gin.Context) {
	uid, _ := c.Get("user_id")
	var in wishlistInput
	if err := c.ShouldBindJSON(&in); err != nil || in.ProductID == 0 {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}
	var count int64
	h.db.Model(&models.WishlistItem{}).Where("user_id = ? AND product_id = ?", uid, in.ProductID).Count(&count)
	if count > 0 {
		utils.BadRequest(c, "San pham da co trong danh sach yeu thich")
		return
	}
	item := models.WishlistItem{UserID: uid.(uint), ProductID: in.ProductID}
	if err := h.db.Create(&item).Error; err != nil {
		utils.ServerError(c, "Khong the them")
		return
	}
	utils.Created(c, "Da them vao yeu thich", item)
}

func (h *WishlistHandler) Remove(c *gin.Context) {
	uid, _ := c.Get("user_id")
	pid, _ := strconv.Atoi(c.Param("product_id"))
	h.db.Where("user_id = ? AND product_id = ?", uid, pid).Delete(&models.WishlistItem{})
	utils.OK(c, "Da xoa khoi yeu thich", nil)
}
