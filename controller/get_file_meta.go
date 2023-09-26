package controller

import (
	"database/sql"
	"errors"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"rest-api-file-server/service"
	"strconv"
)

type GetFileMeta struct {
	logger  *zap.Logger
	service service.FileMetaService
}

func NewGetFileMeta(logger *zap.Logger, service service.FileMetaService) *ListFilesMeta {
	return &ListFilesMeta{logger: logger, service: service}
}

func (ctr *ListFilesMeta) GetFileMeta(w http.ResponseWriter, r *http.Request) {
	fileId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctr.logger.Error("get file meta", zap.Error(err))
		return
	}

	file, err := ctr.service.GetFileMeta(fileId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "File not found", http.StatusNotFound)
			ctr.logger.Error("get file meta", zap.Error(err))
			return
		}
		http.Error(w, "Fail get file meta", http.StatusInternalServerError)
		ctr.logger.Error("get file meta", zap.Error(err))
		return
	}
	WriteResponse(w, http.StatusOK, file)
}
