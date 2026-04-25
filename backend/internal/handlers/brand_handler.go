package handlers

import (
	"strconv"

	"watch-shop/backend/internal/models"
	"watch-shop/backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BrandHandler struct {
	db *gorm.DB
}

func NewBrandHandler(db *gorm.DB) *BrandHandler {
	return &BrandHandler{db: db}
}

type brandInput struct {
	TenThuongHieu string `json:"ten_thuong_hieu"`
	MoTa          string `json:"mo_ta"`
	Logo          string `json:"logo"`
	QuocGia       string `json:"quoc_gia"`
	TrangThai     *bool  `json:"trang_thai"`
}

func (h *BrandHandler) List(c *gin.Context) {
	var items []models.Brand
	q := h.db.Order("id DESC")
	if kw := c.Query("keyword"); kw != "" {
		q = q.Where("ten_thuong_hieu ILIKE ?", "%"+kw+"%")
	}
	q.Find(&items)
	utils.OK(c, "OK", items)
}

func (h *BrandHandler) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var item models.Brand
	if err := h.db.First(&item, id).Error; err != nil {
		utils.NotFound(c, "Khong tim thay thuong hieu")
		return
	}
	utils.OK(c, "OK", item)
}

func (h *BrandHandler) Create(c *gin.Context) {
	var in brandInput
	if err := c.ShouldBindJSON(&in); err != nil || in.TenThuongHieu == "" {
		utils.BadRequest(c, "Vui long nhap ten thuong hieu")
		return
	}
	item := models.Brand{
		TenThuongHieu: in.TenThuongHieu,
		MoTa:          in.MoTa,
		Logo:          in.Logo,
		QuocGia:       in.QuocGia,
		TrangThai:     true,
	}
	if in.TrangThai != nil {
		item.TrangThai = *in.TrangThai
	}
	if err := h.db.Create(&item).Error; err != nil {
		utils.ServerError(c, "Khong the them thuong hieu")
		return
	}
	utils.Created(c, "Them thuong hieu thanh cong", item)
}

func (h *BrandHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var item models.Brand
	if err := h.db.First(&item, id).Error; err != nil {
		utils.NotFound(c, "Khong tim thay thuong hieu")
		return
	}
	var in brandInput
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}
	if in.TenThuongHieu != "" {
		item.TenThuongHieu = in.TenThuongHieu
	}
	item.MoTa = in.MoTa
	item.Logo = in.Logo
	item.QuocGia = in.QuocGia
	if in.TrangThai != nil {
		item.TrangThai = *in.TrangThai
	}
	h.db.Save(&item)
	utils.OK(c, "Cap nhat thanh cong", item)
}

func (h *BrandHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.db.Delete(&models.Brand{}, id).Error; err != nil {
		utils.ServerError(c, "Xoa that bai")
		return
	}
	utils.OK(c, "Xoa thanh cong", nil)
}
