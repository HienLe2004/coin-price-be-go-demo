package main

import (
	"log"

	_ "github.com/HienLe2004/coin-price-be-go-demo/docs"
	"github.com/HienLe2004/coin-price-be-go-demo/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	//Load env
	//Connect db
}

// @title Demo Documenting API
// @version 1
// @Description coin price go backend
// @contact.name Ngoc Hien
// @contact.url https://github.com/HienLe2004/coin-price-be-go-demo

// @host localhost:8080
// @BasePath /api/v1
func main() {

	server := gin.Default()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterRoutes(server)
	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
	//Set up router, routs, server, websocket
}
