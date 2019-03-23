This repository contains a client library for LaunchDarkly's REST API. This client was automatically
generated from our [OpenAPI specification](https://github.com/launchdarkly/ld-openapi).

This REST API is for custom integrations, data export, or automating your feature flag workflows. *DO NOT* use this client library to include feature flags in your web or mobile application. To integrate feature flags with your application, please see the [SDK documentation](https://docs.launchdarkly.com/v2.0/docs)

# Go API client for ldapi

Build custom integrations with the LaunchDarkly REST API

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 2.0.15
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [https://support.launchdarkly.com](https://support.launchdarkly.com)

## Installation

```golang
import "github.com/launchdarkly/api-client-go"
```

## Documentation for API Endpoints

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuditLogApi* | [**GetAuditLogEntries**](docs/AuditLogApi.md#getauditlogentries) | **Get** /auditlog | Get a list of all audit log entries. The query parameters allow you to restrict the returned results by date ranges, resource specifiers, or a full-text search query.
*AuditLogApi* | [**GetAuditLogEntry**](docs/AuditLogApi.md#getauditlogentry) | **Get** /auditlog/{resourceId} | Use this endpoint to fetch a single audit log entry by its resouce ID.
*CustomRolesApi* | [**DeleteCustomRole**](docs/CustomRolesApi.md#deletecustomrole) | **Delete** /roles/{customRoleKey} | Delete a custom role by key.
*CustomRolesApi* | [**GetCustomRole**](docs/CustomRolesApi.md#getcustomrole) | **Get** /roles/{customRoleKey} | Get one custom role by key.
*CustomRolesApi* | [**GetCustomRoles**](docs/CustomRolesApi.md#getcustomroles) | **Get** /roles | Return a complete list of custom roles.
*CustomRolesApi* | [**PatchCustomRole**](docs/CustomRolesApi.md#patchcustomrole) | **Patch** /roles/{customRoleKey} | Modify a custom role by key.
*CustomRolesApi* | [**PostCustomRole**](docs/CustomRolesApi.md#postcustomrole) | **Post** /roles | Create a new custom role.
*EnvironmentsApi* | [**DeleteEnvironment**](docs/EnvironmentsApi.md#deleteenvironment) | **Delete** /projects/{projectKey}/environments/{environmentKey} | Delete an environment in a specific project.
*EnvironmentsApi* | [**GetEnvironment**](docs/EnvironmentsApi.md#getenvironment) | **Get** /projects/{projectKey}/environments/{environmentKey} | Get an environment given a project and key.
*EnvironmentsApi* | [**PatchEnvironment**](docs/EnvironmentsApi.md#patchenvironment) | **Patch** /projects/{projectKey}/environments/{environmentKey} | Modify an environment by ID.
*EnvironmentsApi* | [**PostEnvironment**](docs/EnvironmentsApi.md#postenvironment) | **Post** /projects/{projectKey}/environments | Create a new environment in a specified project with a given name, key, and swatch color.
*FeatureFlagsApi* | [**CopyFeatureFlag**](docs/FeatureFlagsApi.md#copyfeatureflag) | **Post** /flags/{projectKey}/{environmentKey}/{featureFlagKey}/copy | Copies the feature flag configuration from one environment to the same feature flag in another environment.
*FeatureFlagsApi* | [**DeleteFeatureFlag**](docs/FeatureFlagsApi.md#deletefeatureflag) | **Delete** /flags/{projectKey}/{featureFlagKey} | Delete a feature flag in all environments. Be careful-- only delete feature flags that are no longer being used by your application.
*FeatureFlagsApi* | [**GetFeatureFlag**](docs/FeatureFlagsApi.md#getfeatureflag) | **Get** /flags/{projectKey}/{featureFlagKey} | Get a single feature flag by key.
*FeatureFlagsApi* | [**GetFeatureFlagStatus**](docs/FeatureFlagsApi.md#getfeatureflagstatus) | **Get** /flag-statuses/{projectKey}/{environmentKey}/{featureFlagKey} | Get the status for a particular feature flag.
*FeatureFlagsApi* | [**GetFeatureFlagStatuses**](docs/FeatureFlagsApi.md#getfeatureflagstatuses) | **Get** /flag-statuses/{projectKey}/{environmentKey} | Get a list of statuses for all feature flags. The status includes the last time the feature flag was requested, as well as the state of the flag.
*FeatureFlagsApi* | [**GetFeatureFlags**](docs/FeatureFlagsApi.md#getfeatureflags) | **Get** /flags/{projectKey} | Get a list of all features in the given project.
*FeatureFlagsApi* | [**PatchFeatureFlag**](docs/FeatureFlagsApi.md#patchfeatureflag) | **Patch** /flags/{projectKey}/{featureFlagKey} | Perform a partial update to a feature.
*FeatureFlagsApi* | [**PostFeatureFlag**](docs/FeatureFlagsApi.md#postfeatureflag) | **Post** /flags/{projectKey} | Creates a new feature flag.
*ProjectsApi* | [**DeleteProject**](docs/ProjectsApi.md#deleteproject) | **Delete** /projects/{projectKey} | Delete a project by key. Caution-- deleting a project will delete all associated environments and feature flags. You cannot delete the last project in an account.
*ProjectsApi* | [**GetProject**](docs/ProjectsApi.md#getproject) | **Get** /projects/{projectKey} | Fetch a single project by key.
*ProjectsApi* | [**GetProjects**](docs/ProjectsApi.md#getprojects) | **Get** /projects | Returns a list of all projects in the account.
*ProjectsApi* | [**PatchProject**](docs/ProjectsApi.md#patchproject) | **Patch** /projects/{projectKey} | Modify a project by ID.
*ProjectsApi* | [**PostProject**](docs/ProjectsApi.md#postproject) | **Post** /projects | Create a new project with the given key and name.
*RootApi* | [**GetRoot**](docs/RootApi.md#getroot) | **Get** / | 
*TeamMembersApi* | [**DeleteMember**](docs/TeamMembersApi.md#deletemember) | **Delete** /members/{memberId} | Delete a team member by ID.
*TeamMembersApi* | [**GetMember**](docs/TeamMembersApi.md#getmember) | **Get** /members/{memberId} | Get a single team member by ID.
*TeamMembersApi* | [**GetMembers**](docs/TeamMembersApi.md#getmembers) | **Get** /members | Returns a list of all members in the account.
*TeamMembersApi* | [**PatchMember**](docs/TeamMembersApi.md#patchmember) | **Patch** /members/{memberId} | Modify a team member by ID.
*TeamMembersApi* | [**PostMembers**](docs/TeamMembersApi.md#postmembers) | **Post** /members | Invite new members.
*UserSegmentsApi* | [**DeleteUserSegment**](docs/UserSegmentsApi.md#deleteusersegment) | **Delete** /segments/{projectKey}/{environmentKey}/{userSegmentKey} | Delete a user segment.
*UserSegmentsApi* | [**GetUserSegment**](docs/UserSegmentsApi.md#getusersegment) | **Get** /segments/{projectKey}/{environmentKey}/{userSegmentKey} | Get a single user segment by key.
*UserSegmentsApi* | [**GetUserSegments**](docs/UserSegmentsApi.md#getusersegments) | **Get** /segments/{projectKey}/{environmentKey} | Get a list of all user segments in the given project.
*UserSegmentsApi* | [**PatchUserSegment**](docs/UserSegmentsApi.md#patchusersegment) | **Patch** /segments/{projectKey}/{environmentKey}/{userSegmentKey} | Perform a partial update to a user segment.
*UserSegmentsApi* | [**PostUserSegment**](docs/UserSegmentsApi.md#postusersegment) | **Post** /segments/{projectKey}/{environmentKey} | Creates a new user segment.
*UserSettingsApi* | [**GetUserFlagSetting**](docs/UserSettingsApi.md#getuserflagsetting) | **Get** /users/{projectKey}/{environmentKey}/{userKey}/flags/{featureFlagKey} | Fetch a single flag setting for a user by key.
*UserSettingsApi* | [**GetUserFlagSettings**](docs/UserSettingsApi.md#getuserflagsettings) | **Get** /users/{projectKey}/{environmentKey}/{userKey}/flags | Fetch a single flag setting for a user by key.
*UserSettingsApi* | [**PutFlagSetting**](docs/UserSettingsApi.md#putflagsetting) | **Put** /users/{projectKey}/{environmentKey}/{userKey}/flags/{featureFlagKey} | Specifically enable or disable a feature flag for a user based on their key.
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /users/{projectKey}/{environmentKey}/{userKey} | Delete a user by ID.
*UsersApi* | [**GetSearchUsers**](docs/UsersApi.md#getsearchusers) | **Get** /user-search/{projectKey}/{environmentKey} | Search users in LaunchDarkly based on their last active date, or a search query. It should not be used to enumerate all users in LaunchDarkly-- use the List users API resource.
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /users/{projectKey}/{environmentKey}/{userKey} | Get a user by key.
*UsersApi* | [**GetUsers**](docs/UsersApi.md#getusers) | **Get** /users/{projectKey}/{environmentKey} | List all users in the environment. Includes the total count of users. In each page, there will be up to &#39;limit&#39; users returned (default 20). This is useful for exporting all users in the system for further analysis. Paginated collections will include a next link containing a URL with the next set of elements in the collection.
*WebhooksApi* | [**DeleteWebhook**](docs/WebhooksApi.md#deletewebhook) | **Delete** /webhooks/{resourceId} | Delete a webhook by ID.
*WebhooksApi* | [**GetWebhook**](docs/WebhooksApi.md#getwebhook) | **Get** /webhooks/{resourceId} | Get a webhook by ID.
*WebhooksApi* | [**GetWebhooks**](docs/WebhooksApi.md#getwebhooks) | **Get** /webhooks | Fetch a list of all webhooks.
*WebhooksApi* | [**PatchWebhook**](docs/WebhooksApi.md#patchwebhook) | **Patch** /webhooks/{resourceId} | Modify a webhook by ID.
*WebhooksApi* | [**PostWebhook**](docs/WebhooksApi.md#postwebhook) | **Post** /webhooks | Create a webhook.


## Documentation For Models

 - [AuditLogEntries](docs/AuditLogEntries.md)
 - [AuditLogEntry](docs/AuditLogEntry.md)
 - [AuditLogEntryTarget](docs/AuditLogEntryTarget.md)
 - [Clause](docs/Clause.md)
 - [CustomProperty](docs/CustomProperty.md)
 - [CustomRole](docs/CustomRole.md)
 - [CustomRoleBody](docs/CustomRoleBody.md)
 - [CustomRoles](docs/CustomRoles.md)
 - [Environment](docs/Environment.md)
 - [EnvironmentPost](docs/EnvironmentPost.md)
 - [FeatureFlag](docs/FeatureFlag.md)
 - [FeatureFlagBody](docs/FeatureFlagBody.md)
 - [FeatureFlagConfig](docs/FeatureFlagConfig.md)
 - [FeatureFlagStatus](docs/FeatureFlagStatus.md)
 - [FeatureFlagStatuses](docs/FeatureFlagStatuses.md)
 - [FeatureFlags](docs/FeatureFlags.md)
 - [Id](docs/Id.md)
 - [Link](docs/Link.md)
 - [Links](docs/Links.md)
 - [Member](docs/Member.md)
 - [Members](docs/Members.md)
 - [MembersBody](docs/MembersBody.md)
 - [ModelFallthrough](docs/ModelFallthrough.md)
 - [PatchComment](docs/PatchComment.md)
 - [PatchOperation](docs/PatchOperation.md)
 - [Policy](docs/Policy.md)
 - [Prerequisite](docs/Prerequisite.md)
 - [Project](docs/Project.md)
 - [ProjectBody](docs/ProjectBody.md)
 - [Projects](docs/Projects.md)
 - [Role](docs/Role.md)
 - [Rollout](docs/Rollout.md)
 - [Rule](docs/Rule.md)
 - [Statement](docs/Statement.md)
 - [Statements](docs/Statements.md)
 - [Target](docs/Target.md)
 - [User](docs/User.md)
 - [UserFlagSetting](docs/UserFlagSetting.md)
 - [UserFlagSettings](docs/UserFlagSettings.md)
 - [UserRecord](docs/UserRecord.md)
 - [UserSegment](docs/UserSegment.md)
 - [UserSegmentBody](docs/UserSegmentBody.md)
 - [UserSegmentRule](docs/UserSegmentRule.md)
 - [UserSegments](docs/UserSegments.md)
 - [UserSettingsBody](docs/UserSettingsBody.md)
 - [Users](docs/Users.md)
 - [Variation](docs/Variation.md)
 - [Webhook](docs/Webhook.md)
 - [WebhookBody](docs/WebhookBody.md)
 - [Webhooks](docs/Webhooks.md)
 - [WeightedVariation](docs/WeightedVariation.md)


## Author

support@launchdarkly.com

## Sample Code

```
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/launchdarkly/api-client-go"
)

func main() {
	apiKey := os.Getenv("LD_API_KEY")
	if apiKey == "" {
		panic("LD_API_KEY env var was empty!")
	}
	client := ldapi.NewAPIClient(ldapi.NewConfiguration())
	ctx := context.WithValue(context.Background(), ldapi.ContextAPIKey, ldapi.APIKey{
		Key: apiKey,
	})

	// Create a multi-variate feature flag
	body := ldapi.FeatureFlagBody{
		Name: "Test Flag Go",
		Key:  "test-go",
		Variations: []ldapi.Variation{
			{Value: intfPtr([]interface{}{1, 2})},
			{Value: intfPtr([]interface{}{3, 4})},
			{Value: intfPtr([]interface{}{5})}}}
	flag, _, err := client.FeatureFlagsApi.PostFeatureFlag(ctx, "openapi", body, nil)
	if err != nil {
		panic(fmt.Errorf("create failed: %s", err))
	}
	fmt.Printf("Created flag: %+v\n", flag)
	// Clean up new flag
	defer func() {
		if _, err := client.FeatureFlagsApi.DeleteFeatureFlag(ctx, "openapi", body.Key); err != nil {
			panic(fmt.Errorf("delete failed: %s", err))
		}
	}()
}

func intfPtr(i interface{}) *interface{} {
	return &i
}
```
