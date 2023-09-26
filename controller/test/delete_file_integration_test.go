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

func TestIntegration_GivenFile_WhenDeleteFile_ThenFileDeleted(t *testing.T) {
	logger := zap.NewNop()
	appEnv := env.NewEnv()
	fileServerConfig := config.NewFileServerConfig(appEnv)
	pgConfig := config.NewPostgresConfig(appEnv)
	pgDb := pg.NewPgDatabase(pgConfig)
	pgMigration := store.NewPgMigrator(logger, pgConfig)
	pgMigration.RunMigrations()
	localFileContentService := service.NewLocalFileContentService(fileServerConfig, pgDb)
	uploadController := controller.NewUploadController(logger, localFileContentService)
	localFileMetaService := service.NewLocalFileMetaService(fileServerConfig, pgDb)
	deleteFileController := controller.NewDeleteFileController(logger, localFileMetaService)

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
	userFilePath := "/file/folder/file.txt"
	req := newReq(http.MethodPut, userFilePath, body, map[string]string{"file-system-path": "/folder/file.txt"})
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()

	uploadController.Upload(w, req)

	// validate
	require.Equal(t, http.StatusOK, w.Result().StatusCode)
	message, err := io.ReadAll(w.Result().Body)
	require.NoError(t, err)
	require.NotEmpty(t, message)

	deleteReq := newReq(http.MethodDelete, "/file/1", nil, map[string]string{"id": "1"})
	w = httptest.NewRecorder()

	deleteFileController.DeleteFile(w, deleteReq)

	require.Equal(t, http.StatusOK, w.Result().StatusCode)
}
