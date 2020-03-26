package service

import (
	"connector/conf"
	"connector/lib"
	"io/ioutil"
	"net/http"
)

var infoLog lib.InfoLogger
var debugLog lib.DebugLogger
var errorLog lib.ErrorLogger

//InitGateway InitGateway
func InitGateway() {
	managerAddr := conf.ManagerAddr
	connectorServer := conf.ConnectorServer
	url := managerAddr + `/manager/api/gateway?host=` + connectorServer
	debugLog.Println("GET manager url is:", url)

	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		return
	}

	reader := response.Body
	debugLog.Println("response get body is:", reader)
}

//GenerateGatewayData job is set nginx
func GenerateGatewayData(data []byte, CompanyID string) {

	file := conf.GatewayCompanyConfigDir + CompanyID + `_gateway.json`
	debugLog.Println(`[gateway] write gateway file:`, file)

	err := ioutil.WriteFile(file, data, 0777)
	if err != nil {
		errorLog.Println(`写入文件失败:`, err)
	}

	url := conf.NgxSharedSetUrl + "/" + CompanyID
	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		errorLog.Println("get url result err !")
	}
}

// 从不同 gateway 中提取信息，并归类
/*
func groupEleFromGatewayFiles(gateways io.ReadCloser) *Data {

	data := &Data{}
	data.ports.http = 80
	data.ports.https = 443

	for _, gateway := range gateways {
		gatewayJSONFilePath := conf.GatewayCompanyConfigDir + `conf_gateway` + gateway.CompanyID + `_gateway.json`
		err := ioutil.WriteFile(gatewayJSONFilePath, gateway, 0777)
		if err != nil {
			errorLog.Println(`写入文件失败:`, err)
		}
		praseAddressPorts(gateway, data)
	}
	return data
}
*/
