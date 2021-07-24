package log

import (
	"io"
	"net/http"
	"os"
	"runtime"

	"go-template/setting"

	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// LumberJackLogger : Initialize logger
func LumberJackLogger(filePath string, maxSize int, maxBackups int, maxAge int) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxSize, // megabytes
		MaxBackups: maxBackups,
		MaxAge:     maxAge, //days
	}
}

// InitLogToStdoutDebug : set logrus to debug param
func InitLogToStdoutDebug() {
	logrus.SetFormatter(&prefixed.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

// InitLogToFile : set logrus to output in file
func InitLogToFile() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	out := LumberJackLogger(setting.Conf.DefaultSetting.LogPath, 500, 3, 28)
	logMultiWriter := io.MultiWriter(os.Stdout, out)
	logrus.SetOutput(logMultiWriter)
	logrus.SetLevel(logrus.WarnLevel)
}

// Init logrus
func init() {
	environment := os.Getenv("MODE")
	logrus.WithFields(logrus.Fields{
		"Runtime Version": runtime.Version(),
		"Number of CPUs":  runtime.NumCPU(),
		"Arch":            runtime.GOARCH,
	}).Info("Application Initializing")
	logrus.Info("environment", environment)
	switch environment {
	case "dev":
		InitLogToStdoutDebug()
	case "prod":
		InitLogToFile()
	}
	logrus.Debugf("Environment : %s", environment)
}

// Debug logs a message with debug log level.
func Debug(msg string) {
	logrus.Debug(msg)
}

// Debugf logs a formatted message with debug log level.
func Debugf(msg string, args ...interface{}) {
	logrus.Debugf(msg, args...)
}

// Info logs a message with info log level.
func Info(msg string) {
	logrus.Info(msg)
}

// Infof logs a formatted message with info log level.
func Infof(msg string, args ...interface{}) {
	logrus.Infof(msg, args...)
}

// Warn logs a message with warn log level.
func Warn(msg string) {
	logrus.Warn(msg)
}

// Warnf logs a formatted message with warn log level.
func Warnf(msg string, args ...interface{}) {
	logrus.Warnf(msg, args...)
}

// Error logs a message with error log level.
func Error(msg string) {
	logrus.Error(msg)
}

// Errorf logs a formatted message with error log level.
func Errorf(msg string, args ...interface{}) {
	logrus.Errorf(msg, args...)
}

// Fatal logs a message with fatal log level.
func Fatal(msg string) {
	logrus.Fatal(msg)
}

// Fatalf logs a formatted message with fatal log level.
func Fatalf(msg string, args ...interface{}) {
	logrus.Fatalf(msg, args...)
}

// Panic logs a message with panic log level.
func Panic(msg string) {
	logrus.Panic(msg)
}

// Panicf logs a formatted message with panic log level.
func Panicf(msg string, args ...interface{}) {
	logrus.Panicf(msg, args...)
}

// DebugResponse : log response body data for debugging
func DebugResponse(response *http.Response) string {
	bodyBuffer := make([]byte, 5000)
	var str string
	count, err := response.Body.Read(bodyBuffer)
	if err != nil {
		Debug(err.Error())
		return ""
	}
	for ; count > 0; count, err = response.Body.Read(bodyBuffer) {
		if err != nil {
			Debug(err.Error())
			continue
		}
		str += string(bodyBuffer[:count])
	}
	Debugf("response data : %v", str)
	return str
}
