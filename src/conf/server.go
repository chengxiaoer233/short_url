/**
* @Description：
* @Author: cdx
* @Date: 2022/11/24 2:22 下午
 */

package conf

type RedisConf struct {
	Host   string `toml:"host"`
	Port   string `toml:"port"`
	PassWd string `toml:"passWd"`
	Db     int    `toml:"db"`
}

type MysqlConf struct {
	Db         string `toml:"db"`
	DbHost     string `toml:"dbHost"`
	DbPort     string `toml:"dbPort"`
	DbUser     string `toml:"dbUser"`
	DbPassWord string `toml:"dbPassWord"`
}

type LogStruct struct {
	LogPath    string `json:"logPath"`
	MaxAge     int    `json:"maxAge"`
	RotateTime int    `json:"rotateTime"`
	Level      int    `json:"level"`
}
