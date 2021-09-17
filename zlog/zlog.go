package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
)

func checkFile(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter(filenames string) zapcore.WriteSyncer {
	idx := strings.LastIndex(filenames, "/")

	filepath := filenames
	if idx != -1 && filepath != "." {
		filepath = filenames[:idx]
		if isExist, _ := checkFile(filepath); !isExist {
			err := os.MkdirAll(filepath, os.ModePerm)
			if err != nil {
				log.Fatalln("mkdir failed ", err)
			}
		}
	}

	file, err := os.OpenFile(filenames, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败：", err)
	}
	return zapcore.AddSync(file)
}

func Info(filenames, msg string) {
	writeSyncer := getLogWriter(filenames)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core)
	sugar := logger.Sugar()

	sugar.Info(msg)
}
