package middleware

import (
	"Fp-TokoBelanja/database"
	"Fp-TokoBelanja/helper"
	"Fp-TokoBelanja/model"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helper.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthenticated",
				"message": err.Error(),
			})
			return
		}
		db := database.GetDB()
		userData := verifyToken.(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := model.User{}
		err = db.Select("id").First(&User, uint(userID)).Error

		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Unauthenticated",
				"message": "something wrong",
			})
			return
		}
		c.Set("userData", verifyToken)
		fmt.Println("Authenticated...")

		c.Next()
	}
}
