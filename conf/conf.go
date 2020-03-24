// @Title conf
// @Description conf
// @Author jun.zhang@clouddeep.cn
// @Update 2019.11.27
package conf

import (
	"connector/lib"
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var KmsAddress string
var KmsProtocol string
var KmsTimeout time.Duration
var ServerTimeout time.Duration
var SpaAccessport string
var TimeDifference int
var AccessIsVerify string
var AccessDstPort string
var Type string
var IsPolicyVerified string
var IPMask string
var GatewayCompanyConfigDir string

// Config struct
type Config struct {
	Name string
}

// Init config
// @Param config_path
// @Return err
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	//c.WatchConfig()
	return nil
}

// Init config private
// @Param
// @Return err
func (c *Config) initConfig() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	AccessIsVerify = viper.GetString("access.isVerify")
	AccessDstPort = viper.GetString("access.dstport")
	GatewayCompanyConfigDir = viper.GetString("path.jsonpath")
	Type = viper.GetString("type")
	IsPolicyVerified = viper.GetString("server.is_policy_verified")
	IPMask = viper.GetString("server.ip_mask")
	watchConfig()
	return nil
}

// WatchConfig hot modify
// 监听配置文件是否改变,用于热更新
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
		AccessIsVerify = viper.GetString("access.isVerify")
		AccessDstPort = viper.GetString("access.dstport")
		GatewayCompanyConfigDir = viper.GetString("path.jsonpath")
		Type = viper.GetString("type")
		IsPolicyVerified = viper.GetString("server.is_policy_verified")
		lib.ModifyLevel()
	})
}
