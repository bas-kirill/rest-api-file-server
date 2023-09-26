package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"rest-api-file-server/config"
	"rest-api-file-server/model"
	"rest-api-file-server/service"
)

// UploadController ...
type UploadController struct {
	logger           *zap.Logger
	fileServerConfig *config.FileServerConfig
	service          service.FileContentService
}

// NewUploadController creates a new file save controller
func NewUploadController(logger *zap.Logger, service service.FileContentService) *UploadController {
	return &UploadController{logger: logger, service: service}
}

// Upload create or update file by system path
func (ctr *UploadController) Upload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		ctr.logger.Error("upload file", zap.Error(err))
		WriteResponse(w, http.StatusBadRequest, "Fail to parse file")
		return
	}
	defer file.Close()

	userFilePath := mux.Vars(r)["file-system-path"]
	err = ctr.service.Upload(model.NewFile(file, userFilePath))
	if err != nil {
		ctr.logger.Error("upload file", zap.Error(err))
		WriteResponse(w, http.StatusInternalServerError, fmt.Sprintf("Fail to save `%s`", userFilePath))
		return
	}

	WriteResponse(w, http.StatusOK, fmt.Sprintf("Saved `%s`", userFilePath))
}
