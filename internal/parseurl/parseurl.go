package parseurl

import (
	"fmt"
	"net/url"
)

type ParseUrl struct{}

func NewPlugin() *ParseUrl {
	return &ParseUrl{}
}

type Response struct {
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

func (a *ParseUrl) ParseUrl(s string) (*Response, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url: %w", err)
	}

	result := &Response{
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
