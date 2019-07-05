package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"swagger-mocker/swagger_mocker"
)

const JSON_URL = "http://8789.office.qunjielong.com/v2/api-docs"

func main()  {
	app := gin.New()
	ret := GetSwaggerDocs()
	rootRouter := swagger_mocker.NewRouter()
	for path, apiGroup := range ret.Paths {
		subPaths := splitPath(path)
		log.Print("===> ", path)
		rootRouter.AddChild(subPaths, apiGroup)
	}
	app.Use(func(context *gin.Context) {
		path := context.Request.URL.Path
		paths := splitPath(path)
		api := rootRouter.FindApi(paths)
		context.JSON(200, api)
	})
	err := app.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

func removeEmpty(split []string) []string {
	var ret []string
	for _, str := range split {
		if len(str) != 0 {
			ret = append(ret, str)
		}
	}
	return ret
}

func splitPath(path string) []string {
	return removeEmpty(strings.Split(path, "/"))
}

func GetSwaggerDocs() swagger_mocker.SwaggerDoc {
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
	return ret
}
