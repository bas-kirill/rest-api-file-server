package service

import "go.uber.org/zap"

// FileWebService ...
type FileWebService struct {
	logger *zap.Logger
}

// NewFileWebService creates a new file web service
func NewFileWebService(logger *zap.Logger) *FileWebService {
	return &FileWebService{logger: logger}
}

// SaveFile ...
func (f *FileWebService) SaveFile(s string) error {
	//TODO implement me
	panic("implement me")
}

// GetFile ...
func (f *FileWebService) GetFile(s string) error {
	//TODO implement me
	panic("implement me")
}

// DeleteFile ...
func (f *FileWebService) DeleteFile(s string) error {
	//TODO implement me
	panic("implement me")
}
