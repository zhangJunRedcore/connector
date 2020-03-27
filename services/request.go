package services

import (
	"connector/conf"
	"connector/lib"
	"connector/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	//change body type to []byte
	bytebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	debugLog.Println("json:", string(bytebody))

	//get json body into data
	var data GatewayData
	if err = json.Unmarshal(bytebody, &data); err != nil {
		errorLog.Println("Unmarshal err, %v\n", err)
		return
	}

	debugLog.Println("body is:", data)
	debugLog.Println("body.companyId is:", data.CompanyId)

	dataString, err := utils.Map2String(data.Data)
	if err != nil {
		errorLog.Println(err)
		return
	}

	err = utils.WriteFile(data.CompanyId, []byte(dataString))
	if err != nil {
		errorLog.Println(err)
		return
	}

}
