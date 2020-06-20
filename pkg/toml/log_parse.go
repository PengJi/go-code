package main

import (
	"fmt"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/tsthght/BladeAudit/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)
const (
	defaultLogMaxSize = 300 // MB
)

type LogConfig struct {
	Conf Config `toml:"log"`
}

// Config serializes log related config in toml/json.
type Config struct {
	// Log level.
	Level string `toml:"level" json:"level"`
	// Log format. one of json, text, or console.
	Format string `toml:"format" json:"format"`
	// Disable automatic timestamps in output.
	DisableTimestamp bool `toml:"disable-timestamp" json:"disable-timestamp"`
	// File log config.
	File FileLogConfig `toml:"file" json:"file"`
	// Development puts the logger in development mode, which changes the
	// behavior of DPanicLevel and takes stacktraces more liberally.
	Development bool `toml:"development" json:"development"`
	// DisableCaller stops annotating logs with the calling function's file
	// name and line number. By default, all logs are annotated.
	DisableCaller bool `toml:"disable-caller" json:"disable-caller"`
	// DisableStacktrace completely disables automatic stacktrace capturing. By
	// default, stacktraces are captured for WarnLevel and above logs in
	// development and ErrorLevel and above in production.
	DisableStacktrace bool `toml:"disable-stacktrace" json:"disable-stacktrace"`
	// DisableErrorVerbose stops annotating logs with the full verbose error
	// message.
	DisableErrorVerbose bool `toml:"disable-error-verbose" json:"disable-error-verbose"`
	// SamplingConfig sets a sampling strategy for the logger. Sampling caps the
	// global CPU and I/O load that logging puts on your process while attempting
	// to preserve a representative subset of your logs.
	//
	// Values configured here are per-second. See zapcore.NewSampler for details.
	Sampling *zap.SamplingConfig `toml:"sampling" json:"sampling"`
}

// FileLogConfig serializes file log related config in toml/json.
type FileLogConfig struct {
	// Log filename, leave empty to disable file log.
	Filename string `toml:"filename" json:"filename"`
	// Max size for a single file, in MB.
	MaxSize int `toml:"max-size" json:"max-size"`
	// Max log keep days, default is never deleting.
	MaxDays int `toml:"max-days" json:"max-days"`
	// Maximum number of old log files to retain.
	MaxBackups int `toml:"max-backups" json:"max-backups"`
}

func (cfg *Config) buildOptions(errSink zapcore.WriteSyncer) []zap.Option {
	opts := []zap.Option{zap.ErrorOutput(errSink)}

	if cfg.Development {
		opts = append(opts, zap.Development())
	}

	if !cfg.DisableCaller {
		opts = append(opts, zap.AddCaller())
	}

	stackLevel := zap.ErrorLevel
	if cfg.Development {
		stackLevel = zap.WarnLevel
	}
	if !cfg.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(stackLevel))
	}

	if cfg.Sampling != nil {
		opts = append(opts, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewSampler(core, time.Second, int(cfg.Sampling.Initial), int(cfg.Sampling.Thereafter))
		}))
	}

	return opts
}

type ZapProperties struct {
	Core   zapcore.Core
	Syncer zapcore.WriteSyncer
	Level  zap.AtomicLevel
}

func initLogger(cfg *Config, opts ...zap.Option) (*zap.Logger, *ZapProperties, error) {
	// set log level
	level := zap.NewAtomicLevel()
	// zapcore.Levels ("debug", "info", "warn", "error", "dpanic", "panic", and "fatal").
	err := level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return nil, nil, err
	}

	// get log writer
	writeSyncer, err := getLogWriter(&cfg.File)
	if err != nil {
		return nil, nil, err
	}

	// get log core
	encoder := getEncoder(cfg)
	core := zapcore.NewCore(encoder, writeSyncer, level)
	opts = append(cfg.buildOptions(writeSyncer), opts...)
	logger := zap.New(core, opts...)

	r := &ZapProperties{
		Core:   core,
		Syncer: writeSyncer,
		Level:  level,
	}

	return logger, r, nil
}

// initFileLog initializes file based logging options.
func getLogWriter(cfg *FileLogConfig) (zapcore.WriteSyncer, error) {
	if len(cfg.Filename) < 0{
		stdOut, close, err := zap.Open([]string{"stdout"}...)
		if err != nil {
			close()
			return nil, err
		}
		return stdOut, nil
	}

	if st, err := os.Stat(cfg.Filename); err == nil {
		if st.IsDir() {
			return nil, errors.New("this is a directory")
		}
	}
	if cfg.MaxSize == 0 {
		cfg.MaxSize = defaultLogMaxSize
	}

	lg := &lumberjack.Logger{
		Filename:   cfg.Filename,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxDays,
		LocalTime:  true,
	}
	logWriter := zapcore.AddSync(lg)
	return logWriter, nil

}

func getEncoder(cfg *Config) zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func GetLogger() (*zap.Logger, *ZapProperties, error){
	var logConf LogConfig
	var logFile = "log_config.toml"
	if _, err := toml.DecodeFile(logFile, &logConf); err != nil {
		fmt.Println(err)
	}

	conf := &logConf.Conf
	lg, r, err := initLogger(conf)
	return lg, r, err
}

func main(){
	var logConf LogConfig
	var logFile = "log_config.toml"
	if _, err := toml.DecodeFile(logFile, &logConf); err != nil {
		fmt.Println(err)
	}

	fmt.Println("conf.level", logConf.Conf.Level)

	logger, _, _  := GetLogger()
	logger.Info("test log")
}
