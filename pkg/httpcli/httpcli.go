// Package httpcli is http request client, which only supports return json format.
package httpcli

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

const defaultTimeout = 30 * time.Second

// Request HTTP request
type Request struct {
	customRequest func(req *http.Request, data *bytes.Buffer) // used to define HEADER, e.g. to add sign, etc.
	url           string
	params        map[string]interface{} // parameters after URL
	body          string                 // Body data
	bodyJSON      interface{}            // JSON marshal body data
	timeout       time.Duration          // Client timeout
	headers       map[string]string

	request  *http.Request
	response *Response
	method   string
	err      error
}

// Response HTTP response
type Response struct {
	*http.Response
	err error
}

// -----------------------------------  Request way 1 -----------------------------------

// New create a new Request
func New() *Request {
	_ = "STUB: not implemented"

	// Reset set all fields to default value, use at pool
	return nil
}

func (req *Request) Reset() { _ = "STUB: not implemented"; return }

// SetURL set URL
func (req *Request) SetURL(path string) *Request { _ = "STUB: not implemented"; return nil }

// SetParams parameters after setting the URL
func (req *Request) SetParams(params map[string]interface{}) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetParam parameters after setting the URL
func (req *Request) SetParam(k string, v interface{}) *Request {
	_ = "STUB: not implemented"
	return nil
}

// SetBody set body data, support string and []byte, if it is not string, it will be json marshal.
func (req *Request) SetBody(body interface{}) *Request { _ = "STUB: not implemented"; return nil }

// SetTimeout set timeout
func (req *Request) SetTimeout(t time.Duration) *Request { _ = "STUB: not implemented"; return nil }

// SetContentType set ContentType
func (req *Request) SetContentType(a string) *Request { _ = "STUB: not implemented"; return nil }

// SetHeader set the value of the request header
func (req *Request) SetHeader(k, v string) *Request { _ = "STUB: not implemented"; return nil }

// SetHeaders set the value of Request Headers
func (req *Request) SetHeaders(headers map[string]string) *Request {
	_ = "STUB: not implemented"
	return nil
}

// CustomRequest customize request, e.g. add sign, set header, etc.
func (req *Request) CustomRequest(f func(req *http.Request, data *bytes.Buffer)) *Request {
	_ = "STUB: not implemented"
	return nil
}

// GET send a GET request
func (req *Request) GET() (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// DELETE send a DELETE request
func (req *Request) DELETE() (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// POST send a POST request
func (req *Request) POST() (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// PUT send a PUT request
func (req *Request) PUT() (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// PATCH send PATCH requests
func (req *Request) PATCH() (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// Do a request
func (req *Request) Do(method string, data interface{}) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint

func (req *Request) pull() (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

func (req *Request) push() (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

func (req *Request) send(body io.Reader, buf *bytes.Buffer) (*Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Response return response
func (req *Request) Response() (*Response, error) { _ = "STUB: not implemented"; return nil, nil }

// -----------------------------------  Response -----------------------------------

// Error return err
func (resp *Response) Error() error {
	_ = "STUB: not implemented"

	// BodyString returns the body data of the HttpResponse
	return nil
}

func (resp *Response) BodyString() (string, error) { _ = "STUB: not implemented"; return "", nil }

// ReadBody returns the body data of the HttpResponse
func (resp *Response) ReadBody() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// BindJSON parses the response's body as JSON
func (resp *Response) BindJSON(v interface{}) error { _ = "STUB: not implemented"; return nil }

// -----------------------------------  Request way 2 -----------------------------------

// Option set options.
type Option func(*options)

type options struct {
	params  map[string]interface{}
	headers map[string]string
	timeout time.Duration
}

func (o *options) apply(opts ...Option) { _ = "STUB: not implemented"; return }

func defaultOptions() *options {
	_ = "STUB: not implemented"

	// WithParams set params
	return nil
}

func WithParams(params map[string]interface{}) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

// WithHeaders set headers
func WithHeaders(headers map[string]string) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithTimeout set timeout
func WithTimeout(t time.Duration) Option { _ = "STUB: not implemented"; return *new(Option) }

// Get request, return custom json format
func Get(result interface{}, urlStr string, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Delete request, return custom json format
func Delete(result interface{}, urlStr string, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Post request, return custom json format
func Post(result interface{}, urlStr string, body interface{}, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Put request, return custom json format
func Put(result interface{}, urlStr string, body interface{}, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

// Patch request, return custom json format
func Patch(result interface{}, urlStr string, body interface{}, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

var requestErr = func(err error) error { return fmt.Errorf("request error, err=%v", err) }
var jsonParseErr = func(err error) error { return fmt.Errorf("json parsing error, err=%v", err) }
var notOKErr = func(resp *Response) error {
	body, err := resp.ReadBody()
	if err != nil {
		return err
	}
	if len(body) > 500 {
		body = append(body[:500], []byte(" ......")...)
	}
	return fmt.Errorf("statusCode=%d, body=%s", resp.StatusCode, body)
}

func do(method string, result interface{}, urlStr string, body interface{}, params KV, headers map[string]string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint

func gDo(method string, result interface{}, urlStr string, params KV, headers map[string]string, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint

// StdResult standard return data
type StdResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

// KV string:interface{}
type KV = map[string]interface{}
