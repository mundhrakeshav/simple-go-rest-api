package controllers

import (
	"fmt"
	"mundhrakeshav/go-http/pkg/models"
	"mundhrakeshav/go-http/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()});
		return;
	}
	err := u.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()});
		return;
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"user": user,
	})
}

func (u *UserController) GetUser(ctx *gin.Context) {
	var username string = ctx.Param("name")
	user, err := u.UserService.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, user)
	}}

func (u *UserController) GetAll(ctx *gin.Context) {
	users, err := u.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, users)
	}
}


func (u *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	updatedUser, err := u.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	byteData, err := ctx.GetRawData()
	fmt.Println(byteData, err)
	ctx.JSON(http.StatusOK, "Mil")
}
