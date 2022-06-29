package app

import (
	"context"
	"fmt"
	"net/url"
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
	Scheme      string `json:"scheme,omitempty"`
	Opaque      string `json:"opaque,omitempty"`
	Host        string `json:"host,omitempty"`
	Path        string `json:"path,omitempty"`         // path (relative paths may omit leading slash)
	RawPath     string `json:"raw_path,omitempty"`     // encoded path hint (see EscapedPath method)
	ForceQuery  bool   `json:"force_query,omitempty"`  // append a query ('?') even if RawQuery is empty
	RawQuery    string `json:"raw_query,omitempty"`    // encoded query values, without '?'
	Fragment    string `json:"fragment,omitempty"`     // fragment for references, without '#'
	RawFragment string `json:"raw_fragment,omitempty"` // encoded fragment hint (see EscapedFragment method)
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

	return result, nil
}
