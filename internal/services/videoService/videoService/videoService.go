package videoService

import (
	"testproject/internal/database"
	"testproject/internal/services/videoService/dto"
)

type VideoService struct {
	uow *database.UnitOfWork
}

func NewVideoService(uow *database.UnitOfWork) *VideoService {
	return &VideoService{
		uow: uow,
	}
}

func (v VideoService) All() ([]dto.Video, error) {
	var videos []dto.Video
	err := v.uow.BaseRepository.ReadAll(&videos)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (v VideoService) Creat(video dto.Video) (*dto.Video, error) {
	err := v.uow.BaseRepository.Create(&video)
	if err != nil {
		return nil, err
	}
	return &video, nil
}

func (v VideoService) GetUserVideos(userId uint) ([]dto.Video, error) {
	var videos []dto.Video
	err := v.uow.BaseRepository.ReadAllBy(&videos, "UserID", userId)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (v VideoService) GetFavoriteVideos(userId uint) ([]dto.Video, error) {
	return nil, nil
}
