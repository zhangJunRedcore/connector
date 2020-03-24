package controller

import (
	"bytes"
	"connector/conf"
	"encoding/gob"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

var gatewayJSON []byte

type GatewayJSON struct {
	Body      io.Reader
	CompanyId string `json:"companyId"`
}

//GenerateGatewayJSON work for xx.json
func GenerateGatewayJSON(c *gin.Context) {

	var body GatewayJSON
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "INVALID_PARAMS",
		})
	}
	jsonBody, _ := GetBytes(body)

	log.Println("GenerateGatewayJSON body is:", body)
	log.Println(`jsonBody is :`, jsonBody)
	log.Println(`companyID is: `, body.CompanyId)

	writeGatewayData(jsonBody, body.CompanyId)
	version := getGatewayVersion()

	c.JSON(http.StatusOK, gin.H{
		"version": version,
	})
}

//generateGatewayData write json in file
func writeGatewayData(jsonBody []byte, companyID string) {
	file := conf.GatewayCompanyConfigDir + companyID + `_gateway.json`

	err := ioutil.WriteFile(file, jsonBody, 0777)
	if err != nil {
		log.Println(`写入文件失败:`, err)
	}
}

//generateGatewayData get gateway version
func getGatewayVersion() string {
	return "x.x.x"
}

func GetGatewayJSON(c *gin.Context) {
	companyID := c.Param("companyId")
	files := readFileByPath(conf.GatewayCompanyConfigDir)

	for _, file := range files {
		log.Println("[file]", file, companyID)
		result, _ := regexp.Match(companyID, []byte(string(file)))
		if result == true {
			gatewayJSONFilePath := conf.GatewayCompanyConfigDir + string(file)
			gatewayJSON = readFileByPath(gatewayJSONFilePath)
			log.Println(`[gateway]`, gatewayJSON)
			break
		}
	}
	if gatewayJSON == nil {
		c.JSON(http.StatusOK, gin.H{
			"massage": "gateway json not found",
		})
	}
	url := `http://127.0.0.1/clouddeep/shared/get/` + companyID
	gatewayShared, _ := c.Get(url)
	c.JSON(http.StatusOK, gin.H{
		"gatewayFile":   gatewayJSON,
		"gatewayShared": gatewayShared,
	})
}

func readFileByPath(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("读取文件失败:%s", err)
	}
	return content
}

func GetBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
