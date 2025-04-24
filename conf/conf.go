package conf

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/hildam/AI-Cloud-Drive/common/logger"
	"github.com/spf13/viper"
)

var (
	globalCfg *AppConfig // 应用配置
)

// Init 初始化配置
func Init() error {
	// 加载配置
	cfg, err := loadConfig()
	if err != nil {
		return fmt.Errorf("Init config failed, load config err: %v", err)
	}
	globalCfg = cfg

	// 初始化日志
	if err := logger.Init(cfg.Log); err != nil {
		return fmt.Errorf("Init logger failed, err: %+v", err)
	}
	return nil
}

// loadConfig 加载配置
func loadConfig() (*AppConfig, error) {
	cfg := &AppConfig{}

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(".")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		log.Printf("Init config failed, read config err: %v", err)
		return nil, fmt.Errorf("Init config failed, read config err: %v", err)
	}

	// 监听配置变化
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(cfg); err != nil {
			log.Printf("Init config failed, unmarshal config err: %v", err)
		}
	})

	// 解析配置
	if err := v.Unmarshal(cfg); err != nil {
		log.Printf("Init config failed, unmarshal config err: %v", err)
		return nil, fmt.Errorf("Init config failed, unmarshal config err: %v", err)
	}
	return cfg, nil
}

// GetCfg 获取配置
func GetCfg() *AppConfig {
	return globalCfg
}
