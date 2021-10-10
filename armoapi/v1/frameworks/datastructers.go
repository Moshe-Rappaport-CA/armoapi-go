package frameworks

// Framework represents a collection of controls which are combined together to expose comprehensive behavior
type Framework struct {
	// armotypes.PortalBase `json:",inline"`
	CreationTime string `json:"creationTime"`
	Description  string `json:"description"`
	// Controls             []Control `json:"controls"`
	// for new list of  controls in POST/UPADTE requests
	ControlsIDs *[]string `json:"controlsIDs,omitempty"`
}

func NewFramework() *Framework {
	return &Framework{}
}
