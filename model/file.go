package model

import "mime/multipart"

type File struct {
	File         multipart.File
	FileUserPath string
}

func NewFile(file multipart.File, fileUserPath string) *File {
	return &File{File: file, FileUserPath: fileUserPath}
}
