# \ComponentsAPI

All URIs are relative to *https://api.figma.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetComponent**](ComponentsAPI.md#GetComponent) | **Get** /v1/components/{key} | Get component
[**GetFileComponents**](ComponentsAPI.md#GetFileComponents) | **Get** /v1/files/{file_key}/components | Get file components
[**GetTeamComponents**](ComponentsAPI.md#GetTeamComponents) | **Get** /v1/teams/{team_id}/components | Get team components



## GetComponent

> map[string]interface{} GetComponent(ctx, key).Execute()

Get component



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
	key := "key_example" // string | The unique identifier of the component.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.GetComponent(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponent`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The unique identifier of the component. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentRequest struct via the builder pattern


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


## GetFileComponents

> map[string]interface{} GetFileComponents(ctx, fileKey).Execute()

Get file components



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
	fileKey := "fileKey_example" // string | File to list components from. This must be a main file key, not a branch key, as it is not possible to publish from branches.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.GetFileComponents(context.Background(), fileKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetFileComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileComponents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetFileComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File to list components from. This must be a main file key, not a branch key, as it is not possible to publish from branches. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileComponentsRequest struct via the builder pattern


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


## GetTeamComponents

> map[string]interface{} GetTeamComponents(ctx, teamId).PageSize(pageSize).After(after).Before(before).Execute()

Get team components



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
	teamId := "teamId_example" // string | Id of the team to list components from.
	pageSize := float32(8.14) // float32 | Number of items to return in a paged list of results. Defaults to 30. (optional) (default to 30)
	after := float32(8.14) // float32 | Cursor indicating which id after which to start retrieving components for. Exclusive with before. The cursor value is an internally tracked integer that doesn't correspond to any Ids. (optional)
	before := float32(8.14) // float32 | Cursor indicating which id before which to start retrieving components for. Exclusive with after. The cursor value is an internally tracked integer that doesn't correspond to any Ids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComponentsAPI.GetTeamComponents(context.Background(), teamId).PageSize(pageSize).After(after).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComponentsAPI.GetTeamComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamComponents`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ComponentsAPI.GetTeamComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | Id of the team to list components from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **float32** | Number of items to return in a paged list of results. Defaults to 30. | [default to 30]
 **after** | **float32** | Cursor indicating which id after which to start retrieving components for. Exclusive with before. The cursor value is an internally tracked integer that doesn&#39;t correspond to any Ids. | 
 **before** | **float32** | Cursor indicating which id before which to start retrieving components for. Exclusive with after. The cursor value is an internally tracked integer that doesn&#39;t correspond to any Ids. | 

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

