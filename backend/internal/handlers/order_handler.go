package handlers

import (
	"fmt"
	"strconv"
	"time"

	"watch-shop/backend/internal/models"
	"watch-shop/backend/internal/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderHandler struct {
	db *gorm.DB
}

func NewOrderHandler(db *gorm.DB) *OrderHandler {
	return &OrderHandler{db: db}
}

type checkoutInput struct {
	HoTen       string `json:"ho_ten"`
	SoDienThoai string `json:"so_dien_thoai"`
	DiaChi      string `json:"dia_chi"`
	GhiChu      string `json:"ghi_chu"`
	PhuongThuc  string `json:"phuong_thuc"`
}

func (h *OrderHandler) Checkout(c *gin.Context) {
	uid, _ := c.Get("user_id")
	var in checkoutInput
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}
	if in.HoTen == "" || in.SoDienThoai == "" || in.DiaChi == "" {
		utils.BadRequest(c, "Vui long nhap day du thong tin giao hang")
		return
	}
	if !utils.ValidatePhone(in.SoDienThoai) {
		utils.BadRequest(c, "So dien thoai khong hop le")
		return
	}

	var cart []models.CartItem
	h.db.Preload("Product").Where("user_id = ?", uid).Find(&cart)
	if len(cart) == 0 {
		utils.BadRequest(c, "Gio hang dang trong")
		return
	}

	var tong float64
	for _, it := range cart {
		tong += it.Product.Gia * float64(it.SoLuong)
	}

	phuongThuc := in.PhuongThuc
	if phuongThuc == "" {
		phuongThuc = "cod"
	}

	order := models.Order{
		MaDonHang:   fmt.Sprintf("DH%d", time.Now().UnixMilli()),
		UserID:      uid.(uint),
		HoTen:       in.HoTen,
		SoDienThoai: in.SoDienThoai,
		DiaChi:      in.DiaChi,
		GhiChu:      in.GhiChu,
		PhuongThuc:  phuongThuc,
		TrangThai:   models.OrderStatusPending,
		TongTien:    tong,
		NgayDatHang: time.Now(),
	}

	err := h.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&order).Error; err != nil {
			return err
		}
		for _, it := range cart {
			oi := models.OrderItem{
				OrderID:    order.ID,
				ProductID:  it.ProductID,
				TenSanPham: it.Product.TenSanPham,
				HinhAnh:    it.Product.HinhAnh,
				Gia:        it.Product.Gia,
				SoLuong:    it.SoLuong,
				ThanhTien:  it.Product.Gia * float64(it.SoLuong),
			}
			if err := tx.Create(&oi).Error; err != nil {
				return err
			}
			tx.Model(&models.Product{}).Where("id = ?", it.ProductID).
				Updates(map[string]interface{}{
					"so_luong": gorm.Expr("so_luong - ?", it.SoLuong),
					"da_ban":   gorm.Expr("da_ban + ?", it.SoLuong),
				})
		}
		return tx.Where("user_id = ?", uid).Delete(&models.CartItem{}).Error
	})

	if err != nil {
		utils.ServerError(c, "Khong the dat hang: "+err.Error())
		return
	}
	utils.Created(c, "Dat hang thanh cong", order)
}

func (h *OrderHandler) MyOrders(c *gin.Context) {
	uid, _ := c.Get("user_id")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}

	q := h.db.Model(&models.Order{}).Where("user_id = ?", uid)
	if st := c.Query("trang_thai"); st != "" {
		q = q.Where("trang_thai = ?", st)
	}

	var total int64
	q.Count(&total)

	var items []models.Order
	q.Preload("ChiTietDonHang").Order("id DESC").Limit(size).Offset((page - 1) * size).Find(&items)
	utils.Page(c, items, total, page, size)
}

func (h *OrderHandler) MyOrderDetail(c *gin.Context) {
	uid, _ := c.Get("user_id")
	id, _ := strconv.Atoi(c.Param("id"))
	var order models.Order
	if err := h.db.Preload("ChiTietDonHang").Where("id = ? AND user_id = ?", id, uid).First(&order).Error; err != nil {
		utils.NotFound(c, "Khong tim thay don hang")
		return
	}
	utils.OK(c, "OK", order)
}

