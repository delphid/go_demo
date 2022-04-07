package init

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type Zap struct {
	Level        string `json:"level"`
	AddCaller    bool   `json:"add-caller"`
	LevelEncoder string `json:"level-encoder"`
}

func NewLogger(cfg *Config) *zap.Logger {
	logger := getLogger(cfg)
	return logger
}

func getLogger(cfg *Config) *zap.Logger {
	var logger *zap.Logger
	var level zapcore.Level
	level.UnmarshalText([]byte(cfg.Zap.Level))

	encoderCfg := getEncoderConfig(cfg.Zap.LevelEncoder)
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), zapcore.AddSync(os.Stderr), level)
	if level >= zap.WarnLevel {
		logger = zap.New(core, zap.AddStacktrace(level))
	} else {
		logger = zap.New(core)
	}
	if cfg.Zap.AddCaller {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

func getEncoderConfig(levelEncoder string) (cfg zapcore.EncoderConfig) {
	cfg = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006/01/02 15:04:05.000"),
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	switch {
	case levelEncoder == "LowercaseLevelEncoder": // 小写编码器(默认)
		cfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	case levelEncoder == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		cfg.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case levelEncoder == "CapitalLevelEncoder": // 大写编码器
		cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	case levelEncoder == "CapitalColorLevelEncoder": // 大写编码器带颜色
		cfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		cfg.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return cfg
}
