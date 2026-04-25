package handlers

import (
	"strconv"

	"watch-shop/backend/internal/models"
	"watch-shop/backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	db *gorm.DB
}

func NewCategoryHandler(db *gorm.DB) *CategoryHandler {
	return &CategoryHandler{db: db}
}

type categoryInput struct {
	TenDanhMuc string `json:"ten_danh_muc"`
	MoTa       string `json:"mo_ta"`
	TrangThai  *bool  `json:"trang_thai"`
}

func (h *CategoryHandler) List(c *gin.Context) {
	var items []models.Category
	q := h.db.Order("id DESC")
	if kw := c.Query("keyword"); kw != "" {
		q = q.Where("ten_danh_muc ILIKE ?", "%"+kw+"%")
	}
	q.Find(&items)
	utils.OK(c, "OK", items)
}

func (h *CategoryHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var item models.Category
	if err := h.db.First(&item, id).Error; err != nil {
		utils.NotFound(c, "Khong tim thay danh muc")
		return
	}
	utils.OK(c, "OK", item)
}

func (h *CategoryHandler) Create(c *gin.Context) {
	var in categoryInput
	if err := c.ShouldBindJSON(&in); err != nil || in.TenDanhMuc == "" {
		utils.BadRequest(c, "Vui long nhap ten danh muc")
		return
	}
	item := models.Category{
		TenDanhMuc: in.TenDanhMuc,
		MoTa:       in.MoTa,
		Slug:       utils.MakeSlug(in.TenDanhMuc),
		TrangThai:  true,
	}
	if in.TrangThai != nil {
		item.TrangThai = *in.TrangThai
	}
	if err := h.db.Create(&item).Error; err != nil {
		utils.ServerError(c, "Khong the them danh muc")
		return
	}
	utils.Created(c, "Them danh muc thanh cong", item)
}

func (h *CategoryHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var item models.Category
	if err := h.db.First(&item, id).Error; err != nil {
		utils.NotFound(c, "Khong tim thay danh muc")
		return
	}
	var in categoryInput
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}
	if in.TenDanhMuc != "" {
		item.TenDanhMuc = in.TenDanhMuc
		item.Slug = utils.MakeSlug(in.TenDanhMuc)
	}
	item.MoTa = in.MoTa
	if in.TrangThai != nil {
		item.TrangThai = *in.TrangThai
	}
	h.db.Save(&item)
	utils.OK(c, "Cap nhat thanh cong", item)
}

func (h *CategoryHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.db.Delete(&models.Category{}, id).Error; err != nil {
		utils.ServerError(c, "Xoa that bai")
		return
	}
	utils.OK(c, "Xoa thanh cong", nil)
}
