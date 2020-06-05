/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.3.1
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type FeatureFlagConfig struct {
	On bool `json:"on,omitempty"`
	Archived bool `json:"archived,omitempty"`
	Salt string `json:"salt,omitempty"`
	Sel string `json:"sel,omitempty"`
	LastModified int64 `json:"lastModified,omitempty"`
	Version int32 `json:"version,omitempty"`
	Targets []Target `json:"targets,omitempty"`
	Rules []Rule `json:"rules,omitempty"`
	Fallthrough_ *ModelFallthrough `json:"fallthrough,omitempty"`
	OffVariation int32 `json:"offVariation,omitempty"`
	Prerequisites []Prerequisite `json:"prerequisites,omitempty"`
	// Set to true to send detailed event information for this flag.
	TrackEvents bool `json:"trackEvents,omitempty"`
	// Set to true to send detailed event information when targeting is enabled but no individual targeting rule is matched.
	TrackEventsFallthrough bool `json:"trackEventsFallthrough,omitempty"`
}
