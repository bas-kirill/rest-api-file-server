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

}
