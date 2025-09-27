package zap

import (
	"dream-vue-admin/global"
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type roccLogWriter struct {
	mu         *sync.Mutex
	currentDay string
	file       *os.File
	logDir     string
}

func (w *roccLogWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	currentDay := time.Now().Format("2006-01-02")

	//判断日期
	if currentDay != w.currentDay {
		if w.file != nil {
			_ = w.file.Close()
		}
		path := global.Cfg.Server.LogPath
		filePath := fmt.Sprintf("%s/%s.log", path, currentDay)
		_, err = os.Stat(path)
		if err != nil {
			_ = os.MkdirAll(path, os.ModePerm)
		}
		w.file, _ = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		w.currentDay = currentDay
	}
	return w.file.Write(p)
}

func InitZapLog() *zap.SugaredLogger {

	cfg := zap.NewDevelopmentConfig()
	zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	var writer = roccLogWriter{
		mu:     &sync.Mutex{},
		logDir: global.Cfg.Server.LogPath,
	}

	level, err := zapcore.ParseLevel(global.Cfg.Server.LogLevel)
	if err != nil {
		level = zapcore.InfoLevel
	}

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(os.Stdout),
		level,
	)

	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(&writer),
		level,
	)

	core := zapcore.NewTee(consoleCore, fileCore)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)

	return logger.Sugar()
}
