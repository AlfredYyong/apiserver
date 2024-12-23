package config

import (
	"log/slog"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/lexkong/log"
)

const (
	OutputFile   = "file"
	OutputStdout = "stdout"

	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
)

type Config struct {
	LogConfig
	Name string
}

type LogConfig struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	Level      string
	Output     string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// init config
	if err := c.initConfig(); err != nil {
		return err
	}

	// 初始化日志包
	c.initLog()

	//
	c.watchConfig()

	return nil
}

// initConfig 初始化配置文件
func (c Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.AddConfigPath("conf") // 如果没有指定配置文件，则解析默认的配置文件 conf/config.yaml
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv()            // 读取匹配的环境变量
	viper.SetEnvPrefix("APISERVER") // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),        // 输出位置，可选file, stdout
		LoggerLevel:    viper.GetString("log.logger_level"),   // 日志级别
		LoggerFile:     viper.GetString("log.logger_file"),    // 日志文件
		LogFormatText:  viper.GetBool("log.log_format_text"),  // 日志的输出格式
		RollingPolicy:  viper.GetString("log.rolling_policy"), // rotate 依据，可选有daily和size
		LogRotateDate:  viper.GetInt("log.rotate_date"),       // rotate 转存时间
		LogRotateSize:  viper.GetInt("log.rotate_size"),       // rotate 转存大小
		LogBackupCount: viper.GetInt("log.backup_count"),      // 压缩
	}

	log.InitWithConfig(&passLagerCfg)
}

// initSlog 初始化 slog 配置
func (c *Config) initSlog() {
	conf := &LogConfig{
		FileName:   c.FileName,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge,
		Compress:   c.Compress,
		Level:      c.Level,
		Output:     c.Output,
	}

	//设置日志级别
	level := getLogLevel(conf.Level)
	opts := &slog.HandlerOptions{Level: level}

	//设置日志切割
	lumberjackLogger := &lumberjack.Logger{
		Filename:   conf.FileName,   // 日志文件名
		MaxSize:    conf.MaxSize,    // 单个日志文件最大大小（MB）
		MaxBackups: conf.MaxBackups, // 最多保留多少个备份文件
		MaxAge:     conf.MaxAge,     // 文件最多保存多少天
		Compress:   conf.Compress,   // 是否压缩旧文件
	}
	//设置日志输出方式
	if conf.Output == OutputFile {
		//日志文件输出
		fileHandler := slog.NewTextHandler(lumberjackLogger, opts)
		slog.SetDefault(slog.New(fileHandler))
	} else {
		//日志控制台输出
		consoleHandler := slog.NewTextHandler(os.Stdout, opts)
		slog.SetDefault(slog.New(consoleHandler))
	}
}

// watchConfig 监控配置文件变化并热加载程序
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s", e.Name)
	})
}

// getLogLevel 获取 slog 日志级别
func getLogLevel(level string) *slog.LevelVar {
	var lvl = new(slog.LevelVar)
	switch level {
	case LevelDebug:
		lvl.Set(slog.LevelDebug)
	case LevelInfo:
		lvl.Set(slog.LevelInfo)
	case LevelWarn:
		lvl.Set(slog.LevelWarn)
	case LevelError:
		lvl.Set(slog.LevelError)
	default:
		lvl.Set(slog.LevelInfo)
	}
	return lvl
}
