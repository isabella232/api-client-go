/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 3.2.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type FeatureFlagCopyObject struct {
	// The environment key to be used.
	Key string `json:"key"`
	// If the latest version of the flag matches provided version it will copy, otherwise it will return a conflict.
	CurrentVersion int32 `json:"currentVersion,omitempty"`
}
