package contracts

import "testproject/internal/services/videoService/dto"

type IVideoRepository interface {
	GetFavoriteVideos(userId uint) ([]dto.Video, error)
}
