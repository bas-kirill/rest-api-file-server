package service

import (
	"rest-api-file-server/model"
)

type FileService interface {
	SaveFile(file *model.File) error
	GetFile(string) (string, error)
	DeleteFile(string) error
	ListFiles() ([]string, error)
}
