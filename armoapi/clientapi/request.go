package clientapi

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Request struct {
	c *RESTClient

	timeout time.Duration

	// generic components accessible via method setters
	verb       string
	pathPrefix string
	path       string
	version    string
	params     url.Values
	headers    http.Header

	resource string

	// output
	err  error
	body io.Reader
	// retry Retry
}

// Verb sets the verb this request will use.
func (r *Request) Verb(verb string) *Request {
	r.verb = verb
	return r
}

// NewRequest creates a new request helper object for accessing runtime.Objects on a server.
func NewRequest(c *RESTClient) *Request {

	var pathPrefix string

	if c.base != nil {
		pathPrefix = c.base.Path
	} else {
		pathPrefix = "" // TODO - handle empty path
	}

	var timeout time.Duration
	if c.client != nil {
		timeout = c.client.Timeout
	}

	r := &Request{
		c:          c,
		timeout:    timeout,
		pathPrefix: pathPrefix,
	}

	// switch {
	// case len(c.content.AcceptContentTypes) > 0:
	// 	r.SetHeader("Accept", c.content.AcceptContentTypes)
	// case len(c.content.ContentType) > 0:
	// 	r.SetHeader("Accept", c.content.ContentType+", */*")
	// }
	return r
}

// Resource sets the resource to access (<resource>/[ns/<namespace>/]<name>)
func (r *Request) Resource(resource string) *Request {
	if r.err != nil {
		return r
	}

	// TODO - validate resource

	r.resource = resource
	return r
}

// Resource sets the resource to access (<resource>/[ns/<namespace>/]<name>)
func (r *Request) Path(path string) *Request {
	if r.err != nil {
		return r
	}

	// TODO - validate path

	r.path = path
	return r
}

// Resource sets the resource to access (<resource>/[ns/<namespace>/]<name>)
func (r *Request) Version(v string) *Request {
	if r.err != nil {
		return r
	}

	// TODO - validate path

	r.version = v
	return r
}

// Encode is a convenience wrapper for encoding to a []byte from an Encoder
func Encode(obj interface{}) ([]byte, error) {
	buf, err := json.Marshal(obj)
	return buf, err
}

func (r *Request) Body(obj interface{}) *Request {
	if r.err != nil {
		return r
	}
	switch t := obj.(type) {
	case string:
		data, err := ioutil.ReadFile(t)
		if err != nil {
			r.err = err
			return r
		}
		r.body = bytes.NewReader(data)
	case []byte:
		r.body = bytes.NewReader(t)
	case io.Reader:
		r.body = t
	default:

		data, err := Encode(obj)
		if err != nil {
			r.err = err
			return r
		}
		r.body = bytes.NewReader(data)
		r.SetHeader("Content-Type", "application/json") // TODO - support more headers
		// default:
		// 	r.err = fmt.Errorf("unknown type used for body: %+v", obj)
	}
	return r
}

func (r *Request) SetHeader(key string, values ...string) *Request {
	if r.headers == nil {
		r.headers = http.Header{}
	}
	r.headers.Del(key)
	for _, value := range values {
		r.headers.Add(key, value)
	}
	return r
}
func (r *Request) Do(ctx context.Context) Result {
	var result Result
	err := r.request(ctx, func(req *http.Request, resp *http.Response) {
		result = r.transformResponse(resp, req)
	})
	if err != nil {
		return Result{err: err}
	}
	return result
}

// request connects to the server and invokes the provided function when a server response is
// received. It handles retry behavior and up front validation of requests. It will invoke
// fn at most once. It will return an error if a problem occurred prior to connecting to the
// server - the provided function is responsible for handling server errors.
func (r *Request) request(ctx context.Context, fn func(*http.Request, *http.Response)) error {

	client := r.c.client
	if client == nil {
		client = http.DefaultClient
	}

	req, err := r.newHTTPRequest(ctx)
	if err != nil {
		return err
	}

	_, err = client.Do(req)
	return err
}

func (r *Request) newHTTPRequest(ctx context.Context) (*http.Request, error) {
	url := r.URL().String()
	req, err := http.NewRequest(r.verb, url, r.body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	req.Header = r.headers
	return req, nil
}

// URL returns the current working URL.
func (r *Request) URL() *url.URL {
	return &url.URL{}
}

// transformResponse converts an API response into a structured API object
func (r *Request) transformResponse(resp *http.Response, req *http.Request) Result {
	var body []byte
	if resp.Body != nil {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return Result{err: err}
		}
		body = data
	}
	return Result{
		body:       body,
		statusCode: resp.StatusCode,
	}
}
