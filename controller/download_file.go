package controller

import (
	"database/sql"
	"errors"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"rest-api-file-server/config"
	"rest-api-file-server/service"
	"strconv"
)

type DownloadController struct {
	logger           *zap.Logger
	fileServerConfig *config.FileServerConfig
	service          service.FileContentService
}

// NewDownloadController creates a get file controller
func NewDownloadController(logger *zap.Logger, fileServerConfig *config.FileServerConfig, service service.FileContentService) *DownloadController {
	return &DownloadController{logger: logger, fileServerConfig: fileServerConfig, service: service}
}

// Download get file by path
func (ctr *DownloadController) Download(w http.ResponseWriter, r *http.Request) {
	fileId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "fail to parse file id", http.StatusInternalServerError)
		ctr.logger.Error("download file", zap.Error(err))
		return
	}

	serverFilePath, err := ctr.service.Download(fileId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "File not found", http.StatusNotFound)
			ctr.logger.Error("fail download", zap.Error(err))
			return
		}
		http.Error(w, "Fail download file", http.StatusInternalServerError)
		ctr.logger.Error("download file", zap.Error(err))
		return
	}

	http.ServeFile(w, r, serverFilePath)
}
