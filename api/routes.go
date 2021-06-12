package api

import "github.com/gin-gonic/gin"

func InitRoutes() {

	blockchainApi := NewBlockchainApi()

	router := gin.Default()

	router.GET("mine-block", blockchainApi.MineBlock)

	router.GET("chain", blockchainApi.GetChain)

	router.GET("block/:index", blockchainApi.GetBlock)

	router.GET("is-valid", blockchainApi.IsValid)

	router.Run()

}
