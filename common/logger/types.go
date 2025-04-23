package logger

import "github.com/rs/zerolog"

// Config 日志配置
type Config struct {
	ConsoleOutput bool   `mapstructure:"console_output"` // 是否输出到控制台
	FileOutput    bool   `mapstructure:"file_output"`    // 是否输出到文件
	FilePath      string `mapstructure:"file_path"`      // 日志文件路径
	Level         string `mapstructure:"level"`          // 日志级别
	WithCaller    bool   `mapstructure:"with_caller"`    // 是否包含调用者信息
	EchoHeader    string `mapstructure:"echo_header"`    // Echo日志头部
	EchoPrefix    string `mapstructure:"echo_prefix"`    // Echo日志前缀
}

// Check 日志配置检查
func (cfg *Config) Check() {
	// 为空去默认值
	if cfg.FilePath == "" {
		cfg.FilePath = "logs/app.log"
	}
	if cfg.Level == "" {
		cfg.Level = zerolog.InfoLevel.String()
	}
}
