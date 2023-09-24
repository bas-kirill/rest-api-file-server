package service

import (
	"database/sql"
	"errors"
	"fmt"
	"rest-api-file-server/config"
	"rest-api-file-server/model"
	"rest-api-file-server/store/pg"
	"time"
)

type LocalFileMetaService struct {
	fileServerConfig *config.FileServerConfig
	pg               *pg.DB
}

func NewLocalFileMetaService(fileServerConfig *config.FileServerConfig, pg *pg.DB) *LocalFileMetaService {
	return &LocalFileMetaService{fileServerConfig: fileServerConfig, pg: pg}
}

// GetFileMeta get meta information about file
func (f LocalFileMetaService) GetFileMeta(fileId int) (*model.DBFile, error) {
	tx, err := f.pg.Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail to open transaction to list files: %v", err)
	}

	listFilesSql := `select file_id, filename, created_at from files where file_id = $1`

	row := tx.QueryRow(listFilesSql, fileId)
	var dbFile model.DBFile
	err = row.Scan(&dbFile.FileId, &dbFile.Filepath, &dbFile.CreatedAt)
	if err != nil {
		_ = tx.Rollback()
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("fail to scan file row: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("fail to commit tx during deleting file: %v", err)
	}

	return &dbFile, nil
}

// DeleteFileMeta delete file form server and file name from db (no soft delete)
func (f *LocalFileMetaService) DeleteFileMeta(fileId int) error {
	err := f.hardDeleteFile(fileId)
	if err != nil {
		return err
	}

	return nil
}

func (f *LocalFileMetaService) hardDeleteFile(fileId int) error {
	tx, err := f.pg.Db.Begin()
	if err != nil {
		return fmt.Errorf("fail to open transaction to delete file name: %v", err)
	}

	hardDeleteFileNameSql := `
		delete from files
		where file_id = $1
		returning created_at`

	var createdAt time.Time
	row := tx.QueryRow(hardDeleteFileNameSql, fileId)
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

// ListFilesMeta list meta information about every file
func (f *LocalFileMetaService) ListFilesMeta() ([]model.DBFile, error) {
	tx, err := f.pg.Db.Begin()
	if err != nil {
		return nil, fmt.Errorf("fail to open transaction to list files: %v", err)
	}

	listFilesSql := `select file_id, filename, created_at from files`

	rows, err := tx.Query(listFilesSql)
	if err != nil {
		return nil, fmt.Errorf("fail to execute query: %s", listFilesSql)
	}
	dbFiles := make([]model.DBFile, 0)
	for rows.Next() {
		var dbFile model.DBFile
		err = rows.Scan(&dbFile.FileId, &dbFile.Filepath, &dbFile.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("fail to scan file name: %v", err)
		}
		dbFiles = append(dbFiles, dbFile)
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		return nil, fmt.Errorf("fail to commit tx during deleting file: %v", err)
	}

	return dbFiles, nil
}
