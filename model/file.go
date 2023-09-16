package model

import "mime/multipart"

type File struct {
	File           multipart.File
	FileSystemPath string
}

func NewFile(file multipart.File, fileSystemPath string) *File {
	return &File{File: file, FileSystemPath: fileSystemPath}
}
