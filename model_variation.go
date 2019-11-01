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

type Variation struct {
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Value *interface{} `json:"value"`
}
