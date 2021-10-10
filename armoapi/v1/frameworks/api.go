package frameworks

import (
	"armoapi-go/armoapi/clientapi"
	"context"
)

// interface
type IFramework interface {
	Get() ([]Framework, error)
	// Get(*meta.GetOptions) ([]Framework, error)
}

type FrameworkAPI struct {
	clientApi clientapi.IRESTClient
}

func (frameworkAPI *FrameworkAPI) Get() ([]Framework, error) {
	frameworksLIst := []Framework{}

	err := frameworkAPI.clientApi.
		Get().
		Version("v1").            // api version
		Path("").                 // api path
		Resource("").             // use for validation - TODO
		Do(context.Background()). // do request
		Into(&frameworksLIst)     // copy response to object
	return frameworksLIst, err
}

// ArmoAPIV1().Framework().Get(GetOptions())
// ArmoAPI().V1().Framework().List(ListOptions())
