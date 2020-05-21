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

type Webhook struct {
	Links *Links `json:"_links,omitempty"`
	// The unique resource id.
	Id string `json:"_id,omitempty"`
	// The URL of the remote webhook.
	Url string `json:"url,omitempty"`
	// If defined, the webhooks post request will include a X-LD-Signature header whose value will contain an HMAC SHA256 hex digest of the webhook payload, using the secret as the key.
	Secret string `json:"secret,omitempty"`
	// Whether this webhook is enabled or not.
	On bool `json:"on,omitempty"`
	// The name of the webhook.
	Name string `json:"name,omitempty"`
	Statements []Statement `json:"statements,omitempty"`
	// Tags assigned to this webhook.
	Tags []string `json:"tags,omitempty"`
}
