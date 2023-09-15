package service

type FileService interface {
	SaveFile(string) error
	GetFile(string) error
	DeleteFile(string) error
}
