package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"rest-api-file-server/config"
	"rest-api-file-server/controller"
	"rest-api-file-server/env"
	"rest-api-file-server/middleware"
	"rest-api-file-server/service"
	"rest-api-file-server/store"
	"rest-api-file-server/store/pg"
	"time"
)

func setUpLogger() *zap.Logger {
	consoleLevel := zapcore.DebugLevel
	consoleLogConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoder := zapcore.NewConsoleEncoder(consoleLogConfig)

	fileLogConfig := zap.NewProductionEncoderConfig()
	fileLogConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileLogConfig)
	consoleOutput := zapcore.Lock(os.Stdout)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleOutput, consoleLevel)

	logFilePath := fmt.Sprintf("./log/%s.log", time.Now().Format(time.RFC3339))
	fileOutput, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("fail to open log file: " + err.Error())
	}
	fileLevel := zapcore.DebugLevel
	fileCore := zapcore.NewCore(fileEncoder, zapcore.AddSync(fileOutput), fileLevel)

	core := zapcore.NewTee(consoleCore, fileCore)
	logger := zap.New(core)
	zap.ReplaceGlobals(logger)
	return logger
}

func main() {
	appEnv := env.NewEnv()
	httpServerConfig := config.NewHttpServerConfig(appEnv)
	fileServerConfig := config.NewFileServerConfig(appEnv)
	pgConfig := config.NewPostgresConfig(appEnv)
	logger := setUpLogger()

	pgMigration := store.NewPgMigrator(logger, pgConfig)
	pgMigration.RunMigrations()

	pgDb := pg.NewPgDatabase(pgConfig)
	localFileContentService := service.NewLocalFileContentService(fileServerConfig, pgDb)
	localFileMetaService := service.NewLocalFileMetaService(fileServerConfig, pgDb)

	uploadFileController := controller.NewUploadController(logger, localFileContentService)
	downloadFileController := controller.NewDownloadController(logger, fileServerConfig, localFileContentService)

	getFileMetaController := controller.NewGetFileMeta(logger, localFileMetaService)
	deleteFileController := controller.NewDeleteFileController(logger, localFileMetaService)
	listFilesController := controller.NewListFiles(logger, localFileMetaService)

	httpLoggerMiddleware := middleware.NewHttpLoggerMiddleware(logger)

	router := mux.NewRouter()
	router.Use(httpLoggerMiddleware.Log)
	router.HandleFunc("/file/{file-system-path:.+}", uploadFileController.Upload).Methods(http.MethodPut)
	router.HandleFunc("/file/{id}/content", downloadFileController.Download).Methods(http.MethodGet)

	router.HandleFunc("/file", listFilesController.ListFiles).Methods(http.MethodGet)
	router.HandleFunc("/file/{id}", getFileMetaController.GetFileMeta).Methods(http.MethodGet)
	router.HandleFunc("/file/{id}", deleteFileController.DeleteFile).Methods(http.MethodDelete)

	methods := handlers.AllowedMethods([]string{"GET"})
	origins := handlers.AllowedOrigins([]string{"*"}) // todo: add TLS for frontend
	httpServer := http.Server{
		Addr:         httpServerConfig.HttpAddr,
		Handler:      handlers.CORS(methods, origins)(router),
		ReadTimeout:  httpServerConfig.ReadTimeout,
		WriteTimeout: httpServerConfig.WriteTimeout,
	}
	if httpServerConfig.TlsEnabled {
		logger.Info("http+tls server started", zap.Any("config", httpServerConfig))
		logger.Fatal("stop http+tls server", zap.Error(httpServer.ListenAndServeTLS(httpServerConfig.CertFile, httpServerConfig.CertKey)))
	} else {
		logger.Info("http server started", zap.Any("config", httpServerConfig))
		logger.Fatal("stop http server", zap.Error(httpServer.ListenAndServe()))
	}
}
