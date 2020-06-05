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

type SemanticPatchOperationInstructions struct {
	// The name of the modification you would like to perform on a resource.
	Kind string `json:"kind,omitempty"`
}
