package conf

import "github.com/hildam/AI-Cloud-Drive/common/logger"

// AppConfig 应用配置
type AppConfig struct {
	Server    ServerConfig    `mapstructure:"server"`    // 服务器配置
	Database  DatabaseConfig  `mapstructure:"database"`  // 数据库配置
	JWT       JWTConfig       `mapstructure:"jwt"`       // JWT配置
	Storage   StorageConfig   `mapstructure:"storage"`   // 存储配置
	CORS      CORSConfig      `mapstructure:"cors"`      // CORS配置
	RAG       RAGConfig       `mapstructure:"rag"`       // RAG配置
	Embedding EmbeddingConfig `mapstructure:"embedding"` // 嵌入模型配置
	LLM       LLMConfig       `mapstructure:"llm"`       // 语言模型配置
	Log       logger.Config   `mapstructure:"log"`       // 日志配置
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `mapstructure:"port"` // 端口
	Env  string `mapstructure:"env"`  // 环境
}

// GetPort 获取端口
func (s ServerConfig) GetPort() string {
	if s.Port == "" {
		return ":4090"
	}
	return ":" + s.Port
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`     // 主机
	Port     string `mapstructure:"port"`     // 端口
	User     string `mapstructure:"user"`     // 用户名
	Password string `mapstructure:"password"` // 密码
	Name     string `mapstructure:"name"`     // 数据库名
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret          string `mapstructure:"secret"`           // 密钥
	ExpirationHours int    `mapstructure:"expiration_hours"` // 过期时间（小时）
}

// MinioConfig Minio配置
type MinioConfig struct {
	Endpoint        string `mapstructure:"endpoint"`          // 端点
	Bucket          string `mapstructure:"bucket"`            // 桶
	AccessKeyID     string `mapstructure:"access_key_id"`     // 访问密钥ID
	AccessKeySecret string `mapstructure:"access_key_secret"` // 访问密钥密钥
	UseSSL          bool   `mapstructure:"use_ssl"`           // 是否使用SSL
	Region          string `mapstructure:"region"`            // 区域
}

// StorageConfig 存储配置
type StorageConfig struct {
	Type  string      `mapstructure:"type"`  // local/oss/minio
	Local LocalConfig `mapstructure:"local"` // 本地存储配置
	OSS   OSSConfig   `mapstructure:"oss"`   // OSS配置
	Minio MinioConfig `mapstructure:"minio"` // Minio配置
}

// LocalConfig 本地存储配置
type LocalConfig struct {
	BaseDir string `mapstructure:"base_dir"` // 本地存储根目录（如 /data/storage）
}

// OSSConfig OSS配置
type OSSConfig struct {
	Endpoint        string `mapstructure:"endpoint"`          // 端点
	Bucket          string `mapstructure:"bucket"`            // 桶
	AccessKeyID     string `mapstructure:"access_key_id"`     // 访问密钥ID
	AccessKeySecret string `mapstructure:"access_key_secret"` // 访问密钥密钥
}

// CORSConfig CORS配置
type CORSConfig struct {
	AllowOrigins     []string `mapstructure:"allow_origins"`     // 允许的源
	AllowMethods     []string `mapstructure:"allow_methods"`     // 允许的方法
	AllowHeaders     []string `mapstructure:"allow_headers"`     // 允许的头
	ExposeHeaders    []string `mapstructure:"expose_headers"`    // 暴露的头
	AllowCredentials bool     `mapstructure:"allow_credentials"` // 允许凭证
	MaxAge           string   `mapstructure:"max_age"`           // 使用字符串表示时间，便于配置
}

// RAGConfig RAG配置
type RAGConfig struct {
	ChunkSize   int `mapstructure:"chunk_size"`   // 块大小
	OverlapSize int `mapstructure:"overlap_size"` // 重叠大小
}

// EmbeddingConfig 嵌入模型配置
type EmbeddingConfig struct {
	// 使用哪种嵌入模型: remote 或 ollama
	Service string                `mapstructure:"service"` // 选择的服务
	Remote  RemoteEmbeddingConfig `mapstructure:"remote"`  // 远程嵌入模型配置
	Ollama  OllamaEmbeddingConfig `mapstructure:"ollama"`  // Ollama嵌入模型配置
}

// RemoteEmbeddingConfig 远程嵌入模型配置
type RemoteEmbeddingConfig struct {
	APIKey  string `mapstructure:"api_key"`  // API密钥
	Model   string `mapstructure:"model"`    // 模型
	BaseURL string `mapstructure:"base_url"` // 基础URL
}

// OllamaEmbeddingConfig Ollama嵌入模型配置
type OllamaEmbeddingConfig struct {
	URL   string `mapstructure:"url"`   // URL
	Model string `mapstructure:"model"` // 模型
}

// LLMConfig 语言模型配置
type LLMConfig struct {
	APIKey      string  `mapstructure:"api_key"`     // API密钥
	Model       string  `mapstructure:"model"`       // 模型
	BaseURL     string  `mapstructure:"base_url"`    // 基础URL
	MaxTokens   int     `mapstructure:"max_tokens"`  // 最大令牌数
	Temperature float32 `mapstructure:"temperature"` // 温度
}
