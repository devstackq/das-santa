package handler

import (
	"bytes"
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


	json_data, err := json.Marshal(result)
	if err != nil {
		log.Println(err, 1)
		c.Status(400)
		return
	}
	req, err := http.NewRequest("POST", fmt.Sprint(basePath, "/api/round"), bytes.NewBuffer(json_data))
	if err != nil {
		log.Println(err, 1)
		c.Status(400)
		return
	}

	client := &http.Client{}
	//req.Host = basePath

	req.Header.Set("Authorization", token)

	respRound, err := client.Do(req)
	if err != nil {
		log.Println(err, 1)
		c.Status(400)
		return
	}
	var rez map[string]interface{}

	err = json.NewDecoder(respRound.Body).Decode(&rez)
	if err != nil {
		log.Println(err, 1)

		c.Status(400)
		return
	}
	fmt.Println(rez, "response route")
	//TODO: send http.Post(url, result)

	c.JSON(200, result)
}
