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

func TestIntegration_GivenFiles_WhenListFiles_ThenReturnFilenames(t *testing.T) {
	logger := zap.NewNop()
	appEnv := env.NewEnv()
	fileServerConfig := config.NewFileServerConfig(appEnv)
	pgConfig := config.NewPostgresConfig(appEnv)
	pgDb := pg.NewPgDatabase(pgConfig)
	pgMigration := store.NewPgMigrator(logger, pgConfig)
	pgMigration.RunMigrations()
	localFileContentService := service.NewLocalFileContentService(fileServerConfig, pgDb)
	localFileMetaService := service.NewLocalFileMetaService(fileServerConfig, pgDb)
	uploadController := controller.NewUploadController(logger, localFileContentService)
	listFilesController := controller.NewListFiles(logger, localFileMetaService)

	t.Cleanup(func() {
		_, err := pgDb.Db.Exec("truncate table files restart identity")
		if err != nil {
			panic(err)
		}
	})

	files := []struct {
		filename     string
		userFileName string
		content      string
	}{
		{
			filename:     "data_1.txt",
			userFileName: "/folder/data_1.txt",
			content:      "Ilon Mask",
		},
		{
			userFileName: "/folder/data_2.txt",
			content:      "Marilyn Monroe",
		},
		{
			filename:     "data_3.txt",
			userFileName: "/folder/data_3.txt",
			content:      "Albert Einstein",
		},
	}

	for _, file := range files {
		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		// create a new form-data header name data and filename
		dataPart, err := writer.CreateFormFile("file", file.userFileName)
		require.NoError(t, err)

		_, err = io.Copy(dataPart, strings.NewReader(file.content))
		require.NoError(t, err)
		require.NoError(t, writer.Close())

		// create http request & response
		uploadReq := newReq(http.MethodPut, "/file"+file.userFileName, body, map[string]string{"file-system-path": file.userFileName})
		uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
		w := httptest.NewRecorder()

		uploadController.Upload(w, uploadReq)

		// validate
		require.Equal(t, http.StatusOK, w.Result().StatusCode)
		message, err := io.ReadAll(w.Result().Body)
		require.NoError(t, err)
		require.NotEmpty(t, message)
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	listFilesController.ListFiles(w, req)

	require.Equal(t, http.StatusOK, w.Result().StatusCode)
	response, err := io.ReadAll(w.Body)
	require.NoError(t, err)

	for _, file := range files {
		require.Contains(t, string(response), file.filename)
	}
}
