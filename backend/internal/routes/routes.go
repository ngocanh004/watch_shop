package routes

import (
	"watch-shop/backend/internal/config"
	"watch-shop/backend/internal/handlers"
	"watch-shop/backend/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, cfg *config.Config) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	r.Static("/uploads", cfg.UploadDir)

	auth := handlers.NewAuthHandler(db, cfg)
	user := handlers.NewUserHandler(db)
	cat := handlers.NewCategoryHandler(db)
	brand := handlers.NewBrandHandler(db)
	product := handlers.NewProductHandler(db)
	cart := handlers.NewCartHandler(db)
	wish := handlers.NewWishlistHandler(db)
	order := handlers.NewOrderHandler(db)
	upload := handlers.NewUploadHandler(cfg)

	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api.POST("/auth/dang-ky", auth.Register)
	api.POST("/auth/dang-nhap", auth.Login)

	api.GET("/categories", cat.List)
	api.GET("/categories/:id", cat.Get)
	api.GET("/brands", brand.List)
	api.GET("/brands/:id", brand.Get)
	api.GET("/products", product.List)
	api.GET("/products/:id", product.Get)

	authReq := api.Group("/")
	authReq.Use(middleware.JWTAuth(cfg))
	{
		authReq.GET("/me", auth.Profile)
		authReq.PUT("/me", auth.UpdateProfile)
		authReq.POST("/auth/doi-mat-khau", auth.ChangePassword)

		authReq.GET("/cart", cart.List)
		authReq.POST("/cart", cart.Add)
		authReq.PUT("/cart/:id", cart.Update)
		authReq.DELETE("/cart/:id", cart.Remove)
		authReq.DELETE("/cart", cart.Clear)

		authReq.GET("/wishlist", wish.List)
		authReq.POST("/wishlist", wish.Add)
		authReq.DELETE("/wishlist/:product_id", wish.Remove)

		authReq.POST("/orders", order.Checkout)
		authReq.GET("/orders/me", order.MyOrders)
		authReq.GET("/orders/me/:id", order.MyOrderDetail)
		authReq.PUT("/orders/me/:id/huy", order.Cancel)
	}

	admin := api.Group("/admin")
	admin.Use(middleware.JWTAuth(cfg), middleware.RequireAdmin())
	{
		admin.GET("/thong-ke", order.Statistics)

		admin.GET("/users", user.List)
		admin.GET("/users/:id", user.Get)
		admin.PUT("/users/:id", user.Update)
		admin.DELETE("/users/:id", user.Delete)
		admin.POST("/users/:id/reset-password", user.ResetPassword)

		admin.POST("/categories", cat.Create)
		admin.PUT("/categories/:id", cat.Update)
		admin.DELETE("/categories/:id", cat.Delete)

		admin.POST("/brands", brand.Create)
		admin.PUT("/brands/:id", brand.Update)
		admin.DELETE("/brands/:id", brand.Delete)

		admin.POST("/products", product.Create)
		admin.PUT("/products/:id", product.Update)
		admin.DELETE("/products/:id", product.Delete)

		admin.GET("/orders", order.AdminList)
		admin.GET("/orders/:id", order.AdminDetail)
		admin.PUT("/orders/:id/trang-thai", order.UpdateStatus)

		admin.POST("/upload", upload.Image)
	}

	return r
}
