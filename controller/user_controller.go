package controller

import (
	"Fp-TokoBelanja/helper"
	"Fp-TokoBelanja/model/input"
	"Fp-TokoBelanja/model/response"
	"Fp-TokoBelanja/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (h *userController) RegisterUser(c *gin.Context) {
	var input input.UserRegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	userData, err := h.userService.RegisterUser(input)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	userResponse := response.UserRegisterResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		Password:  userData.Password,
		Balance:   userData.Balance,
		CreatedAt: userData.CreatedAt,
	}

	c.JSON(
		http.StatusCreated,
		helper.NewResponse(
			http.StatusCreated,
			"created",
			userResponse,
		),
	)
}

func (h *userController) LoginUser(c *gin.Context) {
	var input input.UserLoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	token, err := h.userService.LoginUser(input)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	userResponse := response.UserLoginResponse{
		Token: token,
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			userResponse,
		),
	)
}

func (h *userController) PatchTopUpUser(c *gin.Context) {
	var input input.UserPatchTopUpInput

	id_user := c.MustGet("currentUser").(int)

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	userData, err := h.userService.TopUpUser(id_user, input)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	message := "Your balance has been succcessfully updated to Rp " + strconv.Itoa(userData.Balance)
	userResponse := response.UserPatchTopUpResponse{
		Message: message,
	}

	c.JSON(
		http.StatusOK,
		helper.NewResponse(
			http.StatusOK,
			"ok",
			userResponse,
		),
	)
}

func (h *userController) RegisterAdmin(c *gin.Context) {
	var input input.UserRegisterInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	userData, err := h.userService.RegisterAdmin(input)
	if err != nil {
		errors := helper.GetErrorData(err)
		c.JSON(
			http.StatusUnprocessableEntity,
			helper.NewErrorResponse(
				http.StatusUnprocessableEntity,
				"failed",
				errors,
			),
		)
		return
	}

	userResponse := response.UserRegisterResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		Password:  userData.Password,
		Balance:   userData.Balance,
		CreatedAt: userData.CreatedAt,
	}

	c.JSON(
		http.StatusCreated,
		helper.NewResponse(
			http.StatusCreated,
			"created",
			userResponse,
		),
	)
}
