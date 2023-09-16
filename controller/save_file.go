package controller

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"rest-api-file-server/config"
	"rest-api-file-server/model"
	"rest-api-file-server/service"
)

// SaveFileController ...
type SaveFileController struct {
	logger           *zap.Logger
	fileServerConfig *config.FileServerConfig
	service          service.FileService
}

// NewSaveFileController creates a new file save controller
func NewSaveFileController(logger *zap.Logger, service service.FileService) *SaveFileController {
	return &SaveFileController{logger: logger, service: service}
}

// SaveFile create or update file by system path
func (s *SaveFileController) SaveFile(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, "Fail to parse file")
		return
	}
	defer file.Close()

	userFilePath := r.URL.String()
	err = s.service.SaveFile(model.NewFile(file, userFilePath))
	if err != nil {
		WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("Fail to save `%s`", userFilePath))
		return
	}

	WriteResponse(w, http.StatusOK, fmt.Sprintf("Saved `%s`", userFilePath))
}
