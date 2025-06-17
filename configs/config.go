package configs

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

var (
	App *ConfigInfo
)

func Setup(env string) {
	// Initialize the configuration
	file := fmt.Sprintf("configs/config.%s.yml", env)
	if _, err := os.Stat(file); err != nil {
		if os.IsNotExist(err) {
			color.Red("** config file not exist **")
			return
		}
	}
	v := viper.New()
	v.SetConfigFile(file)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	config := &ConfigInfo{}
	if err := v.Unmarshal(&config); err != nil {
		panic(err)
	}
	color.Blue("config init", config)
	App = config
}

type ConfigInfo struct {
	Name         string       `mapstructure:"name" json:"name"`
	Host         string       `mapstructure:"host" json:"host"`
	Port         int          `mapstructure:"port" json:"port"`
	DatabaseInfo DatabaseInfo `mapstructure:"database" json:"database"`
	MysqlInfo    MysqlInfo    `mapstructure:"mysql" json:"mysql"`
	RedisInfo    RedisInfo    `mapstructure:"redis" json:"redis"`
	EtcdInfo     EtcdInfo     `mapstructure:"etcd" json:"etcd"`
	MongoInfo    MongoInfo    `mapstructure:"mongo" json:"mongo"`
	LoggerInfo   LoggerInfo   `mapstructure:"logger" json:"logger"`
	JwtInfo      JwtInfo      `mapstructure:"jwt" json:"jwt"`
	JobInfo      JobInfo      `mapstructure:"job" json:"job"`
	MqInfo       MqInfo       `mapstructure:"mq" json:"mq"`
	TracerInfo   TracerInfo   `mapstructure:"tracer" json:"tracer"`
}

type DatabaseInfo struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

type TracerInfo struct {
	Enable bool `mapstructure:"enable"`
}

type MqInfo struct {
	Enable bool   `mapstructure:"enable"`
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Pass   string `mapstructure:"pass"`
	VHost  string `mapstructure:"vhost"`
}

type MysqlInfo struct {
	Enable   bool   `mapstructure:"enable"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	Charset  string `mapstructure:"charset"`
}

type RedisInfo struct {
	Enable     bool   `mapstructure:"enable"`
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Password   string `mapstructure:"password"`
	Database   int    `mapstructure:"database"`
	Expiration int    `mapstructure:"expiration"`
}

type EtcdInfo struct {
	Enable    bool     `mapstructure:"enable"`
	Endpoints []string `mapstructure:"endpoints"`
	Timeout   int      `mapstructure:"timeout"`
	Username  string   `mapstructure:"username"`
	Password  string   `mapstructure:"password"`
}

type MongoInfo struct {
	Enable   bool   `mapstructure:"enable"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"database"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type LoggerInfo struct {
	Enable bool   `mapstructure:"enable"`
	Path   string `mapstructure:"path"`
	Level  string `mapstructure:"level"`
	File   string `mapstructure:"filename"`
}

type JobInfo struct {
	Enable  bool `mapstructure:"enable"`
	Timeout int  `mapstructure:"timeout"`
}

type JwtInfo struct {
	Secret string `mapstructure:"secret"`
	Expire int    `mapstructure:"expire"`
}

type ElasticInfo struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type MinioInfo struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"access_key_id"`
	SecretAccessKey string `yaml:"secret_access_key"`
	BucketName      string `yaml:"bucket_name"`
	UseSSL          bool   `yaml:"use_ssl"`
}

type UploadInfo struct {
	UploadPath string `yaml:"upload_path"`
	UploadUrl  string `yaml:"upload_url"`
}
