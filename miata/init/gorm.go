package init

import (
	"gorm.io/gorm/schema"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type Mysql struct {
	Path         string `json:"path"`
	Dbname       string `json:"db-name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	LogLevel     string `json:"log-level"`
	LevelEncoder string `json:"level-encoder"`
}

// https://github.com/go-sql-driver/mysql#dsn-data-source-name
const ConnParams = "?charset=utf8mb4&parseTime=true&loc=Local&multiStatements=true"

func NewDB(config *Config, log *zap.Logger) *gorm.DB {
	cfg := config.Mysql
	log.Info("Establishing database connection",
		zap.String("host", cfg.Path),
		zap.String("database", cfg.Dbname),
		zap.String("username", cfg.Username))
	dsn := cfg.Username + ":" + cfg.Password + "@tcp(" + cfg.Path + ")/" + cfg.Dbname + ConnParams
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	logger := zapgorm2.New(cfg.GetZapLogger())

	var db *gorm.DB
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Error("DB 连接失败:", zap.String("err", err.Error()))
		return nil
	}

	return db
}

func (m *Mysql) GetZapLogger() *zap.Logger {
	var logger *zap.Logger
	var level zapcore.Level
	level.UnmarshalText([]byte(m.LogLevel))
	encoderCfg := getEncoderConfig(m.LevelEncoder)
	logCore := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), zapcore.AddSync(os.Stderr), level)

	if level == zap.ErrorLevel {
		logger = zap.New(logCore, zap.AddStacktrace(level))
	} else {
		logger = zap.New(logCore)
	}
	return logger
}
