package handler

import (
	"testgolang/internal/domain"

	"github.com/gin-gonic/gin"
)

type GameService interface {
	InsertService(game domain.Game) (int, error)
	GetdataService() ([]domain.Game, int, error)
}

type GameHandler struct {
	GameHan GameService
}

func NewGameHandler(Gameserv GameService) *GameHandler {
	return &GameHandler{Gameserv}
}

func (gh GameHandler) InsertGame(ctx *gin.Context) {
	var game domain.Game
	err := ctx.ShouldBindJSON(&game)
	if err != nil {
		ctx.JSON(404, gin.H{
			"code":    "404",
			"message": "error format json",
		})
		return
	}

	code, err := gh.GameHan.InsertService(game)

	if err != nil {
		ctx.JSON(code, gin.H{
			"code":    code,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success insert data",
	})

}

func (gh GameHandler) Getdata(ctx *gin.Context) {

	data, code, err := gh.GameHan.GetdataService()

	if err != nil {
		ctx.JSON(code, gin.H{
			"code":    code,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(code, gin.H{
		"code":    code,
		"data":    data,
		"message": "success get data",
	})
}
