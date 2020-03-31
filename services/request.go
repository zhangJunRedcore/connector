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
	Data  		map[string]interface{} `json:"data"`
	CompanyId   string                   `json:"companyId"`
}

type GatewayDataList struct {
	DataList  []map[string]interface{} `json:"data"`
	ErrCode   string                   `json:"errCode"`
}

//InitGateway InitGateway
func InitGateway() {
	managerAddr := conf.ManagerAddr
	url := managerAddr + `/manager/api/gateway?companyId=` + viper.GetString("server.company_id") + "&gatewayId=" + viper.GetString("server.gateway_id")
	debugLog.Println("GET manager url is:", url)

	response, err := http.Get(url)
	debugLog.Println("/manager/api/gateway response", response)
	if err != nil || response.StatusCode != http.StatusOK {
		errorLog.Println("get manager err !", url, err)
		return
	}

	//read response body
	bytebody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("change response.Body type to []byte err, %v\n", err)
		return
	}

	//get json body into data
	var data GatewayDataList
	if err = json.Unmarshal(bytebody, &data); err != nil {
		errorLog.Println("Unmarshal err:", err)
		return
	}

	if data.ErrCode != "0" || len(data.DataList) == 0 {
		errorLog.Println("can not get data from manager, errcode:", data.ErrCode)
		return
	}

	companyId := data.DataList[0]["companyId"]
	if companyId == nil {
		errorLog.Println("can not get companyId from data, companyId:", companyId)
		return
	}
	debugLog.Println("data.Data is:", data.DataList[0])
	debugLog.Println("data.errCode is:", data.ErrCode)

	dataString, err := utils.Map2String(data.DataList[0])
	if err != nil {
		errorLog.Println(err)
		return
	}

	err = utils.WriteFile(companyId.(string), []byte(dataString))
	if err != nil {
		errorLog.Println(err)
		return
	}
}
