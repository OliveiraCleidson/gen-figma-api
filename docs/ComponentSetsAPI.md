# \ComponentSetsAPI

All URIs are relative to *https://api.figma.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComponentSet**](ComponentSetsAPI.md#GetComponentSet) | **Get** /v1/component_sets/{key} | Get component set
[**GetFileComponentSets**](ComponentSetsAPI.md#GetFileComponentSets) | **Get** /v1/files/{file_key}/component_sets | Get file component sets
[**GetTeamComponentSets**](ComponentSetsAPI.md#GetTeamComponentSets) | **Get** /v1/teams/{team_id}/component_sets | Get team component sets



## GetComponentSet

> map[string]interface{} GetComponentSet(ctx, key).Execute()

Get component set



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/oliveiracleidson/gen-figma-api"
)

func main() {
	key := "key_example" // string | The unique identifier of the component set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentSetsAPI.GetComponentSet(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentSetsAPI.GetComponentSet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponentSet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ComponentSetsAPI.GetComponentSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The unique identifier of the component set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileComponentSets

> map[string]interface{} GetFileComponentSets(ctx, fileKey).Execute()

Get file component sets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/oliveiracleidson/gen-figma-api"
)

func main() {
	fileKey := "fileKey_example" // string | File to list component sets from. This must be a main file key, not a branch key, as it is not possible to publish from branches.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentSetsAPI.GetFileComponentSets(context.Background(), fileKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentSetsAPI.GetFileComponentSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileComponentSets`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ComponentSetsAPI.GetFileComponentSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File to list component sets from. This must be a main file key, not a branch key, as it is not possible to publish from branches. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileComponentSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeamComponentSets

> map[string]interface{} GetTeamComponentSets(ctx, teamId).PageSize(pageSize).After(after).Before(before).Execute()

Get team component sets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/oliveiracleidson/gen-figma-api"
)

func main() {
	teamId := "teamId_example" // string | Id of the team to list component sets from.
	pageSize := float32(8.14) // float32 | Number of items to return in a paged list of results. Defaults to 30. (optional) (default to 30)
	after := float32(8.14) // float32 | Cursor indicating which id after which to start retrieving component sets for. Exclusive with before. The cursor value is an internally tracked integer that doesn't correspond to any Ids. (optional)
	before := float32(8.14) // float32 | Cursor indicating which id before which to start retrieving component sets for. Exclusive with after. The cursor value is an internally tracked integer that doesn't correspond to any Ids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentSetsAPI.GetTeamComponentSets(context.Background(), teamId).PageSize(pageSize).After(after).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentSetsAPI.GetTeamComponentSets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamComponentSets`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ComponentSetsAPI.GetTeamComponentSets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | Id of the team to list component sets from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamComponentSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **float32** | Number of items to return in a paged list of results. Defaults to 30. | [default to 30]
 **after** | **float32** | Cursor indicating which id after which to start retrieving component sets for. Exclusive with before. The cursor value is an internally tracked integer that doesn&#39;t correspond to any Ids. | 
 **before** | **float32** | Cursor indicating which id before which to start retrieving component sets for. Exclusive with after. The cursor value is an internally tracked integer that doesn&#39;t correspond to any Ids. | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

