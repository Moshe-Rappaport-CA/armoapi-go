package v1

import "armoapi-go/armoapi/v1/frameworks"

const APIVersion = "v1"

type ApiInterface interface {
	Framework() frameworks.IFramework
	// Control() control.IControl
	// Rule() rule.IRule
}

type ClientApi struct {
}

func (c *ClientApi) Framework() frameworks.IFramework {
	return &frameworks.FrameworkAPI{}
}
