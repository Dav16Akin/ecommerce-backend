package main

import (
	"log"

	"github.com/Dav16Akin/ecommerce-rest-backend/internal/database"
	handler "github.com/Dav16Akin/ecommerce-rest-backend/internal/handler"
	model "github.com/Dav16Akin/ecommerce-rest-backend/internal/models"
	"github.com/Dav16Akin/ecommerce-rest-backend/internal/repository"
	"github.com/Dav16Akin/ecommerce-rest-backend/internal/service"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

type Application struct {
	UserHandler handler.UserHandler
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}


	db, err := database.ConnectToDB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&model.User{},
	)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := &Application{
		UserHandler: userHandler,
	}

	router := gin.Default()

	router.GET("/v1/user", app.UserHandler.GetUser)

	router.Run(":8000")

}
