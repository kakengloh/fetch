package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// XHR struct
type XHR struct {
	BaseURL string
	Headers Headers
	client  http.Client
}

// Headers is the request headers
type Headers map[string]interface{}

// Params is the request parameters
type Params map[string]interface{}

// New XML HTTP Request
func New(baseURL string, headers Headers, timeout time.Duration) (*XHR, error) {

	if baseURL == "" {
		return nil, fmt.Errorf("parameter \"baseURL\" is required")
	}

	if headers == nil {
		headers = Headers{}
	}

	return &XHR{
		BaseURL: baseURL,
		Headers: headers,
		client:  http.Client{Timeout: timeout},
	}, nil

}

// GetJSON request
func (x *XHR) GetJSON(path string, params Params, headers Headers, response interface{}) (int, error) {

	// build headers
	_headers := x.Headers

	for k, v := range headers {
		_headers[k] = v
	}

	return GetJSON(x.BaseURL+path, params, _headers, response)
}

// PostJSON request
func (x *XHR) PostJSON(path string, body map[string]interface{}, headers Headers, response interface{}) (int, error) {

	// build headers
	_headers := x.Headers

	for k, v := range headers {
		_headers[k] = v
	}

	return PostJSON(x.BaseURL+path, body, _headers, response)
}

// PutJSON request
func (x *XHR) PutJSON(path string, body map[string]interface{}, headers Headers, response interface{}) (int, error) {

	// build headers
	_headers := x.Headers

	for k, v := range headers {
		_headers[k] = v
	}

	return PutJSON(x.BaseURL+path, body, _headers, response)
}

// DeleteJSON request
func (x *XHR) DeleteJSON(path string, params Params, headers Headers, response interface{}) (int, error) {

	// build headers
	_headers := x.Headers

	for k, v := range headers {
		_headers[k] = v
	}

	return DeleteJSON(x.BaseURL+path, params, _headers, response)
}

// GetJSON static
func GetJSON(url string, params Params, headers Headers, response interface{}) (int, error) {

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return -1, err
	}

	// build query
	q := req.URL.Query()

	for k, v := range params {
		q.Add(k, v.(string))
	}

	req.URL.RawQuery = q.Encode()

	// build headers
	if headers == nil {
		headers = Headers{}
	}

	// auto insert JSON headers
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	for k, v := range headers {
		req.Header.Add(k, v.(string))
	}

	// send request
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return -1, err
	}

	return resp.StatusCode, nil

}

// PostJSON static
func PostJSON(url string, body map[string]interface{}, headers Headers, response interface{}) (int, error) {

	// JSON encode body
	b, err := json.Marshal(body)

	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(b))

	// auto insert JSON headers
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	for k, v := range headers {
		req.Header.Add(k, v.(string))
	}

	// send request
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return -1, err
	}

	return resp.StatusCode, nil
}

// PutJSON static
func PutJSON(url string, body map[string]interface{}, headers Headers, response interface{}) (int, error) {

	// JSON encode body
	b, err := json.Marshal(body)

	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewReader(b))

	// auto insert JSON headers
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	for k, v := range headers {
		req.Header.Add(k, v.(string))
	}

	// send request
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return -1, err
	}

	return resp.StatusCode, nil
}

// DeleteJSON static
func DeleteJSON(url string, params Params, headers Headers, response interface{}) (int, error) {

	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		return -1, err
	}

	// build headers
	if headers == nil {
		headers = Headers{}
	}

	// auto insert JSON headers
	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"

	for k, v := range headers {
		req.Header.Add(k, v.(string))
	}

	// send request
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return -1, err
	}

	return resp.StatusCode, nil

}
