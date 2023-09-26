package service

import (
	"rest-api-file-server/model"
)

// FileContentService for upload / download files
type FileContentService interface {
	Upload(file *model.File) error
	Download(int) (string, error)
}

// FileMetaService is a service for managing files
type FileMetaService interface {
	GetFileMeta(int) (*model.DBFile, error)
	DeleteFileMeta(int) error
	ListFilesMeta() ([]model.DBFile, error)
}
