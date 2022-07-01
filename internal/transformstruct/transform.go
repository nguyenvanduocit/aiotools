package transformstruct

import (
	"encoding/json"
	"fmt"

	"github.com/twpayne/go-jsonstruct"
)

type TransformStruct struct {
}

func NewPlugin() *TransformStruct {
	return &TransformStruct{}
}

type Request struct {
	SourceCode          string
	SourceLanguage      string
	DestinationLanguage string
}

func (a *TransformStruct) TransformStruct(input Request) (*string, error) {
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
