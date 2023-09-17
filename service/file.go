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

// FileWebService ...
type FileWebService struct {
	fileServerConfig *config.FileServerConfig
	pg               *pg.DB
}

// NewFileWebService creates a new file web service
func NewFileWebService(fileServerConfig *config.FileServerConfig, pg *pg.DB) *FileWebService {
	return &FileWebService{fileServerConfig: fileServerConfig, pg: pg}
}

// SaveFile ...
func (f *FileWebService) SaveFile(file *model.File) error {
	destServerFilePath := filepath.Join(f.fileServerConfig.BaseSystemPath, file.FileUserPath)
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

	err = f.persistFileName(file.FileUserPath)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileWebService) persistFileName(fileName string) error {
	tx, err := f.pg.Db.Begin()
	if err != nil {
		return fmt.Errorf("fail to open transaction during saving file: %v", err)
	}

	insertNewFileNameSql := `
		insert into files (filename)
		values ($1)
		on conflict do nothing
		returning created_at`

	var createdAt time.Time
	row := tx.QueryRow(insertNewFileNameSql, fileName)
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

// GetFile ...
func (f *FileWebService) GetFile(fileSystemPath string) (string, error) {
	serverFilePath := filepath.Join(f.fileServerConfig.BaseSystemPath, fileSystemPath)

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

// DeleteFile delete file form server and file name from db (no soft delete)
func (f *FileWebService) DeleteFile(userFilePath string) error {
	serverFilePath := filepath.Join(f.fileServerConfig.BaseSystemPath, userFilePath)

	_, err := os.Stat(serverFilePath)
	if os.IsNotExist(err) {
		return errors.New("file not found")
	}

	err = os.Remove(serverFilePath)
	if err != nil {
		return errors.New("fail remove file")
	}

	err = f.hardDeleteFile(userFilePath)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileWebService) hardDeleteFile(fileName string) error {
	tx, err := f.pg.Db.Begin()
	if err != nil {
		return fmt.Errorf("fail to open transaction to delete file name: %v", err)
	}

	hardDeleteFileNameSql := `
		delete from files
		where filename = $1
		returning created_at`

	var createdAt time.Time
	row := tx.QueryRow(hardDeleteFileNameSql, fileName)
	err = row.Scan(&createdAt)
	if err != nil {
		_ = tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			// it is correct situation because we want to keep idempotency from API pov
			return nil
		}
		return fmt.Errorf("fail to delete: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("fail to commit tx during deleting file: %v", err)
	}

	return nil
}

func (f *FileWebService) ListFiles() ([]string, error) {
	tx, err := f.pg.Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail to open transaction to list files: %v", err)
	}

	listFilesSql := `select filename from files`

	var fileName string
	rows, err := tx.Query(listFilesSql)
	if err != nil {
		return nil, fmt.Errorf("fail to execute query: %s", listFilesSql)
	}
	fileNames := make([]string, 0)
	for rows.Next() {
		err = rows.Scan(&fileName)
		if err != nil {
			return nil, fmt.Errorf("fail to scan file name: %v", err)
		}
		fileNames = append(fileNames, fileName)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("fail to commit tx during deleting file: %v", err)
	}

	return fileNames, nil
}
