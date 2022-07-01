package dnslookup

import (
	"net"
)

type DnsLookup struct{}

func NewPlugin() *DnsLookup {
	return &DnsLookup{}
}

type Input struct {
	Host string
}

type Output struct {
	NS    []string
	CNAME []string
	IP    []string
	MX    []string
	TXT   []string
}

func (d *DnsLookup) Lookup(input Input) (*Output, error) {
	result := &Output{}
	var err error

	result.NS, err = listNameServer(input.Host)
	if err != nil {
		return result, err
	}

	result.CNAME, err = listCName(input.Host)
	if err != nil {
		return result, err
	}

	result.IP, err = listIP(input.Host)
	if err != nil {
		return result, err
	}

	result.MX, err = listMX(input.Host)
	if err != nil {
		return result, err
	}

	result.TXT, err = listTxt(input.Host)
	if err != nil {
		return result, err
	}

	return result, nil
}

func listNameServer(host string) ([]string, error) {
	var result []string
	nameservers, err := net.LookupNS(host)
	if err != nil {
		return result, err
	}

	for _, nameserver := range nameservers {
		result = append(result, nameserver.Host)
	}

	return result, nil
}

func listCName(url string) ([]string, error) {
	cname, err := net.LookupCNAME(url)
	if err != nil {
		return nil, err
	}

	return []string{cname}, nil
}

func listIP(url string) ([]string, error) {
	ips, err := net.LookupIP(url)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, ip := range ips {
		result = append(result, ip.String())
	}

	return result, nil
}

func listMX(url string) ([]string, error) {
	mxs, err := net.LookupMX(url)
	if err != nil {
		return nil, err
	}

	var result []string

	for _, mx := range mxs {
		result = append(result, mx.Host)
	}

	return result, nil
}

func listTxt(url string) ([]string, error) {
	txts, err := net.LookupTXT(url)
	if err != nil {
		return nil, err
	}

	var result []string

	for _, txt := range txts {
		result = append(result, txt)
	}

	return result, nil
}
