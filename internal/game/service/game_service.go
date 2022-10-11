package service

import (
	"errors"
	"testgolang/internal/domain"
)

type GameRepo interface {
	Insert(game domain.Game) error
	Getdata() ([]domain.Game, error)
}

type GameService struct {
	GameSer GameRepo
}

func NewGameService(Gamerepo GameRepo) *GameService {
	return &GameService{Gamerepo}
}

func (gs GameService) InsertService(game domain.Game) (int, error) {

	err := gs.GameSer.Insert(game)

	if err != nil {
		return 500, errors.New("errors insert database")
	}

	return 200, nil

}

func (gs GameService) GetdataService() ([]domain.Game, int, error) {
	data, err := gs.GameSer.Getdata()

	if err != nil {
		return data, 500, errors.New("errors insert database")
	}

	return data, 200, nil

}
