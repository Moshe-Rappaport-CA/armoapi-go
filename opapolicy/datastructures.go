package opapolicy

import (
	"time"

	"github.com/armosec/armoapi-go/armotypes"
)

type AlertScore float32
type RuleLanguages string

const (
	RegoLanguage  RuleLanguages = "Rego"
	RegoLanguage2 RuleLanguages = "rego"
)

// RegoResponse the expected response of single run of rego policy
type RuleResponse struct {
	AlertMessage  string                            `json:"alertMessage"`
	RuleStatus    string                            `json:"ruleStatus"`
	PackageName   string                            `json:"packagename"`
	AlertScore    AlertScore                        `json:"alertScore"`
	AlertObject   AlertObject                       `json:"alertObject"`
	Context       []string                          `json:"context,omitempty"`       // TODO - Remove
	Rulename      string                            `json:"rulename,omitempty"`      // TODO - Remove
	ExceptionName string                            `json:"exceptionName,omitempty"` // Not in use
	Exception     *armotypes.PostureExceptionPolicy `json:"exception,omitempty"`
}

type AlertObject struct {
	K8SApiObjects   []map[string]interface{} `json:"k8sApiObjects,omitempty"`
	ExternalObjects map[string]interface{}   `json:"externalObjects,omitempty"`
}

type FrameworkReport struct {
	Name            string          `json:"name"`
	ControlReports  []ControlReport `json:"controlReports"`
	Score           float32         `json:"score,omitempty"`
	ARMOImprovement float32         `json:"ARMOImprovement,omitempty"`
	WCSScore        float32         `json:"wcsScore,omitempty"`
}
type ControlReport struct {
	armotypes.PortalBase `json:",inline"`
	ControlID            string `json:"id"`

	Name            string       `json:"name"`
	RuleReports     []RuleReport `json:"ruleReports"`
	Remediation     string       `json:"remediation"`
	Description     string       `json:"description"`
	Score           float32      `json:"score,omitempty"`
	BaseScore       float32      `json:"baseScore,omitempty"`
	ARMOImprovement float32      `json:"ARMOImprovement,omitempty"`
}
type RuleReport struct {
	Name               string                   `json:"name"`
	Remediation        string                   `json:"remediation"`
	RuleStatus         RuleStatus               `json:"ruleStatus"` // did we run the rule or not (if there where compile errors, the value will be failed)
	RuleResponses      []RuleResponse           `json:"ruleResponses"`
	ListInputResources []map[string]interface{} `json:"-"`
}
type RuleStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// PostureReport
type PostureReport struct {
	CustomerGUID         string            `json:"customerGUID"`
	ClusterName          string            `json:"clusterName"`
	ReportID             string            `json:"reportID"`
	JobID                string            `json:"jobID"`
	ReportGenerationTime time.Time         `json:"generationTime"`
	FrameworkReports     []FrameworkReport `json:"frameworks"`
}

// RuleMatchObjects defines which objects this rule applied on
type RuleMatchObjects struct {
	APIGroups   []string `json:"apiGroups"`   // apps
	APIVersions []string `json:"apiVersions"` // v1/ v1beta1 / *
	Resources   []string `json:"resources"`   // dep.., pods,
}

// RuleMatchObjects defines which objects this rule applied on
type RuleDependency struct {
	PackageName string `json:"packageName"` // package name
}

// PolicyRule represents single rule, the fundamental executable block of policy
type PolicyRule struct {
	armotypes.PortalBase `json:",inline"`
	CreationTime         string             `json:"creationTime"`
	Rule                 string             `json:"rule"` // multiline string!
	RuleLanguage         RuleLanguages      `json:"ruleLanguage"`
	Match                []RuleMatchObjects `json:"match"`
	RuleDependencies     []RuleDependency   `json:"ruleDependencies"`
	Description          string             `json:"description"`
	Remediation          string             `json:"remediation"`
	RuleQuery            string             `json:"ruleQuery"` // default "armo_builtins" - DEPRECATED
}

// Control represents a collection of rules which are combined together to single purpose
type Control struct {
	armotypes.PortalBase `json:",inline"`
	ControlID            string `json:"id"`

	CreationTime string       `json:"creationTime"`
	Description  string       `json:"description"`
	Remediation  string       `json:"remediation"`
	Rules        []PolicyRule `json:"rules"`
	// for new list of  rules in POST/UPADTE requests
	RulesIDs *[]string `json:"rulesIDs,omitempty"`
}

type UpdatedControl struct {
	Control `json:",inline"`
	Rules   []interface{} `json:"rules"`
}

// Framework represents a collection of controls which are combined together to expose comprehensive behavior
type Framework struct {
	armotypes.PortalBase `json:",inline"`
	CreationTime         string    `json:"creationTime"`
	Description          string    `json:"description"`
	Controls             []Control `json:"controls"`
	// for new list of  controls in POST/UPADTE requests
	ControlsIDs *[]string `json:"controlsIDs,omitempty"`
}

type UpdatedFramework struct {
	Framework `json:",inline"`
	Controls  []interface{} `json:"controls"`
}

type NotificationPolicyType string
type NotificationPolicyKind string

// Supported NotificationTypes
const (
	TypeValidateRules   NotificationPolicyType = "validateRules"
	TypeExecPostureScan NotificationPolicyType = "execPostureScan"
	TypeUpdateRules     NotificationPolicyType = "updateRules"
)

// Supported NotificationKinds
const (
	KindFramework NotificationPolicyKind = "Framework"
	KindControl   NotificationPolicyKind = "Control"
	KindRule      NotificationPolicyKind = "Rule"
)

type PolicyNotification struct {
	NotificationType NotificationPolicyType     `json:"notificationType"`
	Rules            []PolicyIdentifier         `json:"rules"`
	ReportID         string                     `json:"reportID"`
	JobID            string                     `json:"jobID"`
	Designators      armotypes.PortalDesignator `json:"designators"`
}

type PolicyIdentifier struct {
	Kind NotificationPolicyKind `json:"kind"`
	Name string                 `json:"name"`
}