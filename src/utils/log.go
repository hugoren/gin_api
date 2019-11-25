package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


func Log(appName string) *zap.Logger{
	/*rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "alarm.log"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"appName":"`+appName+`"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)
	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return logger

	*/

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:"Time",
		LevelKey:"Level",
		NameKey: "Logger",
		CallerKey: "Caller",
		MessageKey:"Msg",
		LineEnding: zapcore.DefaultLineEnding,
		EncodeLevel: zapcore.LowercaseLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller: zapcore.FullCallerEncoder,
	}
	atom := zap.NewAtomicLevelAt(zap.DebugLevel)
	config := zap.Config{
		Level:atom,
		Development:false,
		DisableCaller:false,
		Encoding: "json",
		EncoderConfig: encoderConfig,
		InitialFields:map[string]interface{}{"appName": appName},
		OutputPaths:[]string{"stdout","alarm.log"},
		ErrorOutputPaths: []string{"stdout","alarm.log"},
	}
	logger, err := config.Build()
	if err != nil {
		logger.Error("logger init failed"+err.Error())
	}
	//logger.Info("logger init success")

	return logger
}
