package logs

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	switch v := message.(type) {
	case map[string]interface{}:
		{
			jsonByte, _ := json.Marshal(v)
			fmt.Println("### LOG ###", string(jsonByte))
			log.Error(string(jsonByte), fields...)
		}
	case string:
		log.Error(v, fields...)
	case error:
		log.Error(v.Error(), fields...)
	}
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}
