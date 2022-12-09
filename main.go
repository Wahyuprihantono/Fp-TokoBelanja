package main

import (
	"Fp-TokoBelanja/config"
	"Fp-TokoBelanja/controller"
	"Fp-TokoBelanja/middleware"
	"Fp-TokoBelanja/repository"
	"Fp-TokoBelanja/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.StartDB()
	db := config.GetDBConnection()

	userRepository := repository.NewUserRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	productRepository := repository.NewProductRepository(db)

	userService := service.NewUserService(userRepository)
	categoryService := service.NewCategoryService(categoryRepository, productRepository)
	productService := service.NewProductService(productRepository, categoryRepository)

	userController := controller.NewUserController(userService)
	categoryController := controller.NewCategoryController(categoryService)
	productController := controller.NewProductController(productService)

	// Users
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.LoginUser)
	router.PATCH("/users/topup", middleware.AuthMiddleware, userController.PatchTopUpUser)
	// Create Admin
	router.POST("/users/admin", userController.RegisterAdmin)

	// Categories
	router.POST("/categories", middleware.AuthMiddleware, categoryController.CreateCategory)
	router.GET("/categories", middleware.AuthMiddleware, categoryController.GetAllCategories)
	router.PATCH("/categories/:categoryId", middleware.AuthMiddleware, categoryController.PatchCategory)
	router.DELETE("/categories/:categoryId", middleware.AuthMiddleware, categoryController.DeleteCategory)

	// Products
	router.POST("/products", middleware.AuthMiddleware, productController.CreateProduct)
	router.GET("/products", middleware.AuthMiddleware, productController.GetAllProducts)
	router.PUT("/products/:productId", middleware.AuthMiddleware, productController.UpdateProduct)
	router.DELETE("/products/:productId", middleware.AuthMiddleware, productController.DeleteProduct)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
