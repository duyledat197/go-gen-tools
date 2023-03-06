package logger

import (
	"fmt"
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Replaced in tests.
var (
	stdout zapcore.WriteSyncer = os.Stdout
	stderr zapcore.WriteSyncer = os.Stderr
	atom                       = zap.NewAtomicLevelAt(zap.ErrorLevel)
)

func NewZapLogger(logLevel string, isLocalEnv bool) *zap.Logger {
	var (
		zapLogger *zap.Logger
		zapLogLvl zapcore.Level
	)
	err := zapLogLvl.Set(logLevel)
	if err != nil {
		log.Println("cannot parse logLevel, err:", err.Error())
		zapLogLvl = zap.WarnLevel
	}
	atom.SetLevel(zapLogLvl)
	if isLocalEnv {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		config.Level = atom
		zapLogger, err := config.Build()
		if err != nil {
			log.Println("cannot build logger, err:", err.Error())
		}
		return zapLogger
	}

	// Configure console output.
	consoleEncoder := newJSONEncoder()

	// We use zapcore.NewTee to direct logs of different levels to different outputs
	core := zapcore.NewTee(
		zapcore.NewCore(
			consoleEncoder,
			zapcore.Lock(stderr),
			zap.LevelEnablerFunc(func(l zapcore.Level) bool {
				return l >= zapcore.ErrorLevel && atom.Enabled(l)
			}),
		),
		zapcore.NewCore(
			consoleEncoder,
			zapcore.Lock(stdout),
			zap.LevelEnablerFunc(func(l zapcore.Level) bool {
				return l < zapcore.ErrorLevel && atom.Enabled(l)
			}),
		),
	)
	zapLogger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	zap.RedirectStdLog(zapLogger)
	return zapLogger
}

// Create a new JSON log encoder with the correct settings.
func newJSONEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "severity",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeLevel:    appendLogLevel,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

func appendLogLevel(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch l {
	case zapcore.DebugLevel:
		enc.AppendString("debug")
	case zapcore.InfoLevel:
		enc.AppendString("info")
	case zapcore.WarnLevel:
		enc.AppendString("warning")
	case zapcore.ErrorLevel:
		enc.AppendString("error")
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		enc.AppendString("critical")
	default:
		enc.AppendString(fmt.Sprintf("Level(%d)", l))
	}
}
