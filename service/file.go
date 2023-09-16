package service

import (
	"errors"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
	"rest-api-file-server/config"
	"rest-api-file-server/model"
)

// FileWebService ...
type FileWebService struct {
	logger           *zap.Logger
	fileServerConfig *config.FileServerConfig
}

// NewFileWebService creates a new file web service
func NewFileWebService(logger *zap.Logger, fileServerConfig *config.FileServerConfig) *FileWebService {
	return &FileWebService{logger: logger, fileServerConfig: fileServerConfig}
}

// SaveFile ...
func (f *FileWebService) SaveFile(file *model.File) error {
	destServerFilePath := filepath.Join(f.fileServerConfig.BaseSystemPath, file.FileSystemPath)
	destServerFileDir := filepath.Dir(destServerFilePath)
	err := os.MkdirAll(destServerFileDir, 0744)
	if err != nil {
		f.logger.Error("fail create folder", zap.String("folder", destServerFileDir), zap.Error(err))
		return err
	}

	destFile, err := os.Create(destServerFilePath)
	if err != nil {
		f.logger.Error("fail create file", zap.String("file", destServerFilePath), zap.Error(err))
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file.File)
	if err != nil {
		f.logger.Error("fail to copy file", zap.String("file", destServerFilePath), zap.Error(err))
		return err
	}

	return nil
}

// GetFile ...
func (f *FileWebService) GetFile(fileSystemPath string) (string, error) {
	serverFilePath := filepath.Join(f.fileServerConfig.BaseSystemPath, fileSystemPath)

	if !f.fileExists(serverFilePath) {
		return "", errors.New("file do not exist")
	}

	return serverFilePath, nil
}

// DeleteFile ...
func (f *FileWebService) DeleteFile(userFilePath string) error {
	serverFilePath := filepath.Join(f.fileServerConfig.BaseSystemPath, userFilePath)

	if !f.fileExists(serverFilePath) {
		return errors.New("file not found")
	}

	err := os.Remove(serverFilePath)
	if err != nil {
		return errors.New("fail remove file")
	}

	return nil
}

func (f *FileWebService) fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
