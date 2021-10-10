package clientapi

import (
	"net/http"
	"net/url"
	"strings"
)

type IRESTClient interface {
	Verb(verb string) *Request
	Post() *Request
	Put() *Request
	Get() *Request
	Delete() *Request
}

type RESTClient struct {
	base   *url.URL
	client *http.Client
}

// NewRESTClient creates a new RESTClient. This client performs generic REST functions
// such as Get, Put, Post, and Delete on specified paths.
func NewRESTClient(baseURL *url.URL, client *http.Client) (*RESTClient, error) {

	// set default BE url

	base := *baseURL
	if !strings.HasSuffix(base.Path, "/") {
		base.Path += "/"
	}
	base.RawQuery = ""
	base.Fragment = ""

	return &RESTClient{
		base:   &base,
		client: client,
	}, nil
}
func (c *RESTClient) Verb(verb string) *Request {
	return NewRequest(c).Verb(verb)
}

// Post begins a POST request. Short for c.Verb("POST").
func (c *RESTClient) Post() *Request {
	return c.Verb("POST")
}

// Put begins a PUT request. Short for c.Verb("PUT").
func (c *RESTClient) Put() *Request {
	return c.Verb("PUT")
}

// Get begins a GET request. Short for c.Verb("GET").
func (c *RESTClient) Get() *Request {
	return c.Verb("GET")
}

// Delete begins a DELETE request. Short for c.Verb("DELETE").
func (c *RESTClient) Delete() *Request {
	return c.Verb("DELETE")
}
