package controller

import (
	"connector/lib"
	"connector/service"
	"connector/tools"

	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var gatewayJSON []byte
var infoLog lib.InfoLogger
var debugLog lib.DebugLogger
var errorLog lib.ErrorLogger

type Data struct {
	Ports       Ports
	Manager     string `json:"manager"`
	HostAddress string `json:"hostAddress"`
}

type Ports struct {
	HTTP  int `json:"http"`
	HTTPS int `json:"https"`
}

type GatewayJSON struct {
	Body      io.Reader
	CompanyID string `json:"companyID"`
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
	jsonBody, _ := tools.GetBytes(body)

	debugLog.Println("GenerateGatewayJSON body is:", body)
	debugLog.Println(`byteBody is :`, jsonBody)
	debugLog.Println(`CompanyID is: `, body.CompanyID)

	service.GenerateGatewayData(jsonBody, body.CompanyID)

	c.JSON(http.StatusOK, gin.H{
		"msg": "",
	})
}
