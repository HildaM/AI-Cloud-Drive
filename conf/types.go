package conf

import (
	"github.com/hildam/AI-Cloud-Drive/common/logger"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

// 存储类型枚举
const (
	LocalStorageType = "local" // 本地文件存储
	OssStorageType   = "oss"   // 阿里云oss
	MinioStorageType = "minio" // minio存储
)

// AppConfig 应用配置
type AppConfig struct {
	Server    ServerConfig    `mapstructure:"server"`    // 服务器配置
	Database  *DatabaseConfig `mapstructure:"database"`  // 数据库配置
	JWT       JWTConfig       `mapstructure:"jwt"`       // JWT配置
	Storage   *StorageConfig  `mapstructure:"storage"`   // 存储配置
	CORS      CORSConfig      `mapstructure:"cors"`      // CORS配置
	RAG       RAGConfig       `mapstructure:"rag"`       // RAG配置
	Embedding EmbeddingConfig `mapstructure:"embedding"` // 嵌入模型配置
	LLM       LLMConfig       `mapstructure:"llm"`       // 语言模型配置
	Log       logger.Config   `mapstructure:"log"`       // 日志配置
	Milvus    *MilvusConfig   `mapstructure:"milvus"`
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
	// 通用配置
	Type          string `mapstructure:"type"`            // local/oss/minio
	UrlExpireTime int64  `mapstructure:"url_expire_time"` // 临时 url过期时间

	// 存储配置
	Local *LocalConfig `mapstructure:"local"` // 本地存储配置
	OSS   *OSSConfig   `mapstructure:"oss"`   // OSS配置
	Minio *MinioConfig `mapstructure:"minio"` // Minio配置
}

// CheckCfg 检查参数
func (s *StorageConfig) CheckCfg() {
	if s.Type == "" {
		s.Type = LocalStorageType
	}
	if s.UrlExpireTime <= 0 {
		s.UrlExpireTime = 6000 // 单位秒
	}
}

// MilvusConfig Milvus向量数据库配置
type MilvusConfig struct {
	Address         string `mapstructure:"address"`
	CollectionName  string `mapstructure:"collection_name"`
	VectorDimension int    `mapstructure:"vector_dimension"`
	IndexType       string `mapstructure:"index_type"`
	MetricType      string `mapstructure:"metric_type"`
	Nlist           int    `mapstructure:"nlist"`
	// 搜索参数
	Nprobe int `mapstructure:"nprobe"`
	// 字段最大长度配置
	IDMaxLength      string `mapstructure:"id_max_length"`
	ContentMaxLength string `mapstructure:"content_max_length"`
	DocIDMaxLength   string `mapstructure:"doc_id_max_length"`
	DocNameMaxLength string `mapstructure:"doc_name_max_length"`
	KbIDMaxLength    string `mapstructure:"kb_id_max_length"`
}

// GetMetricType 获取类型
func (m *MilvusConfig) GetMetricType() entity.MetricType {
	// 获取配置的度量类型
	var metricType entity.MetricType
	switch m.MetricType {
	case "L2":
		metricType = entity.L2 // 欧几里得距离：测量向量间的直线距离，适合图像特征等数值型向量
	case "IP":
		metricType = entity.IP // 内积距离：适合已归一化的向量，计算效率高
	default:
		metricType = entity.COSINE // 余弦相似度：测量向量方向的相似性，适合文本语义搜索
	}
	return metricType
}

// GetMilvusIndex 根据配置构建索引
func (m *MilvusConfig) GetMilvusIndex() (entity.Index, error) {
	// 选择索引类型的距离度量方式
	metricType := m.GetMetricType()

	// 创建索引
	var (
		idx entity.Index
		err error
	)
	if m.Nlist <= 0 {
		m.Nlist = 128 // 为空，取默认值
	}

	switch m.IndexType {
	case "IVF_FLAT":
		// IVF_FLAT: 倒排文件索引 + 原始向量存储
		// 优点：搜索精度高；缺点：内存占用较大
		// nlist: 聚类数量，值越大精度越高但速度越慢，通常设置为 sqrt(n) 到 4*sqrt(n)，其中n为向量数量
		idx, err = entity.NewIndexIvfFlat(metricType, m.Nlist)
	case "IVF_SQ8":
		// IVF_SQ8: 倒排文件索引 + 标量量化压缩存储（8位）
		// 优点：比IVF_FLAT节省内存；缺点：轻微精度损失
		// nlist: 与IVF_FLAT相同，根据数据规模调整
		idx, err = entity.NewIndexIvfSQ8(metricType, m.Nlist)
	case "HNSW":
		// HNSW: 层次可导航小世界图索引，高效且精确但内存占用大
		// M: 每个节点的最大边数，影响图的连通性和构建/查询性能
		//    - 值越大，构建越慢，内存占用越大，但查询越精确
		//    - 通常取值范围为8-64之间，默认值8在大多数场景下平衡了性能和精度
		// efConstruction: 构建索引时每层搜索的候选邻居数量
		//    - 值越大，构建越慢，索引质量越高
		//    - 通常取值范围为40-800，默认值40在大多数场景下表现良好
		// 注：这两个参数需要根据数据特性和性能要求综合调优，目前使用经验值
		idx, err = entity.NewIndexHNSW(metricType, 8, 40) // M=8, efConstruction=40
	default:
		// 默认使用IVF_FLAT，兼顾搜索精度和性能
		idx, err = entity.NewIndexIvfFlat(metricType, m.Nlist)
	}
	return idx, err
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
