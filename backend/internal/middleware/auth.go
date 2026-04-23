package middleware

import (
	"strings"

	"watch-shop/backend/internal/config"
	"watch-shop/backend/internal/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuth(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			utils.Unauthorized(c, "Thieu token xac thuc")
			c.Abort()
			return
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Unauthorized(c, "Token khong dung dinh dang")
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(parts[1], cfg.JWTSecret)
		if err != nil {
			utils.Unauthorized(c, "Token het han hoac khong hop le")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("vai_tro", claims.VaiTro)
		c.Next()
	}
}

func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		vaiTro, _ := c.Get("vai_tro")
		if vaiTro != "admin" {
			utils.Forbidden(c, "Ban khong co quyen thuc hien")
			c.Abort()
			return
		}
		c.Next()
	}
}
