/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.1
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type Policy struct {

	Resources *Resources `json:"resources,omitempty"`

	Actions *Actions `json:"actions,omitempty"`

	// Effect of the policy - allow or deny.
	Effect string `json:"effect,omitempty"`
}
