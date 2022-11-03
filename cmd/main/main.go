package main

import (
	"context"
	"mundhrakeshav/go-http/pkg/controllers"
	"mundhrakeshav/go-http/pkg/db"
	"mundhrakeshav/go-http/pkg/router"
	"mundhrakeshav/go-http/pkg/services"
	"os"

	logger "mundhrakeshav/go-http/pkg/log"

	gin "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	server          *gin.Engine
	user_service    services.UserService
	user_controller controllers.UserController
	ctx             context.Context)

func init() {
	gin.SetMode(gin.ReleaseMode)
	//Load .env
	godotenv.Load(".env")
	user_service = services.NewUserService(db.DBStore.GetUsersCollection(), ctx)
	user_controller = controllers.NewUserController(user_service)
	server = gin.Default()
}

func main() {
	defer db.DBStore.Disconnect()
	defer logger.Log.Sync() // flushes buffer, if any
	router_group := server.Group("/v1")
	router.RegisterUserRoutes(&user_controller, router_group)
	PORT := os.Getenv("PORT")
	logger.Log.Info("Listening on port" + PORT)
	if err := server.Run(PORT); err != nil {
		logger.Log.Fatal(err.Error())
	}
}


// npx nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/main/main.go 