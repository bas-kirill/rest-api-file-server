package main

import (
	"fmt"
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
	"time"
)

func setUpLogger() *zap.Logger {
	// todo: set up log file prefix from credentials + custom logger for project
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
	logger := setUpLogger()

	fileWebService := service.NewFileWebService(logger, fileServerConfig)

	saveFileController := controller.NewSaveFileController(logger, fileWebService)
	getFileController := controller.NewGetFileController(logger, fileWebService)
	deleteFileController := controller.NewDeleteFileController(logger, fileWebService)

	httpLoggerMiddleware := middleware.NewHttpLoggerMiddleware(logger)

	router := mux.NewRouter()
	router.Use(httpLoggerMiddleware.Log)
	router.HandleFunc("/", saveFileController.SaveFile).Methods(http.MethodPost)
	router.HandleFunc("/{file-system-path}", getFileController.GetFile).Methods(http.MethodGet)
	router.HandleFunc("/{file-system-path}", deleteFileController.DeleteFile).Methods(http.MethodDelete)

	server := http.Server{
		Addr:         ":36000",
		Handler:      router,
		ReadTimeout:  httpServerConfig.ReadTimeout,
		WriteTimeout: httpServerConfig.WriteTimeout,
	}

	logger.Info("http server started", zap.String("server-address", server.Addr), zap.Any("config", httpServerConfig))
	if httpServerConfig.TlsEnabled {
		logger.Fatal("stop http server", zap.Error(server.ListenAndServeTLS(httpServerConfig.CertFile, httpServerConfig.CertKey)))
	} else {
		logger.Fatal("stop http server", zap.Error(server.ListenAndServe()))
	}
}
