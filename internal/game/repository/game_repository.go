package repository

import (
	"testgolang/internal/domain"

	"gorm.io/gorm"
)

type GameRepo struct {
	db *gorm.DB
}

func NewGameRepo(db *gorm.DB) *GameRepo {
	return &GameRepo{db}
}

func (gr GameRepo) Insert(game domain.Game) error {

	err := gr.db.Create(&game).Error
	return err
}

func (gr GameRepo) Getdata() ([]domain.Game, error) {
	var game []domain.Game
	err := gr.db.Find(&game).Error
	return game, err

}
