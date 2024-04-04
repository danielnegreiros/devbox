package rest

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/danielnegreiros/go-proxmox-cli/internal/app/errs"
	httpclient "github.com/danielnegreiros/go-proxmox-cli/internal/app/http_client"
)

type HttpRequest struct {
	Timeout       int
	EndPoint      string
	Method        string
	Body          []byte
	AcceptedCodes []int
	Data          any
	Cookie        *http.Cookie
	Header        map[string]string
}

func (r *HttpRequest) Execute() (any, int) {

	client := httpclient.NewHttpClient(r.Timeout, 10, 10)

	req, err := http.NewRequest(r.Method, r.EndPoint, bytes.NewBuffer(r.Body))
	errs.PanicIfErr(err)

	if r.Cookie != nil {
		req.AddCookie(r.Cookie)
	}

	for k, v := range r.Header {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	errs.PanicIfErr(err)
	defer resp.Body.Close()

	contentBytes, err := io.ReadAll(resp.Body)
	errs.PanicIfErr(err)

	log.Printf("%s %s %d", req.Method, req.URL.Path, resp.StatusCode)
	PanicIBadCode(resp.StatusCode, r.AcceptedCodes, contentBytes)

	err = json.Unmarshal(contentBytes, r.Data)
	errs.PanicIfErr(err)

	return r.Data, resp.StatusCode
}

func PanicIBadCode(code int, allowedS []int, msg []byte) {
	for _, allow := range allowedS {
		if allow == code {
			return
		}
	}
	log.Println(string(msg))
	log.Fatal(code)
}
