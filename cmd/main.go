package main

import (
	"fmt"
	"testproject/api"
	"testproject/api/authController"
	"testproject/api/searchControler"
	"testproject/api/videoController"
	"testproject/configs"
	"testproject/internal/database"
	"testproject/internal/services/authService/authService"
	"testproject/internal/services/searchService/invertedIndexService"
	"testproject/internal/services/videoService/videoService"
)

func main() {
	fmt.Print("Be name khoda")
	conf, err := configs.LoadDbConfigs("../configs")
	if err != nil {
		panic(err)
	}
	db := database.LoadDb(&conf)
	uow := database.NewUnitOfWork(db)
	searchServiceInstance := invertedIndexService.NewInvertedIndex(uow, 100)
	authServiceInstance := authService.NewAuthService(uow)
	videoServiceInstance := videoService.NewVideoService(uow)
	searchServiceInstance.AddData("Arman Navaran Do Do Eskachi ana mana Kalachi", 1)
	searchServiceInstance.AddData("tap tap Abbasi khoda do do ro nandazi ", 2)
	searchControllerInstance := searchControler.NewInvertedIndexController(searchServiceInstance)
	authControllerInstance := authController.NewAuthController(authServiceInstance)
	videoControllerInstance := videoController.NewVideoController(videoServiceInstance)
	server := api.NewServer(searchControllerInstance, authControllerInstance, videoControllerInstance)
	err = server.Run("127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
}
