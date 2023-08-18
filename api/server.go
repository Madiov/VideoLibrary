package api

import (
	"github.com/gin-gonic/gin"
	"testproject/api/authController"
	"testproject/api/searchControler"
	"testproject/api/videoController"
)

type Server struct {
	gin              *gin.Engine
	searchController searchControler.ISearchController
	authController   authController.IAuthController
	videoController  videoController.IVideoController
}

func NewServer(searchController searchControler.ISearchController,
	authController authController.IAuthController,
	videoController videoController.IVideoController) *Server {
	server := Server{
		gin:              gin.Default(),
		searchController: searchController,
		authController:   authController,
		videoController:  videoController,
	}
	server.SetRoutes()
	return &server
}

func (s *Server) SetRoutes() {
	//Auth Routes
	authGroup := s.gin.Group("/auth")
	authGroup.POST("/login", s.authController.Login)
	authGroup.POST("/", s.authController.Register)

	//Api Authorization
	apiGroup := s.gin.Group("/api")
	apiGroup.Use(s.authController.MiddleAuth)

	//Search Routes
	searchGroup := apiGroup.Group("/search")
	searchGroup.GET("/", s.searchController.Query)

	//video Routes
	videoGroup := apiGroup.Group("/video")
	videoGroup.GET("/", s.videoController.All)
}

func (s *Server) Run(address string) error {
	return s.gin.Run(address)
}
