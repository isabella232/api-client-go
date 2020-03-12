/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 2.0.31
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type DestinationMParticle struct {
	ApiKey string `json:"apiKey,omitempty"`
	Secret string `json:"secret,omitempty"`
	UserIdentity string `json:"userIdentity,omitempty"`
	Environment string `json:"environment,omitempty"`
}
