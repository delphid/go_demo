package model

import (
	"go.uber.org/zap"
	"gorm.io/gorm"

	config "miata/init"
)

type Store struct {
	log *zap.Logger
	db  *gorm.DB
	cfg *config.Config
}

func NewStore(
	log *zap.Logger,
	db *gorm.DB,
	cfg *config.Config,
) *Store {
	return &Store{
		log: log,
		db:  db,
		cfg: cfg,
	}
}