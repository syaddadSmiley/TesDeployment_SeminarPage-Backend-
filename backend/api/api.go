package api

import (
	"fmt"

	repository "github.com/syaddadSmiley/SeminarPage/repository"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type API struct {
	userRepo  repository.UserRepo
	adminRepo repository.AdminRepo
	gin       *gin.Engine
}

func NewAPI(userRepo repository.UserRepo, adminRepo repository.AdminRepo) *API {
	gin := gin.Default()
	api := &API{
		userRepo:  userRepo,
		adminRepo: adminRepo,
		gin:       gin,
	}

	//User//
	gin.Any("/Login", api.Login)
	gin.Any("/Register", api.Register)
	gin.POST("/Logout", api.AuthMiddleWare(api.Logout))
	gin.GET("/MyProfile", api.AuthMiddleWare(api.GetProfile))
	gin.DELETE("/DeleteUser/:id", api.AuthMiddleWare(api.DeleteUserByID))
	gin.PUT("/UpdateUser/:id", api.AuthMiddleWare(api.UpdateUserByID)) //ADA BUG
	gin.POST("/UpdatePassword", api.AuthMiddleWare(api.UpdatePassword))

	gin.GET("/GetBarang", api.AuthMiddleWare(api.SearchProductUser))
	gin.GET("/GetBarang/sort", api.AuthMiddleWare(api.SortingProduct))
	gin.GET("/GetBarang/filter", api.AuthMiddleWare(api.FilterByGame))
	gin.GET("/api/GetBarang/:id", api.AuthMiddleWare(api.GetProductByID))

	gin.POST("/Komentar", api.AuthMiddleWare(api.CreateKomentar))
	gin.DELETE("/Komentar/hapus", api.AuthMiddleWare(api.DeleteKomentar))

	gin.POST("/Wishlist/add", api.AuthMiddleWare(api.CreateWishlist))
	gin.DELETE("/Wishlist/hapus", api.AuthMiddleWare(api.DeleteWishlist))
	gin.GET("/Wishlist/get", api.AuthMiddleWare(api.GetWishlist))
	gin.GET("/Wishlist/all/:id", api.AuthMiddleWare(api.GetAllWishlist))

	gin.POST("/Coupons/use", api.AuthMiddleWare(api.UseCoupon))

	gin.POST("/Basket", api.AuthMiddleWare(api.Basket))
	gin.POST("GetBasket", api.AuthMiddleWare(api.CheckBasket))

	gin.GET("/Pagination", api.AuthMiddleWare(api.Pagination))
	gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowCredentials: true,
	}))
	//admin//
	gin.Any("/RegisterAdmin", api.RegisterAdmin)

	gin.POST("/AddCategory", api.AuthMiddleWare(api.AdminMiddleware(api.CreateCategory)))
	gin.PUT("/EditCategory", api.AuthMiddleWare(api.AdminMiddleware(api.UpdateCategory)))
	gin.DELETE("/DeleteCategory", api.AuthMiddleWare(api.AdminMiddleware(api.DeleteCategory)))
	gin.GET("/GetCategory", api.AuthMiddleWare(api.AdminMiddleware(api.GetCategory)))

	gin.POST("/AddProduct", api.AuthMiddleWare(api.AdminMiddleware(api.CreateProduct)))
	gin.PUT("/EditProduct", api.AuthMiddleWare(api.AdminMiddleware(api.UpdateProduct)))
	gin.DELETE("/DeleteProduct", api.AuthMiddleWare(api.AdminMiddleware(api.DeleteProduct)))

	gin.POST("/Coupons/add", api.AuthMiddleWare(api.AdminMiddleware(api.AddCoupon)))

	gin.POST("Notification", api.AuthMiddleWare(api.Notification))

	//buat testing aja with ginkgo :)
	gin.DELETE("/DeleteUser", api.DeleteUser)
	return api
}

func (api *API) Handler() *gin.Engine {
	return api.gin
}

func (api *API) Start() {
	fmt.Println("http://localhst:8008/")
	api.Handler().Run(":8008")
}
