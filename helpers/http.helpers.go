package helpers

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/parnurzeal/gorequest"
)

// Env ..
type Env struct {
	DebugClient bool   `envconfig:"DEBUG_CLIENT" default:"true"`
	Timeout     string `envconfig:"TIMEOUT" default:"60s"`
	RetryBad    int    `envconfig:"RETRY_BAD" default:"1"`
}

var (
	httpEnv Env
)

// HTTPMethodGet ...
const (
	HTTPMethodGet      = "GET"
	HTTPMethodPost     = "POST"
	HTTPMethodPut      = "PUT"
	HTTPMethodDelete   = "DELETE"
	HTTPMethodPostForm = "POST_FORM"
)

func init() {
	if err := envconfig.Process("HTTP", &httpEnv); err != nil {
		fmt.Println("Failed to get HTTP env:", err)
	}
}

// HTTPGet func
func HTTPGet(url string, header http.Header) (gorequest.Response, []byte, error) {
	request := gorequest.New()
	request.SetDebug(httpEnv.DebugClient)
	timeout, _ := time.ParseDuration(httpEnv.Timeout)
	reqagent := request.Get(url)
	reqagent.Header = header
	resp, body, errs := reqagent.
		Timeout(timeout).
		Retry(httpEnv.RetryBad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return resp, []byte(body), errs[0]
	}
	return resp, []byte(body), nil
}

// HTTPPost func
func HTTPPost(url string, jsondata interface{}) (gorequest.Response, []byte, error) {
	request := gorequest.New()
	request.SetDebug(httpEnv.DebugClient)
	timeout, _ := time.ParseDuration(httpEnv.Timeout)
	// _ := errors.New("Connection Problem")
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Post(url)
	reqagent.Header.Set("Content-Type", "application/json")
	resp, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(httpEnv.RetryBad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return resp, []byte(body), errs[0]
	}
	return resp, []byte(body), nil
}

// HTTPPostWithHeader func
func HTTPPostWithHeader(url string, jsondata interface{}, header http.Header) (gorequest.Response, []byte, error) {
	request := gorequest.New()
	request.SetDebug(httpEnv.DebugClient)
	timeout, _ := time.ParseDuration(httpEnv.Timeout)
	reqagent := request.Post(url)
	reqagent.Header = header
	resp, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(httpEnv.RetryBad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return resp, []byte(body), errs[0]
	}
	return resp, []byte(body), nil
}

// HTTPPostFormWithHeader func
func HTTPPostFormWithHeader(url string, jsondata interface{}, header http.Header) (gorequest.Response, []byte, error) {
	request := gorequest.New()
	request.SetDebug(httpEnv.DebugClient)
	timeout, _ := time.ParseDuration(httpEnv.Timeout)
	reqagent := request.Post(url)
	reqagent.Header = header
	resp, body, errs := reqagent.
		Type("multipart").
		SendFile("client_credentials", "grand_type").
		Timeout(timeout).
		Retry(httpEnv.RetryBad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return resp, []byte(body), errs[0]
	}
	return resp, []byte(body), nil
}

// HTTPPutWithHeader func
func HTTPPutWithHeader(url string, jsondata interface{}, header http.Header) (gorequest.Response, []byte, error) {
	request := gorequest.New()
	request.SetDebug(httpEnv.DebugClient)
	timeout, _ := time.ParseDuration(httpEnv.Timeout)
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Put(url)
	reqagent.Header = header
	resp, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(httpEnv.RetryBad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return resp, []byte(body), errs[0]
	}
	return resp, []byte(body), nil
}

// HTTPDeleteWithHeader func
func HTTPDeleteWithHeader(url string, jsondata interface{}, header http.Header) (gorequest.Response, []byte, error) {
	request := gorequest.New()
	request.SetDebug(httpEnv.DebugClient)
	timeout, _ := time.ParseDuration(httpEnv.Timeout)
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Delete(url)
	reqagent.Header = header
	resp, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(httpEnv.RetryBad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return resp, []byte(body), errs[0]
	}
	return resp, []byte(body), nil
}

// SendHTTPRequest ..
func SendHTTPRequest(method, url string, header http.Header, body interface{}) (gorequest.Response, []byte, error) {
	var response gorequest.Response
	var data []byte
	var err error
	switch method {
	case HTTPMethodGet:
		response, data, err = HTTPGet(url, header)
	case HTTPMethodPost:
		response, data, err = HTTPPostWithHeader(url, body, header)
	case HTTPMethodPut:
		response, data, err = HTTPPutWithHeader(url, body, header)
	case HTTPMethodDelete:
		response, data, err = HTTPDeleteWithHeader(url, body, header)
	case HTTPMethodPostForm:
		response, data, err = HTTPPostFormWithHeader(url, body, header)
	}
	return response, data, err
}

// PostFormDataV2 ..
func PostFormDataV2(uri string, params map[string]string, auth, path, paramName string) (*http.Request, error) {

	var err error
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if path != "" {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		part, err := writer.CreateFormFile(paramName, filepath.Base(path))
		if err != nil {
			return nil, err
		}
		io.Copy(part, file)
	}

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", auth)
	return req, err
}
