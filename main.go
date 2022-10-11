package main

import (
	"testgolang/internal/database"
	"testgolang/internal/game/handler"
	"testgolang/internal/game/repository"
	"testgolang/internal/game/service"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db := database.ConnectDatabase()
	GameRepo := repository.NewGameRepo(db)
	GameService := service.NewGameService(GameRepo)
	Gamehand := handler.NewGameHandler(GameService)

	router.POST("/create", Gamehand.InsertGame)
	router.GET("/getdata", Gamehand.Getdata)

	router.Run()
}
