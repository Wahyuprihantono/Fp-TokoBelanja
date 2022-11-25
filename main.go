package main

import (
	"os"

	_handler "Fp-TokoBelanja/app/delivery"
	_repository "Fp-TokoBelanja/app/repository"
	_usecase "Fp-TokoBelanja/app/usecase"
	"Fp-TokoBelanja/config"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.StartDB()
	db := config.GetDBConnection()

	userRepository := _repository.NewUserRepository(db)
	userUsecase := _usecase.NewUserUsecase(userRepository)

	categoryRepository := _repository.NewCategoryRepository(db)
	categoryUsecase := _usecase.NewCategoryUsecase(categoryRepository)

	api := router.Group("/")
	_handler.NewUserHandler(api, userUsecase)
	_handler.NewCategoryHandler(api, categoryUsecase)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
