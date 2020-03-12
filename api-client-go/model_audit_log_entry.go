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

type AuditLogEntry struct {
	Links *Links `json:"_links,omitempty"`
	// The unique resource id.
	Id string `json:"_id,omitempty"`
	Date int64 `json:"date,omitempty"`
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ShortDescription string `json:"shortDescription,omitempty"`
	Comment string `json:"comment,omitempty"`
	Member *Member `json:"member,omitempty"`
	TitleVerb string `json:"titleVerb,omitempty"`
	Title string `json:"title,omitempty"`
	Target *AuditLogEntryTarget `json:"target,omitempty"`
}
