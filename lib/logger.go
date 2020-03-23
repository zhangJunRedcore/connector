package lib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/spf13/viper"
)

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
	LogLevelSilent
	// LogLevelPanic
)

const (
	// Label Label
	Label = "DeepTun"
	// Phase Phase
	Phase = "trans"
)

// DebugLogger DebugLogger
type DebugLogger struct{}

// InfoLogger InfoLogger
type InfoLogger struct{}

// ErrorLogger ErrorLogger
type ErrorLogger struct{}

// Logger Logger
type Logger struct {
	Debug DebugLogger
	Info  InfoLogger
	Error ErrorLogger
}

var l zerolog.Logger
var output zerolog.ConsoleWriter

// InitLogger InitLogger
func NewLogger(level int, path string) *Logger {
	path = viper.GetString("log.dirpath")
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			// 必须分成两步：先创建文件夹、再修改权限
			os.MkdirAll(path, 0777) //0777也可以os.ModePerm
			os.Chmod(path, 0777)
		} else {
			panic(err)
		}
	}

	filePath := filepath.Join(path, "deepctl.log")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println("logger open file err:", err)
	}
	wr := diode.NewWriter(file, 1000, 0, func(missed int) {
		fmt.Printf(" %d /n", missed)
	})
	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	zerolog.CallerSkipFrameCount = 3 //跳过调用方法的堆栈帧数  22默认2
	zerolog.SetGlobalLevel(zerolog.Level(0))
	output = zerolog.ConsoleWriter{Out: wr, NoColor: true}

	output.FormatLevel = func(i interface{}) string {
		return fmt.Sprintf("level: %s,", i)
	}
	output.FormatTimestamp = func(i interface{}) string {
		return fmt.Sprintf("label: %s, timestamp: %s,", Label, time.Now().In(cstSh).Format("2006/01/02 15:04:05"))
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s: ", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s,", i)
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("message: %s", i)
	}

	l = zerolog.
		New(output).
		With().
		Logger()

	result := Logger{}
	return &result
}

// 修改配置文件的日志级别时，调用ModifyLevel，重新生成实例 logger
// [TO DO] 压测时，进行日志级别的修改测试
func ModifyLevel() {
	zerolog.SetGlobalLevel(zerolog.Level(viper.GetInt("log.level")))
	l = zerolog.
		New(output).
		With().
		Logger()
}

// Println Println
func (o *DebugLogger) Println(args ...interface{}) {
	l.Debug().Msg(fmt.Sprint(args...))
}

// Printf Printf
func (o *DebugLogger) Printf(format string, args ...interface{}) {
	l.Debug().Msgf(format, args...)
}

// Println Println
func (o *InfoLogger) Println(args ...interface{}) {
	l.Info().Msg(fmt.Sprint(args...))
}

// Printf Printf
func (o *InfoLogger) Printf(format string, args ...interface{}) {
	l.Info().Msgf(format, args...)
}

// Println Println
func (o *ErrorLogger) Println(args ...interface{}) {
	l.Error().Msg(fmt.Sprint(args...))
}

// Printf Printf
func (o *ErrorLogger) Printf(format string, args ...interface{}) {
	l.Error().Msgf(format, args...)
}
