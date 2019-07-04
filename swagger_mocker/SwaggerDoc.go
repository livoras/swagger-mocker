package swagger_mocker

type SwaggerDoc struct {
	BasePath string `json:"basePath"`
	Host string `json:"host"`
	Swagger string `json:"swagger"`
	Tags []struct{
		Description string `json:"description"`
		Name string `json:"name"`
	} `json:"tags"`
	Info struct{
		Description string `json:"description"`
		Title string `json:"title"`
		Version string `json:"version"`
	} `json:"info"`
	Definitions map[string]struct{
		Properties struct{
			Code struct{
				Description string `json:"description"`
				Format string `json:"format"`
				Type string `json:"type"`
			} `json:"code"`
			Data struct{
				Ref string `json:"$ref"`
			} `json:"data"`
			Msg struct{
				Description string `json:"description"`
				Type string `json:"type"`
			} `json:"msg"`
			Success struct{
				Type string `json:"type"`
			} `json:"success"`
		} `json:"properties"`
		Title string `json:"title"`
		Type string `json:"type"`
	} `json:"definitions"`
	Paths map[string]struct{
		Post Api `json:"post"`
		Get Api `json:"get"`
		PUT Api `json:"put"`
		DELETE Api `json:"delete"`
	} `json:"paths"`
}

type Api struct {
	Consumes []string `json:"consumes"`
	Deprecated bool `json:"deprecated"`
	Description string `json:"description"`
	OperationId string `json:"operationId"`
	Parameters []struct{
		Description string `json:"description"`
		In string `json:"in"`
		Name string `json:"name"`
		Required bool `json:"required"`
		Type string `json:"type"`
		Format string `json:"format"`
		Schema struct{
			Ref string `json:"$ref"`
		} `json:"schema"`
	} `json:"parameters"`
	Produces []string `json:"produces"`
	Responses map[uint]struct{
		Description string `json:"description"`
		Schema struct{
			Ref string `json:"$ref"`
		} `json:"schema"`
	} `json:"responses"`
	Summary string `json:"summary"`
	Tags []string `json:"tags"`
}
