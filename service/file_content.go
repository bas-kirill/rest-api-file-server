package service

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"rest-api-file-server/config"
	"rest-api-file-server/model"
	"rest-api-file-server/store/pg"
	"time"
)

// LocalFileContentService ...
type LocalFileContentService struct {
	fileServerConfig *config.FileServerConfig
	pg               *pg.DB
}

// NewLocalFileContentService creates a new file local service
func NewLocalFileContentService(fileServerConfig *config.FileServerConfig, pg *pg.DB) *LocalFileContentService {
	return &LocalFileContentService{fileServerConfig: fileServerConfig, pg: pg}
}

// Upload upload file content to local storage
func (svc *LocalFileContentService) Upload(file *model.File) error {
	destServerFilePath := filepath.Join(svc.fileServerConfig.BaseSystemPath, file.FileUserPath)
	destServerFileDir := filepath.Dir(destServerFilePath)
	err := os.MkdirAll(destServerFileDir, 0744)
	if err != nil {
		return err
	}

	destFile, err := os.Create(destServerFilePath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, file.File)
	if err != nil {
		return err
	}

	err = svc.persistFile(file.FileUserPath)
	if err != nil {
		return err
	}

	return nil
}

func (svc *LocalFileContentService) persistFile(fileName string) error {
	tx, err := svc.pg.Db.Begin()
	if err != nil {
		return fmt.Errorf("fail to open transaction during saving file: %v", err)
	}

	insertNewFileSql := `
		insert into files (filename)
		values ($1)
		on conflict do nothing
		returning created_at`

	var createdAt time.Time
	row := tx.QueryRow(insertNewFileSql, fileName)
	err = row.Scan(&createdAt)
	if err != nil {
		_ = tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			// it is correct situation because we want to keep idempotency from API pov
			return nil
		}
		return fmt.Errorf("fail to scan to file id: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("fail to commit tx during saving file: %v", err)
	}

	return nil
}

// Download downloads file content from local storage
func (svc *LocalFileContentService) Download(fileId int) (string, error) {
	dbFile, err := svc.findById(fileId)
	if err != nil {
		return "", err
	}
	serverFilePath := filepath.Join(svc.fileServerConfig.BaseSystemPath, dbFile.Filepath)

	fileInfo, err := os.Stat(serverFilePath)
	if os.IsNotExist(err) {
		return "", errors.New("file do not exist")
	}

	if fileInfo.IsDir() {
		return "", errors.New("path is dir")
	}

	// now file is filename
	return serverFilePath, nil
}

func (svc *LocalFileContentService) findById(fileId int) (*model.DBFile, error) {
	tx, err := svc.pg.Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail to open transaction during saving file: %v", err)
	}

	findByIdSql := `select file_id, filename, created_at from files where file_id = $1`

	var dbFile model.DBFile
	row := tx.QueryRow(findByIdSql, fileId)
	err = row.Scan(&dbFile.FileId, &dbFile.Filepath, &dbFile.CreatedAt)
	if err != nil {
		_ = tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("fail to scan to file id: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("fail to commit tx during saving file: %v", err)
	}

	return &dbFile, nil
}
