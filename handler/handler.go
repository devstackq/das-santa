package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/devstackq/das-santa.git/models"
	"github.com/devstackq/das-santa.git/service"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

const mapID = "faf7ef78-41b3-4a36-8423-688a61929c08"

const basePath = "https://datsanta.dats.team"

const token = `f79375bd-cd2c-40c7-8639-cb8cdb707edc`

type Handler struct {
	srv *service.Service
}

func New(srv *service.Service) *Handler {
	return &Handler{
		srv: srv,
	}
}
func (h Handler) QasqyrRun(c *gin.Context) {
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

	result, err := h.srv.Ebash(data)
	if err != nil {
		c.Status(400)
		return
	}

	result.MapID = mapID

	log.Println(len(result.Moves), "len moves")

	json_data, err := json.Marshal(result)
	if err != nil {
		c.JSON(400, err)
		return
	}

	req, err := http.NewRequest("POST", fmt.Sprint(basePath, "/api/round"), bytes.NewBuffer(json_data))
	if err != nil {
		c.JSON(400, err)
		return
	}

	client := &http.Client{}

	req.Header.Set("X-API-Key", token)

	respRound, err := client.Do(req)
	if err != nil {
		c.JSON(400, err)
		return
	}
	//var responseRound map[string]interface{}

	var responseRound models.ResponseSendRound

	err = json.NewDecoder(respRound.Body).Decode(&responseRound)
	if err != nil {
		c.JSON(400, err)
		return
	}

	if responseRound.Success {

		req, err := http.NewRequest("GET", fmt.Sprint(basePath, "/api/round", responseRound.RoundID), nil)
		if err != nil {
			c.JSON(400, err)
			return
		}
		respRound, err := client.Do(req)
		if err != nil {
			c.JSON(400, err)
			return
		}
		var responseGetRound models.ResponseGetRound

		err = json.NewDecoder(respRound.Body).Decode(&responseGetRound)
		if err != nil {
			c.JSON(400, err)
			return
		}
		c.JSONP(200, gin.H{"result": responseGetRound})
		return
	}

	c.JSONP(400, gin.H{"result": responseRound})
}
