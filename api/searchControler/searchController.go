package searchControler

import (
	"github.com/gin-gonic/gin"
	"testproject/internal/services/searchService/contracts"
)

type ISearchController interface {
	Query(ctx *gin.Context)
}

type InvertedIndexController struct {
	service contracts.ISearchService
}

func NewInvertedIndexController(service contracts.ISearchService) *InvertedIndexController {
	return &InvertedIndexController{
		service: service,
	}
}

func (c *InvertedIndexController) Query(ctx *gin.Context) {
	searchTerm := ctx.Query("searchTerm")
	result, ok := c.service.Query(searchTerm)
	if !ok {
		ctx.JSON(400, result)
	}
	ctx.JSON(200, result)
}
