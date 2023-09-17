package controller

import (
	"go.uber.org/zap"
	"net/http"
	"rest-api-file-server/service"
)

type DeleteFileController struct {
	logger  *zap.Logger
	service service.FileService
}

// NewDeleteFileController create a delete file controller
func NewDeleteFileController(logger *zap.Logger, service service.FileService) *DeleteFileController {
	return &DeleteFileController{logger: logger, service: service}
}

// DeleteFile idempotent delete file by system path
func (d *DeleteFileController) DeleteFile(w http.ResponseWriter, r *http.Request) {
	userFilePath := r.URL.String()
	err := d.service.DeleteFile(userFilePath)
	if err != nil {
		if err.Error() == "file not found" {
			w.WriteHeader(http.StatusOK)
			return
		}
		d.logger.Error("delete file", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
