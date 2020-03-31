package controller

import (
	"connector/lib"
	"connector/services"
	"connector/utils"
	"connector/conf"
	"net/http"
	"github.com/spf13/viper"
	"github.com/gin-gonic/gin"
)

var infoLog lib.InfoLogger
var debugLog lib.DebugLogger
var errorLog lib.ErrorLogger

//GenerateGatewayJSON work for xx.json
func GenerateGatewayJSON(ctx *gin.Context) {
	app := services.Gin{Ctx: ctx}
	var body services.GatewayData
	if ctx.BindJSON(&body) != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"msg": "INVALID_PARAMS",
		})
	}
	debugLog.Println("GenerateGatewayJSON body is:", body)

	dataString, err := utils.Map2String(body.Data)
	if err != nil {	
		errorLog.Println(err)
		app.Response(200, err.Error(),nil)
		return 
	}

	err = utils.WriteFile(body.CompanyId , []byte(dataString))
	if err != nil {
		errorLog.Println(err)
		app.Response(200, err.Error(),nil)
		return
	}

	confPath := viper.ConfigFileUsed()
	yaml := &utils.Yaml{YamlPath: confPath}
	// modify companyId
	err = yaml.Modify("company_id", body.CompanyId)
	if err != nil {
		errorLog.Println("modify companyId err !", err)
		app.Response(200, err.Error(),nil)
		return
	}
	// modify gatewayId
	if body.Data["gatewayId"] == nil {
		errorLog.Println("gatewayId is nil")
		app.Response(200, "gatewayId is nil", nil)
		return
	}

	err = yaml.Modify("gateway_id", body.Data["gatewayId"].(string))
	if err != nil {
		errorLog.Println("modify gateway err !", err)
		app.Response(200, err.Error(),nil)
		return
	}

	url := conf.NgxSharedSetUrl + "/" + body.CompanyId
	response, err := services.Get(url)
	if err != nil || response.StatusCode != http.StatusOK {
		errorLog.Println("get clouddeep/shared/set result err !", err)
		app.Response(200, err.Error(),nil)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
