package services

import (
	"connector/conf"
	"connector/lib"
	"net/http"
	"github.com/spf13/viper"
)

var infoLog lib.InfoLogger
var debugLog lib.DebugLogger
var errorLog lib.ErrorLogger

type GatewayData struct {
	Data      map[string]interface{} `json:"data"`
	CompanyId string                 `json:"companyId"`
}

//InitGateway InitGateway
func InitGateway() {
	managerAddr := conf.ManagerAddr
	url := managerAddr + `/manager/api/gateway?companyId=` + viper.GetString("server.company_id")
	debugLog.Println("GET manager url is:", url)

	response, err := http.Get(url)
	debugLog.Println("/manager/api/gateway response", response)
	if err != nil || response.StatusCode != http.StatusOK {
		errorLog.Println("get /manager/api/gateway result err !", url, err)
		return
	}

	reader := response.Body
	debugLog.Println("response get body is:", reader)
}
