package main

import (
	"api_golang/internal/app"
	"api_golang/internal/auth"
	"api_golang/internal/users"
	"api_golang/pkg/middleware"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db := app.InitDb(app.DB_URL)
	defer app.CloseDb(db)

	r := gin.Default()

	// Register Middlewares
	r.Use(gin.RecoveryWithWriter(os.Stdout))
	r.Use(middleware.Cors())

	// Register Packages
	users.RegisterPackage(db, r)
	auth.RegisterPackage(db, r)

	// Register Static Files
	r.StaticFile("/", "./static/index.html")

	// Register Routes
	r.Run(":" + app.PORT)
}
