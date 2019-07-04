package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"swagger-mocker/swagger_mocker"
)

const JSON_URL = "http://8789.office.qunjielong.com/v2/api-docs"

func main()  {
	app := gin.New()
	app.Use(func(context *gin.Context) {
		log.Printf("---~ %v", context.Request.Method)
	})
	app.GET("/", func(context *gin.Context) {
		log.Print("OJBK")
		res, err := http.Get(JSON_URL)
		if err != nil {
			log.Fatal(err)
		}
		swagger, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}
		var ret swagger_mocker.SwaggerDoc
		err = json.Unmarshal(swagger, &ret)
		if err != nil {
			log.Fatal(err)
		}
		context.JSON(200, ret)
	})
	err := app.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
