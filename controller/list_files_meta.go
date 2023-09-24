package controller

import (
	"go.uber.org/zap"
	"net/http"
	"rest-api-file-server/model"
	"rest-api-file-server/service"
)

type ListFilesMeta struct {
	logger  *zap.Logger
	service service.FileMetaService
}

func NewListFiles(logger *zap.Logger, service service.FileMetaService) *ListFilesMeta {
	return &ListFilesMeta{logger: logger, service: service}
}

func (ctr *ListFilesMeta) ListFiles(w http.ResponseWriter, r *http.Request) {
	files, err := ctr.service.ListFilesMeta()
	if err != nil {
		http.Error(w, "Fail list files", http.StatusInternalServerError)
		ctr.logger.Error("list file meta", zap.Error(err))
		return
	}
	response := struct {
		Data  []model.DBFile `json:"data"`
		Count int            `json:"count"`
	}{
		Data:  files,
		Count: len(files),
	}
	WriteResponse(w, http.StatusOK, response)
}
