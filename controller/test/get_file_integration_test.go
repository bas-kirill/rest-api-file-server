package test

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"rest-api-file-server/config"
	"rest-api-file-server/controller"
	"rest-api-file-server/env"
	"rest-api-file-server/service"
	"rest-api-file-server/store"
	"rest-api-file-server/store/pg"
	"strings"
	"testing"
)

func TestIntegration_GivenNoFile_WhenGetFile_ThenReturnFileNotFound(t *testing.T) {
	logger := zap.NewNop()
	appEnv := env.NewEnv()
	fileServerConfig := config.NewFileServerConfig(appEnv)
	pgConfig := config.NewPostgresConfig(appEnv)
	pgDb := pg.NewPgDatabase(pgConfig)
	pgMigration := store.NewPgMigrator(logger, pgConfig)
	pgMigration.RunMigrations()
	fileWebService := service.NewLocalFileContentService(fileServerConfig, pgDb)
	getFileController := controller.NewDownloadController(logger, fileServerConfig, fileWebService)

	getFileContentReq := newReq(http.MethodGet, "/file/1/content", nil, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	getFileController.Download(w, getFileContentReq)

	require.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestIntegration_GivenFile_WhenGetFile_ThenReturnFile(t *testing.T) {
	logger := zap.NewNop()
	appEnv := env.NewEnv()
	fileServerConfig := config.NewFileServerConfig(appEnv)
	pgConfig := config.NewPostgresConfig(appEnv)
	pgDb := pg.NewPgDatabase(pgConfig)
	pgMigration := store.NewPgMigrator(logger, pgConfig)
	pgMigration.RunMigrations()
	localFileContentService := service.NewLocalFileContentService(fileServerConfig, pgDb)
	downloadController := controller.NewDownloadController(logger, fileServerConfig, localFileContentService)
	uploadController := controller.NewUploadController(logger, localFileContentService)

	t.Cleanup(func() {
		_, err := pgDb.Db.Exec("truncate table files restart identity")
		if err != nil {
			panic(err)
		}
	})

	file := struct {
		filename string
		content  string
	}{
		filename: "data_1.txt",
		content:  "some data",
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	// create a new form-data header name data and filename data_1.txt
	dataPart, err := writer.CreateFormFile("file", file.filename)
	require.NoError(t, err)

	// copy file content to multipart section
	_, err = io.Copy(dataPart, strings.NewReader(file.content))
	require.NoError(t, err)
	require.NoError(t, writer.Close())

	// create http request & response
	uploadReq := newReq(http.MethodPut, "/file/folder/file.txt", body, map[string]string{"file-system-path": "/folder/file.txt"})
	uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	uploadController.Upload(w, uploadReq)

	downloadReq := newReq(http.MethodGet, "/file/1/content", nil, map[string]string{"id": "1"})
	w = httptest.NewRecorder()

	downloadController.Download(w, downloadReq)

	require.Equal(t, http.StatusOK, w.Result().StatusCode)
	response, err := io.ReadAll(w.Body)
	require.NoError(t, err)
	require.Equal(t, file.content, string(response))
}
