package model

import (
	"mime/multipart"
	"time"
)

// File holds file metadata as a JSON
type File struct {
	File         multipart.File
	FileUserPath string
}

func NewFile(file multipart.File, fileUserPath string) *File {
	return &File{File: file, FileUserPath: fileUserPath}
}

type DBFile struct {
	FileId    int       `json:"id"`
	Filepath  string    `json:"filepath"`
	CreatedAt time.Time `json:"created_at"`
}
