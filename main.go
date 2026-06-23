package main

import (
	"context"
	"fmt"
	"os"

	"github.com/NaufalParamaRafif/CICO-Learn-Backend/initializers"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/internal/handler"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/internal/repository"
	"github.com/NaufalParamaRafif/CICO-Learn-Backend/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func init() {
	initializers.LoadEnvVariables()
	// initializers.ConnectToDB()
	// initializers.SyncDatabase()
}

func main() {
	route := gin.Default()

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println(err.Error())
		panic(err.Error())
	}

	// Repository
	userRepository := repository.NewUserRepository(pool)

	// Usecase
	userUseCase := usecase.NewUserUseCase(&userRepository)

	// Handler
	handler.NewAuthHandler(route, userUseCase)
	handler.NewUserHandler(route, userUseCase)

	route.Run()
}
