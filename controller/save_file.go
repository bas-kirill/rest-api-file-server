package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"rest-api-file-server/service"
)

// SaveFileController ...
type SaveFileController struct {
	logger  *zap.Logger
	service service.FileService
}

// NewSaveFileController creates a new file controller
func NewSaveFileController(logger *zap.Logger, service service.FileService) *SaveFileController {
	return &SaveFileController{logger: logger, service: service}
}

// SaveFile idempotency create new file by system path
func (s *SaveFileController) SaveFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileSystemPath := vars["file-system-path"]
	WriteResponse(w, http.StatusOK, fmt.Sprintf("Got `%s` file system path", fileSystemPath))
}
