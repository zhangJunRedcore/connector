package main

import (
	"connector/conf"
	"connector/lib"
	"connector/router"
	"connector/service"
	"flag"
	"fmt"
	"log"

	"github.com/spf13/viper"
)

var configPath = flag.String("config_path", "/project/deepctl/etc/conf.yaml", "config file path")

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

	var interfaceName string = "interfaceName"

	logger := lib.NewLogger(
		logLevel,
		fmt.Sprintf("(%s) ", interfaceName),
	)
	logger.Info.Println("Starting Logger")
	logger.Debug.Println("Debug log enabled")

	service.InitGateway()

	r := router.InitRouter()

	r.Run(":8888")

}
