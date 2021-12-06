package gorseclient

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type API struct {
	Error error `json:"error"`

	client    *http.Client
	request   *http.Request
	baseUrl   string
	token     string
	debugMode DebugMode
}

type Option struct {
	HttpClient *http.Client `json:"http_client"`
	DebugMode  DebugMode    `json:"debug_mode"`
}

func NewAPI(baseUrl string, token string, option ...Option) *API {
	api := API{
		baseUrl:   baseUrl,
		token:     token,
		client:    http.DefaultClient,
		debugMode: DEBUG_SILENT,
	}

	if len(option) > 0 {
		if option[0].HttpClient != nil {
			api.client = option[0].HttpClient
		}
		if option[0].DebugMode != 0 {
			api.debugMode = option[0].DebugMode
		}
	}

	return &api
}

func (a *API) newRequest(method string, endpoint string, data []byte) *API {
	var body io.Reader
	if data != nil {
		body = bytes.NewBuffer(data)
	}

	request, err := http.NewRequest(method, a.baseUrl+endpoint, body)
	if err != nil {
		a.debugError(err)
		a.Error = err
		return a
	}
	request.Header.Set("Accept", "*/*")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", a.token)
	a.request = request
	return a
}

func (a *API) do() ([]byte, error) {
	resp, err := a.client.Do(a.request)
	if err != nil {
		a.debugError(err)
		return nil, err
	}
	a.debugInfo(fmt.Sprintf("METHOD: %s", a.request.Method))
	a.debugInfo(fmt.Sprintf("Endpoint Hit: %s\n", a.request.URL))
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		a.debugError(err)
		return nil, err
	}

	a.debugInfo(fmt.Sprintf("Resp Status Code : %d", resp.StatusCode))
	a.debugInfo(fmt.Sprintf("Resp Status : %s", resp.Status))
	a.debugInfo(fmt.Sprintf("Resp Body : %s\n", string(body)))

	if resp.StatusCode == http.StatusNotFound && a.request.Method == http.MethodGet {
		return nil, Err404NotFound
	}
	if resp.StatusCode == http.StatusInternalServerError && a.request.Method == http.MethodPost {
		return nil, Err500InternalSystemError
	}

	return body, err
}

func (a *API) debugInfo(info string) {
	if a.debugMode == DEBUG_INFO {
		log.Println(info)
	}
}

func (a *API) debugError(err error) {
	if a.debugMode == DEBUG_INFO || a.debugMode == DEBUG_ERROR {
		log.Printf("ERROR: %+v", err)
	}
}