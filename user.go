/* 
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * OpenAPI spec version: 2.0.0
 * Contact: support@launchdarkly.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

type User struct {

	LastPing string `json:"lastPing,omitempty"`

	EnvironmentId string `json:"environmentId,omitempty"`

	OwnerId string `json:"ownerId,omitempty"`

	User interface{} `json:"user,omitempty"`

	Avatar string `json:"avatar,omitempty"`
}
