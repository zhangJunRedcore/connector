package main

import (
	"connector/conf"
	"connector/lib"
	"connector/router"
	"connector/services"
	"flag"
	"log"

	"github.com/spf13/viper"
)

var configPath = flag.String("config_path", "/project/connector/etc/conf.yaml", "config file path")
//var configPath = flag.String("config_path", "./conf/conf.yaml", "config file path")

func init() {
	log.Println("config_path:", *configPath)
	err := conf.Init(*configPath)
	if err != nil {
		panic(err)
	}
}

func main() {

	// get log level (default: info)
	logLevel := func() int {
		switch viper.GetString("log.level") {
		case "debug":
			return lib.LogLevelDebug
		case "info":
			return lib.LogLevelInfo
		case "error":
			return lib.LogLevelError
		case "silent":
			return lib.LogLevelSilent
		}
		return lib.LogLevelInfo
	}()

	logger := lib.NewLogger(
		logLevel,
	)
	logger.Info.Println("Starting Logger")
	logger.Debug.Println("Debug log enabled")

	services.InitGateway()

	r := router.InitRouter()
	host := ":" + viper.GetString("server.http_port")
	logger.Info.Println(host)

	r.Run(host)
}
