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

func TestIntegration_GivenFile_WhenSaveFile_ThenFileSaved(t *testing.T) {
	logger := zap.NewNop()
	appEnv := env.NewEnv()
	fileServerConfig := config.NewFileServerConfig(appEnv)
	pgConfig := config.NewPostgresConfig(appEnv)
	pgDb := pg.NewPgDatabase(pgConfig)
	pgMigration := store.NewPgMigrator(logger, pgConfig)
	pgMigration.RunMigrations()
	localFileContentService := service.NewLocalFileContentService(fileServerConfig, pgDb)
	uploadFileController := controller.NewUploadController(logger, localFileContentService)

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
	//userFilePath := "/file/folder/file.txt"
	uploadReq := newReq(http.MethodPut, "/file/folder/file.txt", body, map[string]string{"file-system-path": "/folder/file.txt"})
	uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	uploadFileController.Upload(w, uploadReq)

	// validate
	require.Equal(t, http.StatusOK, w.Result().StatusCode)
	message, err := io.ReadAll(w.Result().Body)
	require.NoError(t, err)
	require.NotEmpty(t, message)
}

func TestIntegration_GivenFileSaved_WhenSaveSameFile_ThenFileOverwritten(t *testing.T) {
	logger := zap.NewNop()
	appEnv := env.NewEnv()
	fileServerConfig := config.NewFileServerConfig(appEnv)
	pgConfig := config.NewPostgresConfig(appEnv)
	pgDb := pg.NewPgDatabase(pgConfig)
	pgMigration := store.NewPgMigrator(logger, pgConfig)
	pgMigration.RunMigrations()
	localFileContentService := service.NewLocalFileContentService(fileServerConfig, pgDb)
	uploadController := controller.NewUploadController(logger, localFileContentService)
	downloadController := controller.NewDownloadController(logger, fileServerConfig, localFileContentService)

	t.Cleanup(func() {
		_, err := pgDb.Db.Exec("truncate table files restart identity")
		if err != nil {
			panic(err)
		}
	})

	files := []struct {
		filename string
		content  string
	}{
		{
			filename: "data_1.txt",
			content:  "ping",
		},
		{
			filename: "data_2.txt",
			content:  "pong",
		},
	}

	var userFilePath = "/file/folder/file.txt"
	for _, file := range files {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		// create a new form-data header name data and filename data_1.txt
		dataPart, err := writer.CreateFormFile("file", file.filename)
		require.NoError(t, err)

		_, err = io.Copy(dataPart, strings.NewReader(file.content))
		require.NoError(t, err)
		require.NoError(t, writer.Close())

		// create http request & response
		uploadReq := newReq(http.MethodPut, "/file/folder/file.txt", body, map[string]string{"file-system-path": "/folder/file.txt"})
		uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()

		uploadController.Upload(w, uploadReq)

		// validate
		require.Equal(t, http.StatusOK, w.Result().StatusCode)
		message, err := io.ReadAll(w.Result().Body)
		require.NoError(t, err)
		require.NotEmpty(t, message)
	}

	downloadReq := newReq(http.MethodGet, userFilePath, nil, map[string]string{"id": "1"})
	w := httptest.NewRecorder()

	downloadController.Download(w, downloadReq)

	require.Equal(t, http.StatusOK, w.Result().StatusCode)
	response, err := io.ReadAll(w.Body)
	require.NoError(t, err)
	require.Equal(t, files[1].content, string(response))
}
