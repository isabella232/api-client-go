/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.20
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type StreamSdkVersionData struct {
	// The language of the sdk
	Sdk string `json:"sdk,omitempty"`
	// The version of the sdk
	Version string `json:"version,omitempty"`
}
