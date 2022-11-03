package router

import (
	controller "mundhrakeshav/go-http/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(u *controller.UserController,routerGroup *gin.RouterGroup) {
	userRouter := routerGroup.Group("/user")
	userRouter.POST("/create", u.CreateUser)
	userRouter.GET("/get/:name", u.GetUser)
	userRouter.GET("/getall", u.GetAll)
	userRouter.PATCH("/update", u.UpdateUser)
	userRouter.DELETE("/delete", u.DeleteUser)

}