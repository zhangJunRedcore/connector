package controller

import (
	"bytes"
	"connector/conf"
	"connector/lib"
	"encoding/gob"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

var gatewayJSON []byte
var infoLog lib.InfoLogger
var debugLog lib.DebugLogger
var errorLog lib.ErrorLogger

type Data struct {
	ports       Ports
	manager     string `json:"manager"`
	hostAddress string `json:"hostAddress"`
}

type Ports struct {
	HTTP  int `json:"http"`
	HTTPS int `json:"https"`
}

type GatewayJSON struct {
	Body      io.Reader
	CompanyID string `json:"CompanyID"`
	ErrorCode string `json:"errCode"`
}

//GenerateGatewayJSON work for xx.json
func GenerateGatewayJSON(c *gin.Context) {

	var body GatewayJSON
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "INVALID_PARAMS",
		})
	}
	jsonBody, _ := getBytes(body)

	debugLog.Println("GenerateGatewayJSON body is:", body)
	debugLog.Println(`byteBody is :`, jsonBody)
	debugLog.Println(`CompanyID is: `, body.CompanyID)

	// writeGatewayData(jsonBody, body.CompanyID)

	file := conf.GatewayCompanyConfigDir + body.CompanyID + `_gateway.json`
	debugLog.Println(`[gateway] write gateway file:`, file)

	err := ioutil.WriteFile(file, jsonBody, 0777)
	if err != nil {
		errorLog.Println(`写入文件失败:`, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "",
	})
}

func InitGateway(c *gin.Context) {
	managerAddr := conf.ManagerAddr
	connectorServer := conf.ConnectorServer
	url := managerAddr + `/manager/api/gateway?host=` + connectorServer
	debugLog.Println("GET manager url is:", url)

	response, err := http.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	debugLog.Println("response get body is:", reader)
	// gatewayGroup := groupEleFromGatewayFiles(reader)
	// modifyBgxConf(gatewayGroup.ports)
}

func getBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

/*
func GenerateGatewayData(data []byte, CompanyID string) interface{} {

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
		return nil
	}
	return response.Body
}

// 从不同 gateway 中提取信息，并归类
func groupEleFromGatewayFiles(gateways io.ReadCloser) *Data {

	data := &Data{}
	data.ports.HTTP = 80
	data.ports.HTTPS = 443

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

func praseAddressPorts(i interface{}, data *Data) {
	jsondata, _ := json.Marshal(data)
	debugLog.Println("port json data is :", jsondata)
	// data.manager& data.hostAddress  合并数组  for range
	json.
}
*/
