package controller

import (
	"go.uber.org/zap"
	"net/http"
	"rest-api-file-server/config"
	"rest-api-file-server/service"
)

type GetFileController struct {
	logger           *zap.Logger
	fileServerConfig *config.FileServerConfig
	service          service.FileService
}

// NewGetFileController creates a get file controller
func NewGetFileController(logger *zap.Logger, fileServerConfig *config.FileServerConfig, service service.FileService) *GetFileController {
	return &GetFileController{logger: logger, fileServerConfig: fileServerConfig, service: service}
}

// GetFile get file by path
func (g *GetFileController) GetFile(w http.ResponseWriter, r *http.Request) {
	userFilePath := r.URL.String()

	serverFilePath, err := g.service.GetFile(userFilePath)
	if err != nil {
		WriteResponse(w, http.StatusNotFound, "File not found")
		return
	}

	http.ServeFile(w, r, serverFilePath)
}
