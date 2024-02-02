package httpclient

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

func NewHttpClient(timeout int, maxIdleConn int, MaxIdleConnHost int) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        maxIdleConn,
			MaxIdleConnsPerHost: MaxIdleConnHost,
			DisableKeepAlives:   false,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		},
	}
}

func ComposeCredentials(user string, pass string) []byte {
	body := url.Values{}
	body.Set("username", user)
	body.Add("password", pass)
	return []byte(body.Encode())
}
