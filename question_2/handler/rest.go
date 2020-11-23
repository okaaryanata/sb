package handler

import (
	"encoding/json"
	"net/http"
	"question_2/db"
	"question_2/model"
	"strconv"
)

// RestServer - struct for rest transport
type RestServer struct {
	Db db.Connection
}

type resStruct struct {
	Search string `json:"search"`
	Page   string `json:"page"`
}

type errorResponse struct {
	Message string `json:"message"`
}

// NewRestServer - init rest
func NewRestServer(db db.Connection) RestServer {
	return RestServer{Db: db}
}

// GetData - get movie data
func (h RestServer) GetData(res http.ResponseWriter, req *http.Request) {
	words, ok := req.URL.Query()["search"]

	if !ok || len(words[0]) < 1 {
		js := h.createErrRes("Url Param 'search' is missing")
		res.WriteHeader(http.StatusNotAcceptable)
		res.Write(js)
		return
	}

	paginations, ok := req.URL.Query()["pagination"]

	if !ok || len(paginations[0]) < 1 {
		js := h.createErrRes("Url Param 'pagination' is missing")
		res.WriteHeader(http.StatusNotAcceptable)
		res.Write(js)
		return
	}
	pageStr := paginations[0]
	pageInt, err := strconv.Atoi(pageStr)
	if err != nil {
		js := h.createErrRes(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(js)
		return
	}

	word := words[0]

	resAPI, err := GetDataMovie(word, pageStr)
	if err != nil {
		log := model.LogHistory{SearchWord: word, Pagination: int64(pageInt), Success: false}
		_, err1 := log.Insert(h.Db)
		if err1 != nil {
			js := h.createErrRes(err1.Error())
			res.WriteHeader(http.StatusInternalServerError)
			res.Write(js)
			return
		}
		js := h.createErrRes(err.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(js)
		return
	}

	log := model.LogHistory{SearchWord: word, Pagination: int64(pageInt), Success: true}
	_, err1 := log.Insert(h.Db)
	if err1 != nil {
		js := h.createErrRes(err1.Error())
		res.WriteHeader(http.StatusInternalServerError)
		res.Write(js)
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resAPI)
}

func (h RestServer) createErrRes(message string) []byte {
	err := errorResponse{Message: message}
	res, _ := json.Marshal(err)
	return res
}
