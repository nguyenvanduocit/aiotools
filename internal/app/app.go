package app

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/twpayne/go-jsonstruct"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

type ParseResult struct {
	Scheme      string              `json:"scheme,omitempty"`
	Opaque      string              `json:"opaque,omitempty"`
	Host        string              `json:"host,omitempty"`
	Path        string              `json:"path,omitempty"`
	RawPath     string              `json:"raw_path,omitempty"`
	ForceQuery  bool                `json:"force_query,omitempty"`
	RawQuery    string              `json:"raw_query,omitempty"`
	Fragment    string              `json:"fragment,omitempty"`
	RawFragment string              `json:"raw_fragment,omitempty"`
	Query       map[string][]string `json:"query,omitempty"`
}

func (a *App) ParseUrl(s string) (*ParseResult, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	result := &ParseResult{
		Scheme:      u.Scheme,
		Opaque:      u.Opaque,
		Host:        u.Host,
		Path:        u.Path,
		RawPath:     u.RawPath,
		ForceQuery:  u.ForceQuery,
		RawQuery:    u.RawQuery,
		Fragment:    u.Fragment,
		RawFragment: u.RawFragment,
	}

	if u.RawQuery != "" {
		result.Query = make(map[string][]string)
		for s2, strings := range u.Query() {
			result.Query[s2] = strings
		}
	}

	return result, nil
}

type JsonToGoInput struct {
	JsonCode string
}

func (a *App) JsonToGo(input JsonToGoInput) (*string, error) {
	options := []jsonstruct.GeneratorOption{
		jsonstruct.WithOmitEmpty(jsonstruct.OmitEmptyAuto),
		jsonstruct.WithSkipUnparseableProperties(true),
		jsonstruct.WithUseJSONNumber(false),
		jsonstruct.WithGoFormat(true),
	}

	var value interface{}
	err := json.Unmarshal([]byte(input.JsonCode), &value)
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
