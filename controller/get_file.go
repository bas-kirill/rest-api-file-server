package controller

import (
	"go.uber.org/zap"
	"net/http"
	"rest-api-file-server/service"
)

type GetFileController struct {
	logger  *zap.Logger
	service service.FileService
}

// NewGetFileController creates a get file controller
func NewGetFileController(logger *zap.Logger, service service.FileService) *GetFileController {
	return &GetFileController{logger: logger, service: service}
}

// GetFile get file by system path
func (g *GetFileController) GetFile(w http.ResponseWriter, r *http.Request) {

}
