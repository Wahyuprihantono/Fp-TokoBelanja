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
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryController := controller.NewCategoryController(categoryService)

	// Users
	router.POST("/users/register", userController.RegisterUser)
	router.POST("/users/login", userController.LoginUser)
	router.PATCH("/users/topup", middleware.AuthMiddleware, userController.PatchTopUpUser)
	// Create Admin
	router.POST("/users/admin", userController.RegisterAdmin)

	// Categories
	router.POST("/categories", middleware.AuthMiddleware, categoryController.CreateCategory)
	router.GET("/categories", middleware.AuthMiddleware, categoryController.GetAllCategories)
	router.PATCH("/categories/:id", middleware.AuthMiddleware, categoryController.PatchCategory)
	router.DELETE("/categories/:id", middleware.AuthMiddleware, categoryController.DeleteCategory)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
