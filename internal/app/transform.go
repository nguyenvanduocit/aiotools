package app

import (
	"encoding/json"
	"fmt"

	"github.com/twpayne/go-jsonstruct"
)

type StructTransformRequest struct{}

type TransformStructRequest struct {
	SourceCode          string
	SourceLanguage      string
	DestinationLanguage string
}

func (a *StructTransformRequest) TransformStruct(input TransformStructRequest) (*string, error) {
	options := []jsonstruct.GeneratorOption{
		jsonstruct.WithOmitEmpty(jsonstruct.OmitEmptyAuto),
		jsonstruct.WithSkipUnparseableProperties(true),
		jsonstruct.WithUseJSONNumber(false),
		jsonstruct.WithGoFormat(true),
	}

	var value interface{}
	err := json.Unmarshal([]byte(input.SourceCode), &value)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %w", err)
	}

	observedValue := &jsonstruct.ObservedValue{}
	observedValue = observedValue.Merge(value)

	goCode, err := jsonstruct.NewGenerator(options...).GoCode(observedValue)
	if err != nil {
		return nil, fmt.Errorf("failed to generate go code: %w", err)
	}

	output := string(goCode)

	return &output, nil
}
