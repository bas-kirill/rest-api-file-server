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

type DeleteFileController struct {
	logger  *zap.Logger
	service service.FileMetaService
}

// NewDeleteFileController create a delete file controller
func NewDeleteFileController(logger *zap.Logger, service service.FileMetaService) *DeleteFileController {
	return &DeleteFileController{logger: logger, service: service}
}

// DeleteFile idempotent delete file by system path
func (ctr *DeleteFileController) DeleteFile(w http.ResponseWriter, r *http.Request) {
	fileId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Fail delete file", http.StatusInternalServerError)
		ctr.logger.Error("delete file", zap.Error(err))
		return
	}
	err = ctr.service.DeleteFileMeta(fileId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			w.WriteHeader(http.StatusOK) // need for idempotency
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		ctr.logger.Error("delete file", zap.Error(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}
