# \StylesAPI

All URIs are relative to *https://api.figma.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFileStyles**](StylesAPI.md#GetFileStyles) | **Get** /v1/files/{file_key}/styles | Get file styles
[**GetStyle**](StylesAPI.md#GetStyle) | **Get** /v1/styles/{key} | Get style
[**GetTeamStyles**](StylesAPI.md#GetTeamStyles) | **Get** /v1/teams/{team_id}/styles | Get team styles



## GetFileStyles

> map[string]interface{} GetFileStyles(ctx, fileKey).Execute()

Get file styles



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
	fileKey := "fileKey_example" // string | File to list styles from. This must be a main file key, not a branch key, as it is not possible to publish from branches.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StylesAPI.GetFileStyles(context.Background(), fileKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StylesAPI.GetFileStyles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileStyles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StylesAPI.GetFileStyles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File to list styles from. This must be a main file key, not a branch key, as it is not possible to publish from branches. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileStylesRequest struct via the builder pattern


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


## GetStyle

> map[string]interface{} GetStyle(ctx, key).Execute()

Get style



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
	key := "key_example" // string | The unique identifier of the style.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StylesAPI.GetStyle(context.Background(), key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StylesAPI.GetStyle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStyle`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StylesAPI.GetStyle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | The unique identifier of the style. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStyleRequest struct via the builder pattern


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


## GetTeamStyles

> map[string]interface{} GetTeamStyles(ctx, teamId).PageSize(pageSize).After(after).Before(before).Execute()

Get team styles



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
	teamId := "teamId_example" // string | Id of the team to list styles from.
	pageSize := float32(8.14) // float32 | Number of items to return in a paged list of results. Defaults to 30. (optional) (default to 30)
	after := float32(8.14) // float32 | Cursor indicating which id after which to start retrieving styles for. Exclusive with before. The cursor value is an internally tracked integer that doesn't correspond to any Ids. (optional)
	before := float32(8.14) // float32 | Cursor indicating which id before which to start retrieving styles for. Exclusive with after. The cursor value is an internally tracked integer that doesn't correspond to any Ids. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StylesAPI.GetTeamStyles(context.Background(), teamId).PageSize(pageSize).After(after).Before(before).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StylesAPI.GetTeamStyles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamStyles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `StylesAPI.GetTeamStyles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | Id of the team to list styles from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamStylesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **float32** | Number of items to return in a paged list of results. Defaults to 30. | [default to 30]
 **after** | **float32** | Cursor indicating which id after which to start retrieving styles for. Exclusive with before. The cursor value is an internally tracked integer that doesn&#39;t correspond to any Ids. | 
 **before** | **float32** | Cursor indicating which id before which to start retrieving styles for. Exclusive with after. The cursor value is an internally tracked integer that doesn&#39;t correspond to any Ids. | 

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

