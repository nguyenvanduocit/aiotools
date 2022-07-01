package getjson

import "github.com/tidwall/gjson"

type GetJson struct {
}

func NewGetJson() *GetJson {
	return &GetJson{}
}

type Input struct {
	Code  string `json:"code"`
	Query string `json:"query"`
}

func (g *GetJson) Get(input Input) (string, error) {
	result := gjson.Get(input.Code, input.Query)
	return result.String(), nil
}
