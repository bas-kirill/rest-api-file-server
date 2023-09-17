package controller

import (
	"go.uber.org/zap"
	"net/http"
	"rest-api-file-server/service"
)

type ListFiles struct {
	logger  *zap.Logger
	service service.FileService
}

func NewListFiles(logger *zap.Logger, service service.FileService) *ListFiles {
	return &ListFiles{logger: logger, service: service}
}

func (l *ListFiles) ListFiles(w http.ResponseWriter, r *http.Request) {
	files, err := l.service.ListFiles()
	if err != nil {
		http.Error(w, "Fail list files", http.StatusInternalServerError)
	}
	WriteResponse(w, http.StatusOK, files)
}
