package setjson

import (
	"github.com/tidwall/sjson"
)

type SetJson struct {
}

func NewSetJson() *SetJson {
	return &SetJson{}
}

type Input struct {
	Code  string      `json:"code"`
	Query string      `json:"query"`
	Value interface{} `json:"value"`
}

func (g *SetJson) Set(input Input) (string, error) {
	if input.Value == "" {
		return sjson.Delete(input.Code, input.Query)
	}

	return sjson.Set(input.Code, input.Query, input.Value)
}
