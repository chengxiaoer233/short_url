/**
* @Description：
* @Author: cdx
* @Date: 2022/11/24 2:21 下午
 */

package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
)

type Config struct {
	RedisConf RedisConf `toml:"redis_conf"`
	MysqlConf MysqlConf `toml:"database"`
	LogStruct LogStruct `toml:"log"`
}

var config = &Config{}

func GetConfig() *Config {
	return config
}

func init() {

	appConfigPath := ""
	flag.StringVar(&appConfigPath, "c", "./etc/config-production.toml", "config path")
	flag.Parse()

	toml.DecodeFile(appConfigPath, config)
}
