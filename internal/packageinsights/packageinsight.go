package packageinsights

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/tidwall/gjson"
)

func NewPlugin() *PackageInsights {
	return &PackageInsights{}
}

type PackageInsights struct {
}

type Request struct {
	PackageName string `json:"package_name"`
}

func (a *PackageInsights) GetPackageInsight(request *Request) (*Response, error) {
	endpoint := `https://api.npms.io/v2/package/` + url.QueryEscape(request.PackageName)
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if apiErr := gjson.GetBytes(body, "error"); apiErr.Exists() {
		return nil, errors.New(apiErr.String())
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
