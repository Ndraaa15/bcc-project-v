package handlers

import (
	"bcc-project-v/sdk/config"
	"bcc-project-v/src/helper"
	"bcc-project-v/src/middleware"
	"bcc-project-v/src/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	http       *gin.Engine
	config     config.Initializer
	db         *gorm.DB
	Repository repository.Repository
}

func Init(config config.Initializer, repo *repository.Repository) *handler {
	rest := handler{
		http:       gin.Default(),
		config:     config,
		Repository: *repo,
	}
	rest.registerRoutes()
	return &rest
}

func (h *handler) registerRoutes() {
	repository.NewRepository(h.db)
	h.Repository.SeedCategory()
	h.Repository.SeedStatus()

	// userHandler := newUserHandler(h.Repository, h.config)
	// userHandler.UserRegister

	h.http.GET("/", func(ctx *gin.Context) {
		helper.SuccessResponse(ctx, http.StatusOK, "Hello World", nil)
	})

	v1 := h.http.Group("/api/v1")

	//User
	user := h.http.Group(v1.BasePath() + "/user")
	user.POST("/signup", h.UserRegister)
	user.POST("/login", h.UserLogin)
	user.Use(middleware.NewRepository(h.db).IsUserLoggedIn()).
		PUT("/update", h.UserUpdate)
		// PUT("/update/password", h.UserUpdatePassword)

	//Admin
	seller := h.http.Group(v1.BasePath() + "/seller")
	seller.POST("/signup", h.SellerRegister)
	seller.GET("/login", h.SellerLogin)
	seller.Use(middleware.NewRepository(h.db).IsSellerLoggedIn()).
		PUT("/update", h.SellerUpdate)
		// PUT("/update/password", h.SellerUpdatePassword)

	//Product for seller
	product_seller := h.http.Group(v1.BasePath() + "/seller/market")
	product_seller.
		Use(middleware.NewRepository(h.db).IsSellerLoggedIn()).
		POST("/product/upload", h.PostProduct).
		PUT("/product/:product_id", h.UpdateProduct).
		GET("/product", h.GetSellerProduct).
		GET("/product/:product_id", h.GetSellerProductByID).
		DELETE("/product/:product_id", h.DeleteProductByID).
		POST("check", h.CheckPayment)

	//Product for user
	product_user := h.http.Group(v1.BasePath() + "/user/market")
	product_user.
		Use(middleware.NewRepository(h.db).IsUserLoggedIn()).
		GET("/products", h.GetAllProduct).
		GET("/product/:product_id", h.GetProductByID).
		POST("/cart", h.AddProductToCart).
		DELETE("/cart", h.RemoveProductFromCart).
		GET("/cart", h.GetProductCart).
		GET("/products/filter/:category", h.GetProductByFilter).
		GET("/search", h.SearchProduct).
		POST("/payment/offline", h.OfflinePayment).
		POST("/payment/online", h.OnlinePayment).
		GET("/history", h.GetHistory)
}

func (h *handler) Run() {
	h.http.Run(fmt.Sprintf(":%s", h.config.GetEnv("PORT")))
}
