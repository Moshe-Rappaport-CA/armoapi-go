package apiinterfaces

import v1 "armoapi-go/armoapi/v1"

type Interface interface {
	V1() v1.ApiInterface
}

type Client struct {
	v1 *v1.ClientApi
}

// AppsV1 retrieves the AppsV1Client
func (c *Client) V1() v1.ApiInterface {
	return c.v1
}
