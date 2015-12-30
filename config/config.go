package config

import (
	"io/ioutil"

	"pirate-sharding/core/yaml"
)

//整个config文件对应的结构
type Config struct {
	Addr        string       `yaml:"addr"`
	User        string       `yaml:"user"`
	Password    string       `yaml:"password"`
	LogPath     string       `yaml:"log_path"`
	LogLevel    string       `yaml:"log_level"`
	LogSql      string       `yaml:"log_sql"`
	SlowLogTime int          `yaml:"slow_log_time"`
	AllowIps    string       `yaml:"allow_ips"`
	Nodes       []NodeConfig `yaml:"nodes"`

	Schema SchemaConfig `yaml:"schema"`
}

//node节点对应的配置
type NodeConfig struct {
	Name             string `yaml:"name"`
	DownAfterNoAlive int    `yaml:"down_after_noalive"`
	MaxConnNum       int    `yaml:"max_conns_limit"`

	User     string `yaml:"user"`
	Password string `yaml:"password"`

	Master string `yaml:"master"`
	Slave  string `yaml:"slave"`
}

//schema对应的结构体
type SchemaConfig struct {
	DB        string        `yaml:"db"`
	Nodes     []string      `yaml:"nodes"`
	Default   string        `yaml:"default"` //默认路由规则
	ShardRule []ShardConfig `yaml:"shard"`   //range或hash路由规则
}

//range或hash路由规则
type ShardConfig struct {
	Table         string   `yaml:"table"`
	Key           string   `yaml:"key"`
	Nodes         []string `yaml:"nodes"`
	Locations     []int    `yaml:"locations"`
	Type          string   `yaml:"type"`
	TableRowLimit int      `yaml:"table_row_limit"`
}

func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return ParseConfigData(data)
}
