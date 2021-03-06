# \FeatureFlagsApi

All URIs are relative to *https://app.launchdarkly.com/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CopyFeatureFlag**](FeatureFlagsApi.md#CopyFeatureFlag) | **Post** /flags/{projectKey}/{featureFlagKey}/copy | Copies the feature flag configuration from one environment to the same feature flag in another environment.
[**DeleteFeatureFlag**](FeatureFlagsApi.md#DeleteFeatureFlag) | **Delete** /flags/{projectKey}/{featureFlagKey} | Delete a feature flag in all environments. Be careful-- only delete feature flags that are no longer being used by your application.
[**GetExpiringUserTargets**](FeatureFlagsApi.md#GetExpiringUserTargets) | **Get** /flags/{projectKey}/{featureFlagKey}/expiring-user-targets/{environmentKey} | Get expiring user targets for feature flag
[**GetFeatureFlag**](FeatureFlagsApi.md#GetFeatureFlag) | **Get** /flags/{projectKey}/{featureFlagKey} | Get a single feature flag by key.
[**GetFeatureFlagChangeRequest**](FeatureFlagsApi.md#GetFeatureFlagChangeRequest) | **Get** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{featureFlagChangeRequestId} | Get a single change request for a feature flag
[**GetFeatureFlagChangeRequests**](FeatureFlagsApi.md#GetFeatureFlagChangeRequests) | **Get** /{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests | Get all change requests for a feature flag
[**GetFeatureFlagStatus**](FeatureFlagsApi.md#GetFeatureFlagStatus) | **Get** /flag-statuses/{projectKey}/{environmentKey}/{featureFlagKey} | Get the status for a particular feature flag.
[**GetFeatureFlagStatusAcrossEnvironments**](FeatureFlagsApi.md#GetFeatureFlagStatusAcrossEnvironments) | **Get** /flag-status/{projectKey}/{featureFlagKey} | Get the status for a particular feature flag across environments
[**GetFeatureFlagStatuses**](FeatureFlagsApi.md#GetFeatureFlagStatuses) | **Get** /flag-statuses/{projectKey}/{environmentKey} | Get a list of statuses for all feature flags. The status includes the last time the feature flag was requested, as well as the state of the flag.
[**GetFeatureFlags**](FeatureFlagsApi.md#GetFeatureFlags) | **Get** /flags/{projectKey} | Get a list of all features in the given project.
[**PatchExpiringUserTargets**](FeatureFlagsApi.md#PatchExpiringUserTargets) | **Patch** /flags/{projectKey}/{featureFlagKey}/expiring-user-targets/{environmentKey} | Update, add, or delete expiring user targets on feature flag
[**PatchFeatureFlag**](FeatureFlagsApi.md#PatchFeatureFlag) | **Patch** /flags/{projectKey}/{featureFlagKey} | Perform a partial update to a feature.
[**PostApplyFeatureFlagChangeRequest**](FeatureFlagsApi.md#PostApplyFeatureFlagChangeRequest) | **Post** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{featureFlagChangeRequestId}/apply | Apply change request for a feature flag
[**PostFeatureFlag**](FeatureFlagsApi.md#PostFeatureFlag) | **Post** /flags/{projectKey} | Creates a new feature flag.
[**PostFeatureFlagChangeRequest**](FeatureFlagsApi.md#PostFeatureFlagChangeRequest) | **Post** /{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests | 
[**PostReviewFeatureFlagChangeRequest**](FeatureFlagsApi.md#PostReviewFeatureFlagChangeRequest) | **Post** /projects/{projectKey}/flags/{featureFlagKey}/environments/{environmentKey}/approval-requests/{featureFlagChangeRequestId}/review | Review change request for a feature flag


# **CopyFeatureFlag**
> FeatureFlag CopyFeatureFlag(ctx, projectKey, featureFlagKey, featureFlagCopyBody)
Copies the feature flag configuration from one environment to the same feature flag in another environment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **featureFlagCopyBody** | [**FeatureFlagCopyBody**](FeatureFlagCopyBody.md)| Copy feature flag configurations between environments. | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFeatureFlag**
> DeleteFeatureFlag(ctx, projectKey, featureFlagKey)
Delete a feature flag in all environments. Be careful-- only delete feature flags that are no longer being used by your application.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 

### Return type

 (empty response body)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetExpiringUserTargets**
> UserTargetingExpirationForFlags GetExpiringUserTargets(ctx, projectKey, environmentKey, featureFlagKey)
Get expiring user targets for feature flag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 

### Return type

[**UserTargetingExpirationForFlags**](UserTargetingExpirationForFlags.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlag**
> FeatureFlag GetFeatureFlag(ctx, projectKey, featureFlagKey, optional)
Get a single feature flag by key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
 **optional** | ***GetFeatureFlagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFeatureFlagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **env** | [**optional.Interface of []string**](string.md)| By default, each feature will include configurations for each environment. You can filter environments with the env query parameter. For example, setting env&#x3D;[\&quot;production\&quot;] will restrict the returned configurations to just your production environment. | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlagChangeRequest**
> FeatureFlagChangeRequests GetFeatureFlagChangeRequest(ctx, projectKey, featureFlagKey, environmentKey, featureFlagChangeRequestId)
Get a single change request for a feature flag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **featureFlagChangeRequestId** | **string**| The feature flag change request ID | 

### Return type

[**FeatureFlagChangeRequests**](FeatureFlagChangeRequests.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlagChangeRequests**
> FeatureFlagChangeRequests GetFeatureFlagChangeRequests(ctx, projectKey, featureFlagKey, environmentKey)
Get all change requests for a feature flag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 

### Return type

[**FeatureFlagChangeRequests**](FeatureFlagChangeRequests.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlagStatus**
> FeatureFlagStatus GetFeatureFlagStatus(ctx, projectKey, environmentKey, featureFlagKey)
Get the status for a particular feature flag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 

### Return type

[**FeatureFlagStatus**](FeatureFlagStatus.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlagStatusAcrossEnvironments**
> FeatureFlagStatusAcrossEnvironments GetFeatureFlagStatusAcrossEnvironments(ctx, projectKey, featureFlagKey)
Get the status for a particular feature flag across environments

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 

### Return type

[**FeatureFlagStatusAcrossEnvironments**](FeatureFlagStatusAcrossEnvironments.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlagStatuses**
> FeatureFlagStatuses GetFeatureFlagStatuses(ctx, projectKey, environmentKey)
Get a list of statuses for all feature flags. The status includes the last time the feature flag was requested, as well as the state of the flag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 

### Return type

[**FeatureFlagStatuses**](FeatureFlagStatuses.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetFeatureFlags**
> FeatureFlags GetFeatureFlags(ctx, projectKey, optional)
Get a list of all features in the given project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
 **optional** | ***GetFeatureFlagsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetFeatureFlagsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **env** | [**optional.Interface of []string**](string.md)| By default, each feature will include configurations for each environment. You can filter environments with the env query parameter. For example, setting env&#x3D;[\&quot;production\&quot;] will restrict the returned configurations to just your production environment. | 
 **summary** | **optional.Bool**| By default in api version &gt;&#x3D; 1, flags will _not_ include their list of prerequisites, targets or rules.  Set summary&#x3D;0 to include these fields for each flag returned. | 
 **archived** | **optional.Bool**| When set to 1, only archived flags will be included in the list of flags returned.  By default, archived flags are not included in the list of flags. | 
 **limit** | **optional.Float32**| The number of objects to return. Defaults to -1, which returns everything. | 
 **offset** | **optional.Float32**| Where to start in the list. This is for use with pagination. For example, an offset of 10 would skip the first 10 items and then return the next limit items. | 
 **filter** | **optional.String**| A comma-separated list of filters. Each filter is of the form field:value. | 
 **sort** | **optional.String**| A comma-separated list of fields to sort by. A field prefixed by a - will be sorted in descending order. | 
 **tag** | **optional.String**| Filter by tag. A tag can be used to group flags across projects. | 

### Return type

[**FeatureFlags**](FeatureFlags.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchExpiringUserTargets**
> UserTargetingExpirationForFlags PatchExpiringUserTargets(ctx, projectKey, environmentKey, featureFlagKey, semanticPatchWithComment)
Update, add, or delete expiring user targets on feature flag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **semanticPatchWithComment** | [**interface{}**](interface{}.md)| Requires a Semantic Patch representation of the desired changes to the resource. &#39;https://apidocs.launchdarkly.com/reference#updates-via-semantic-patches&#39;. The addition of comments is also supported. | 

### Return type

[**UserTargetingExpirationForFlags**](UserTargetingExpirationForFlags.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchFeatureFlag**
> FeatureFlag PatchFeatureFlag(ctx, projectKey, featureFlagKey, patchComment)
Perform a partial update to a feature.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **patchComment** | [**PatchComment**](PatchComment.md)| Requires a JSON Patch representation of the desired changes to the project, and an optional comment. &#39;http://jsonpatch.com/&#39; Feature flag patches also support JSON Merge Patch format. &#39;https://tools.ietf.org/html/rfc7386&#39; The addition of comments is also supported. | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostApplyFeatureFlagChangeRequest**
> FeatureFlagChangeRequests PostApplyFeatureFlagChangeRequest(ctx, projectKey, featureFlagKey, environmentKey, featureFlagChangeRequestId, featureFlagChangeRequestApplyConfigBody)
Apply change request for a feature flag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **featureFlagChangeRequestId** | **string**| The feature flag change request ID | 
  **featureFlagChangeRequestApplyConfigBody** | [**FeatureFlagChangeRequestApplyConfigBody**](FeatureFlagChangeRequestApplyConfigBody.md)| Apply a new feature flag change request | 

### Return type

[**FeatureFlagChangeRequests**](FeatureFlagChangeRequests.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFeatureFlag**
> FeatureFlag PostFeatureFlag(ctx, projectKey, featureFlagBody, optional)
Creates a new feature flag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagBody** | [**FeatureFlagBody**](FeatureFlagBody.md)| Create a new feature flag. | 
 **optional** | ***PostFeatureFlagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostFeatureFlagOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **clone** | **optional.String**| The key of the feature flag to be cloned. The key identifies the flag in your code.  For example, setting clone&#x3D;flagKey will copy the full targeting configuration for all environments (including on/off state) from the original flag to the new flag. | 

### Return type

[**FeatureFlag**](FeatureFlag.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostFeatureFlagChangeRequest**
> FeatureFlagChangeRequest PostFeatureFlagChangeRequest(ctx, projectKey, featureFlagKey, environmentKey, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
 **optional** | ***PostFeatureFlagChangeRequestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostFeatureFlagChangeRequestOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **featureFlagChangeRequestConfigBody** | [**optional.Interface of FeatureFlagChangeRequestConfigBody**](FeatureFlagChangeRequestConfigBody.md)| Create a new feature flag change request | 

### Return type

[**FeatureFlagChangeRequest**](FeatureFlagChangeRequest.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostReviewFeatureFlagChangeRequest**
> FeatureFlagChangeRequests PostReviewFeatureFlagChangeRequest(ctx, projectKey, featureFlagKey, environmentKey, featureFlagChangeRequestId, featureFlagChangeRequestReviewConfigBody)
Review change request for a feature flag

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectKey** | **string**| The project key, used to tie the flags together under one project so they can be managed together. | 
  **featureFlagKey** | **string**| The feature flag&#39;s key. The key identifies the flag in your code. | 
  **environmentKey** | **string**| The environment key, used to tie together flag configuration and users under one environment so they can be managed together. | 
  **featureFlagChangeRequestId** | **string**| The feature flag change request ID | 
  **featureFlagChangeRequestReviewConfigBody** | [**FeatureFlagChangeRequestReviewConfigBody**](FeatureFlagChangeRequestReviewConfigBody.md)| Review a feature flag change request | 

### Return type

[**FeatureFlagChangeRequests**](FeatureFlagChangeRequests.md)

### Authorization

[Token](../README.md#Token)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

