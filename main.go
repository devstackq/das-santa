package main

import (
	"github.com/devstackq/das-santa.git/handler"
	"github.com/devstackq/das-santa.git/service"
	"github.com/gin-gonic/gin"
)

func main() {

	srv := service.New()
	handler := handler.New(srv)
	router := gin.Default()
	router.POST("/qasqyr", handler.Qasqyr)
	router.Run("localhost:8080")

}
