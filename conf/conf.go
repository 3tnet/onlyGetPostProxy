package conf

import (
	"encoding/json"
	"os"
	"log"
)

type Config struct {
	ListenPort string `json:"listen_port"`
	ServerHost string `json:"server_host"`
	MethodKey  string `json:"method_key"`
	Scheme     string `json:"scheme"`
	fileName   string
}
func NewConfig() *Config{
	// 默认配置
	return &Config{ "9090",
					"127.0.0.1:8080",
					"_method",
					"http",
					"./onlyGetPostProxyConf.json" }
}
func (c *Config) SetFileName(fileName string){
	c.fileName = fileName
}
func (c *Config) ReadConfig(){

	if _, err := os.Stat(c.fileName); !os.IsNotExist(err) {
		file, err := os.Open(c.fileName)
		if err != nil {
			log.Fatalf("打开配置文件 %s 出错:%s", c.fileName, err)
		}
		defer file.Close()

		err = json.NewDecoder(file).Decode(c)

		if err != nil {
			log.Fatalf("格式有误的 JSON 配置文件:\n%s", file)
		}
	}
}
