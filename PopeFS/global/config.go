package global

import (
	"git.xiaojukeji.com/falcon/pope-fs/util/confutil"
)

const (
	//ConfigBasePath 基础配置文件config
	ConfigBasePath = "conf"
	//MainConfigFile 主配置文件路径
	MainConfigFile = "pope-fs.conf"
	//DataBaseConfigFile 数据库配置文件路径
	DataBaseConfigFile = "database.json"
	//LoggerConfigFile 日志配置文件路径
	LoggerConfigFile = "logger.conf"
	//NodeManagerConfigFile node manager 配置文件路径
	NodeManagerConfigFile = "nodemgr.json"
	//RockstableConfigFile rt 配置文件路径
	RockstableConfigFile = "rockstable.json"
	//CodisConfigFile codis 集群配置文件
	CodisConfigFile = "codis.json"
	//RateLimitConfigFile 文件路径
	RateLimitConfigFile = "limiter.json"
	//ApolloConfigFile Apollo配置文件路径
	ApolloConfigFile = "apollo.json"

	// FeatureConfigFile 特征配置文件，未来迁移到 db
	FeatureConfigFile = "features.json"
	// FeatureRelatedServiceConfigFile 服务配置文件，未来迁移到 db
	FeatureRelatedServiceConfigFile = "services.json"

	// CouponMetaConfigFile coupon meta data
	CouponMetaConfigFile = "coupon.json"

	// ProductMeteConfigFile product meta data
	ProductMetaConfigFile = "product.json"

	//S3BucketMetaConfigFile = "s3Bucket.json"

	// ProductMeteConfigFile product meta data
	CouponGroupMetaConfigFile = "couponGroup.json"

	// FeatureOfsConfigFile ofs config
	FeatureOfsConfigFile = "ofs.json"

	// DISF config file
	DisfTomlFile = "disf.toml"

	// DISF yaml file
	DisfYamlFile = "disf.yaml"
)

//Config 全局变量
var Config Holder

//Holder 存储配置项
type Holder struct {
	Nereus NereusConfig `toml:"pope-fs"`
}

//NereusConfig 配置项
type NereusConfig struct {
	Mode       string // 服务模式，和配置文件相关
	Port       string // http 服务端口号
	UtcEnable  int    // utc多时区支持
	ThriftPort string // thrift 服务端口号
}

//APIConfig API相关配置文件
type APIConfig struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

//ServiceConfig 服务配置文件
type ServiceConfig struct {
	Name                string      `json:"name"`
	Server              []string    `json:"server"`
	API                 []APIConfig `json:"api"`
	ConnTimeoutMs       int         `json:"conntimeout"`
	URL                 string      `json:"url"`
	RspTimeoutMs        int         `json:"rsptimeout"`
	RetryTime           int         `json:"retrytime"`
	RetryCount          int         `json:"retrycount"`
	MaxIdleConnsPerHost int         `json:"maxidleconnsperhost"`
}

// GetMainFile 获取基本配置文件路径
func GetMainFile() string {
	return GetConfFile(MainConfigFile)
}

// 获取disf.yaml
func GetDisfYaml() string {
	return confutil.AppendPaths(ConfigBasePath, DisfYamlFile)
}

// GetConfFile 获取全局配置文件路径
func GetConfFile(filename string) string {
	return confutil.AppendPaths(ConfigBasePath, filename)
}

// GetModeFile 根据配置的执行模式，获取对应模式下的配置文件全路径
func GetModeFile(configFile string) string {
	return confutil.AppendPaths(ConfigBasePath, Config.Nereus.Mode, configFile)
}
