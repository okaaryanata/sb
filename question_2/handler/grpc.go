package handler

import (
	"context"
	"encoding/json"
	"question_2/db"
	"question_2/model"
	"question_2/pb"
	"strconv"
)

// GrpcServer - struct for grpc service
type GrpcServer struct {
	Db db.Connection
}

// GetMovie - get movie from api
func (s *GrpcServer) GetMovie(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	page, word := request.GetPagination(), request.GetSearchWord()
	pbRes := pb.Response{}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pbRes.ErrorMessage = err.Error()
		pbRes.IsSuccess = false
	}
	logHistory := model.LogHistory{}
	logHistory.Pagination = int64(pageInt)
	logHistory.SearchWord = word

	res, err := GetDataMovie(word, page)
	if err != nil {
		logHistory.Success = false
		logHistory.Insert(s.Db)

		pbRes.ErrorMessage = err.Error()
		pbRes.IsSuccess = false
		return &pbRes, err
	}
	logHistory.Success = true
	logHistory.Insert(s.Db)

	err = json.Unmarshal(res, &pbRes)
	if err != nil {
		pbRes.IsSuccess = false
		pbRes.ErrorMessage = err.Error()
		return &pbRes, err
	}
	pbRes.IsSuccess = true
	return &pbRes, nil
}
