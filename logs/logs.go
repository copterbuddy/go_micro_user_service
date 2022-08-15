package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	config := zap.NewDevelopmentConfig()

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zapcore.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zapcore.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zapcore.Field) {

	switch v := message.(type) {
	case error: //error interface
		log.Error(v.Error(), fields...)
	case string: // error message
		log.Error(v)
	}
}

/* 

***Log Error interface อย่างเดียว***
msg,ok := message.(error)
	if ok {
		msg.Error()
	}
*********************************
	
***Log Typeตามการใช้งาน***
Log, _ = zap.NewProduction()
Log, _ = zap.NewDevelopment()

config := zap.NewProductionConfig() // ใช้ใน prod เป็น ไฟล์
config := zap.NewDevelopmentConfig() //ใช้ใน dev เป็น console log
*************************

*/
