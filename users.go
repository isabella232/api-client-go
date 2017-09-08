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

type Users struct {

	Links Links `json:"_links,omitempty"`

	TotalCount float32 `json:"totalCount,omitempty"`

	Items []User `json:"items,omitempty"`
}
