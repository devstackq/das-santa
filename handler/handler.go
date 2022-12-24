package handler

import (
	"encoding/json"
	"fmt"
	"github.com/devstackq/das-santa.git/models"
	"github.com/devstackq/das-santa.git/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
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
	data := models.Map{}

	requestURL := fmt.Sprint(basePath, "/json/map/", mapID, ".json")
	resp, err := http.Get(requestURL)

	if err != nil {
		c.Status(400)
		return
	}
	bb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Status(400)
		return
	}
	err = json.Unmarshal(bb, &data)
	if err != nil {
		c.Status(400)
		return
	}

	if err := h.srv.Ebash(data); err != nil {
		c.Status(400)
		return
	}

	c.Status(200)
}

func (h Handler) GetMap(c *gin.Context) {
}

func (h Handler) SendRoute(c *gin.Context) {}

func (h Handler) GetStatus(c *gin.Context) {}
