package videoController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"testproject/api/ResponseHandler"
	"testproject/internal/services/videoService/contracts"
	"testproject/internal/services/videoService/dto"
)

type IVideoController interface {
	All(ctx *gin.Context)
	Creat(ctx *gin.Context)
	GetUserVideos(ctx *gin.Context)
	GetFavoriteVideos(ctx *gin.Context)
}

type VideoController struct {
	service contracts.IVideoService
}

func NewVideoController(service contracts.IVideoService) *VideoController {
	return &VideoController{
		service: service,
	}
}

func (c *VideoController) Creat(ctx *gin.Context) {
	var response ResponseHandler.Response
	var video dto.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		response = ResponseHandler.HandleResponse(err)
	}
	result, err := c.service.Creat(video)
	if err != nil {
		response = ResponseHandler.HandleResponse(err)
	}
	response.StatusCode = http.StatusCreated
	response.Message = "VideoCreated"
	response.PayLoad = result
	ctx.JSON(response.StatusCode, response)
}

func (c *VideoController) GetUserVideos(ctx *gin.Context) {
	var response ResponseHandler.Response
	userID, err := strconv.Atoi(ctx.GetHeader("UserID"))
	if err != nil {
		response = ResponseHandler.HandleResponse(err)
	}
	if userID < 1 {
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response = ResponseHandler.HandleResponse(err)
	}
	result, err := c.service.GetUserVideos(uint(userID))
	if err != nil {
		response = ResponseHandler.HandleResponse(err)
	}
	response.StatusCode = http.StatusOK
	response.Message = ""
	response.PayLoad = result
	ctx.JSON(response.StatusCode, response)
}

func (c *VideoController) GetFavoriteVideos(ctx *gin.Context) {
	var response ResponseHandler.Response
	userID, err := strconv.Atoi(ctx.GetHeader("UserID"))
	if err != nil {
		response = ResponseHandler.HandleResponse(err)
	}
	if userID < 1 {
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	if err != nil {
		response = ResponseHandler.HandleResponse(err)
	}
	result, err := c.service.GetFavoriteVideos(uint(userID))
	if err != nil {
		response = ResponseHandler.HandleResponse(err)
	}
	response.StatusCode = http.StatusOK
	response.Message = ""
	response.PayLoad = result
	ctx.JSON(response.StatusCode, response)
}

func (c *VideoController) All(ctx *gin.Context) {
	var response ResponseHandler.Response
	result, err := c.service.All()
	if err != nil {
		response = ResponseHandler.HandleResponse(err)
	}
	response.StatusCode = http.StatusOK
	response.Message = ""
	response.PayLoad = result
	ctx.JSON(response.StatusCode, response)
}
