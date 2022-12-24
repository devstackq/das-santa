package handler

import (
	"github.com/devstackq/das-santa.git/service"
	"github.com/gin-gonic/gin"
)

const mapID = "faf7ef78-41b3-4a36-8423-688a61929c08"

const basePath = "https://datsanta.dats.team"

type Handler struct {
	srv *service.Service
}

func New(srv *service.Service) *Handler {
	return &Handler{
		srv: srv,
	}
}
func (h Handler) Qasqyr(c *gin.Context) {
	/*
		1 getMap
		2 run srv.Ebash(request.Map)
		3
	*/

}

func (h Handler) GetMap(c *gin.Context) {
}

func (h Handler) SendRoute(c *gin.Context) {}

func (h Handler) GetStatus(c *gin.Context) {}
