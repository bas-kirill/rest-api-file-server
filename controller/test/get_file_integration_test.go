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
	fileWebService := service.NewFileWebService(fileServerConfig, pgDb)
	getFileController := controller.NewGetFileController(logger, fileServerConfig, fileWebService)

	req := httptest.NewRequest(http.MethodGet, "/file.txt", nil)
	w := httptest.NewRecorder()

	getFileController.GetFile(w, req)

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
	fileWebService := service.NewFileWebService(fileServerConfig, pgDb)
	getFileController := controller.NewGetFileController(logger, fileServerConfig, fileWebService)
	saveFileController := controller.NewSaveFileController(logger, fileWebService)

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
	userFilePath := "/folder/file.txt"
	req := httptest.NewRequest(http.MethodPut, userFilePath, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	saveFileController.SaveFile(w, req)

	req = httptest.NewRequest(http.MethodGet, userFilePath, nil)
	w = httptest.NewRecorder()

	getFileController.GetFile(w, req)

	require.Equal(t, http.StatusOK, w.Result().StatusCode)
	response, err := io.ReadAll(w.Body)
	require.NoError(t, err)
	require.Equal(t, file.content, string(response))

	t.Cleanup(func() {
		err = fileWebService.DeleteFile(userFilePath)
		if err != nil {
			panic(err)
		}
	})
}
