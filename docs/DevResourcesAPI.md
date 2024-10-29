# \DevResourcesAPI

All URIs are relative to *https://api.figma.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDevResource**](DevResourcesAPI.md#DeleteDevResource) | **Delete** /v1/files/{file_key}/dev_resources/{dev_resource_id} | Delete dev resource
[**GetDevResources**](DevResourcesAPI.md#GetDevResources) | **Get** /v1/files/{file_key}/dev_resources | Get dev resources
[**PostDevResources**](DevResourcesAPI.md#PostDevResources) | **Post** /v1/dev_resources | Create dev resources
[**PutDevResources**](DevResourcesAPI.md#PutDevResources) | **Put** /v1/dev_resources | Update dev resources



## DeleteDevResource

> DeleteDevResource(ctx, fileKey, devResourceId).Execute()

Delete dev resource



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
	fileKey := "fileKey_example" // string | The file to delete the dev resource from. This must be a main file key, not a branch key.
	devResourceId := "devResourceId_example" // string | The id of the dev resource to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevResourcesAPI.DeleteDevResource(context.Background(), fileKey, devResourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevResourcesAPI.DeleteDevResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | The file to delete the dev resource from. This must be a main file key, not a branch key. | 
**devResourceId** | **string** | The id of the dev resource to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDevResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDevResources

> map[string]interface{} GetDevResources(ctx, fileKey).NodeIds(nodeIds).Execute()

Get dev resources



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
	fileKey := "fileKey_example" // string | The file to get the dev resources from. This must be a main file key, not a branch key.
	nodeIds := "nodeIds_example" // string | Comma separated list of nodes that you care about in the document. If specified, only dev resources attached to these nodes will be returned. If not specified, all dev resources in the file will be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevResourcesAPI.GetDevResources(context.Background(), fileKey).NodeIds(nodeIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevResourcesAPI.GetDevResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDevResources`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevResourcesAPI.GetDevResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | The file to get the dev resources from. This must be a main file key, not a branch key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nodeIds** | **string** | Comma separated list of nodes that you care about in the document. If specified, only dev resources attached to these nodes will be returned. If not specified, all dev resources in the file will be returned. | 

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


## PostDevResources

> map[string]interface{} PostDevResources(ctx).PostDevResourcesRequest(postDevResourcesRequest).Execute()

Create dev resources



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
	postDevResourcesRequest := *openapiclient.NewPostDevResourcesRequest([]openapiclient.PostDevResourcesRequestDevResourcesInner{*openapiclient.NewPostDevResourcesRequestDevResourcesInner("Name_example", "Url_example", "FileKey_example", "NodeId_example")}) // PostDevResourcesRequest | A list of dev resources that you want to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevResourcesAPI.PostDevResources(context.Background()).PostDevResourcesRequest(postDevResourcesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevResourcesAPI.PostDevResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDevResources`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevResourcesAPI.PostDevResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostDevResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postDevResourcesRequest** | [**PostDevResourcesRequest**](PostDevResourcesRequest.md) | A list of dev resources that you want to create. | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDevResources

> map[string]interface{} PutDevResources(ctx).PutDevResourcesRequest(putDevResourcesRequest).Execute()

Update dev resources



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
	putDevResourcesRequest := *openapiclient.NewPutDevResourcesRequest([]openapiclient.PutDevResourcesRequestDevResourcesInner{*openapiclient.NewPutDevResourcesRequestDevResourcesInner("Id_example")}) // PutDevResourcesRequest | A list of dev resources that you want to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevResourcesAPI.PutDevResources(context.Background()).PutDevResourcesRequest(putDevResourcesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevResourcesAPI.PutDevResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDevResources`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevResourcesAPI.PutDevResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutDevResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putDevResourcesRequest** | [**PutDevResourcesRequest**](PutDevResourcesRequest.md) | A list of dev resources that you want to update. | 

### Return type

**map[string]interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

