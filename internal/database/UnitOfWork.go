package database

import (
	"gorm.io/gorm"
	"testproject/internal/database/contracts"
	"testproject/internal/database/repository"
)

type UnitOfWork struct {
	db              *gorm.DB
	BaseRepository  contracts.IBaseRepository
	UserRepository  contracts.IUserRepository
	VideoRepository contracts.IVideoRepository
}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {
	baseRepository := repository.NewBaseRepository(db)
	userRepository := repository.NewUserRepository(db)
	videoRepository := repository.NewVideoRepository(db)
	return &UnitOfWork{
		BaseRepository:  baseRepository,
		UserRepository:  userRepository,
		VideoRepository: videoRepository,
	}
}