func (h *OrderHandler) Cancel(c *gin.Context) {
	uid, _ := c.Get("user_id")
	id, _ := strconv.Atoi(c.Param("id"))
	var order models.Order
	if err := h.db.Where("id = ? AND user_id = ?", id, uid).First(&order).Error; err != nil {
		utils.NotFound(c, "Khong tim thay don hang")
		return
	}
	if order.TrangThai != models.OrderStatusPending {
		utils.BadRequest(c, "Don hang khong the huy")
		return
	}
	order.TrangThai = models.OrderStatusCancelled
	h.db.Save(&order)

	var items []models.OrderItem
	h.db.Where("order_id = ?", order.ID).Find(&items)
	for _, it := range items {
		h.db.Model(&models.Product{}).Where("id = ?", it.ProductID).Updates(map[string]interface{}{
			"so_luong": gorm.Expr("so_luong + ?", it.SoLuong),
			"da_ban":   gorm.Expr("da_ban - ?", it.SoLuong),
		})
	}
	utils.OK(c, "Da huy don hang", order)
}

func (h *OrderHandler) AdminList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}

	q := h.db.Model(&models.Order{})
	if st := c.Query("trang_thai"); st != "" {
		q = q.Where("trang_thai = ?", st)
	}
	if kw := c.Query("keyword"); kw != "" {
		q = q.Where("ma_don_hang ILIKE ? OR ho_ten ILIKE ? OR so_dien_thoai ILIKE ?",
			"%"+kw+"%", "%"+kw+"%", "%"+kw+"%")
	}

	var total int64
	q.Count(&total)

	var items []models.Order
	q.Preload("User").Order("id DESC").Limit(size).Offset((page - 1) * size).Find(&items)
	utils.Page(c, items, total, page, size)
}

func (h *OrderHandler) AdminDetail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var order models.Order
	if err := h.db.Preload("ChiTietDonHang").Preload("User").First(&order, id).Error; err != nil {
		utils.NotFound(c, "Khong tim thay don hang")
		return
	}
	utils.OK(c, "OK", order)
}

type updateStatusInput struct {
	TrangThai string `json:"trang_thai"`
}

func (h *OrderHandler) UpdateStatus(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var in updateStatusInput
	if err := c.ShouldBindJSON(&in); err != nil {
		utils.BadRequest(c, "Du lieu khong hop le")
		return
	}
	valid := map[string]bool{
		models.OrderStatusPending:   true,
		models.OrderStatusConfirmed: true,
		models.OrderStatusShipping:  true,
		models.OrderStatusDone:      true,
		models.OrderStatusCancelled: true,
	}
	if !valid[in.TrangThai] {
		utils.BadRequest(c, "Trang thai khong hop le")
		return
	}
	if err := h.db.Model(&models.Order{}).Where("id = ?", id).Update("trang_thai", in.TrangThai).Error; err != nil {
		utils.ServerError(c, "Cap nhat that bai")
		return
	}
	utils.OK(c, "Cap nhat thanh cong", nil)
}

func (h *OrderHandler) Statistics(c *gin.Context) {
	var totalOrder, totalUser, totalProduct int64
	h.db.Model(&models.Order{}).Count(&totalOrder)
	h.db.Model(&models.User{}).Count(&totalUser)
	h.db.Model(&models.Product{}).Count(&totalProduct)

	var doanhThu float64
	h.db.Model(&models.Order{}).
		Where("trang_thai = ?", models.OrderStatusDone).
		Select("COALESCE(SUM(tong_tien), 0)").Scan(&doanhThu)

	type statusCount struct {
		TrangThai string `json:"trang_thai"`
		SoLuong   int64  `json:"so_luong"`
	}
	var theoTrangThai []statusCount
	h.db.Model(&models.Order{}).Select("trang_thai, COUNT(*) AS so_luong").
		Group("trang_thai").Scan(&theoTrangThai)

	utils.OK(c, "OK", gin.H{
		"tong_don_hang": totalOrder,
		"tong_nguoi_dung": totalUser,
		"tong_san_pham":   totalProduct,
		"doanh_thu":       doanhThu,
		"theo_trang_thai": theoTrangThai,
	})
}
