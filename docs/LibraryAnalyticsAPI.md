# \LibraryAnalyticsAPI

All URIs are relative to *https://api.figma.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLibraryAnalyticsActions**](LibraryAnalyticsAPI.md#GetLibraryAnalyticsActions) | **Get** /v1/analytics/libraries/{file_key}/actions | Get library analytics action data.
[**GetLibraryAnalyticsUsages**](LibraryAnalyticsAPI.md#GetLibraryAnalyticsUsages) | **Get** /v1/analytics/libraries/{file_key}/usages | Get library analytics usage data.



## GetLibraryAnalyticsActions

> map[string]interface{} GetLibraryAnalyticsActions(ctx, fileKey).GroupBy(groupBy).Cursor(cursor).StartDate(startDate).EndDate(endDate).Order(order).Execute()

Get library analytics action data.



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
	fileKey := "fileKey_example" // string | File key of the library to fetch analytics data for.
	groupBy := "groupBy_example" // string | A dimension to group returned analytics data by. Accepts \"component\" or \"team\".
	cursor := "cursor_example" // string | Cursor indicating what page of data to fetch. Obtained from prior API call. (optional)
	startDate := "startDate_example" // string | ISO 8601 date string (YYYY-MM-DD) of the earliest week to include. Dates are rounded back to the nearest start of a week. Defaults to one year prior. (optional)
	endDate := "endDate_example" // string | ISO 8601 date string (YYYY-MM-DD) of the latest week to include. Dates are rounded forward to the nearest end of a week. Defaults to the latest computed week. (optional)
	order := "order_example" // string | How to order response rows by week. This param can be either \"asc\" or \"desc\" (default). (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAnalyticsAPI.GetLibraryAnalyticsActions(context.Background(), fileKey).GroupBy(groupBy).Cursor(cursor).StartDate(startDate).EndDate(endDate).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAnalyticsAPI.GetLibraryAnalyticsActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryAnalyticsActions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LibraryAnalyticsAPI.GetLibraryAnalyticsActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File key of the library to fetch analytics data for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryAnalyticsActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupBy** | **string** | A dimension to group returned analytics data by. Accepts \&quot;component\&quot; or \&quot;team\&quot;. | 
 **cursor** | **string** | Cursor indicating what page of data to fetch. Obtained from prior API call. | 
 **startDate** | **string** | ISO 8601 date string (YYYY-MM-DD) of the earliest week to include. Dates are rounded back to the nearest start of a week. Defaults to one year prior. | 
 **endDate** | **string** | ISO 8601 date string (YYYY-MM-DD) of the latest week to include. Dates are rounded forward to the nearest end of a week. Defaults to the latest computed week. | 
 **order** | **string** | How to order response rows by week. This param can be either \&quot;asc\&quot; or \&quot;desc\&quot; (default). | [default to &quot;desc&quot;]

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


## GetLibraryAnalyticsUsages

> map[string]interface{} GetLibraryAnalyticsUsages(ctx, fileKey).GroupBy(groupBy).Cursor(cursor).Order(order).Execute()

Get library analytics usage data.



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
	fileKey := "fileKey_example" // string | File key of the library to fetch analytics data for.
	groupBy := "groupBy_example" // string | A dimension to group returned analytics data by. Accepts \"component\" or \"file\".
	cursor := "cursor_example" // string | Cursor indicating what page of data to fetch. Obtained from prior API call. (optional)
	order := "order_example" // string | How to order response rows by number of instances. This param can be either \"asc\" or \"desc\" (default). (optional) (default to "desc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LibraryAnalyticsAPI.GetLibraryAnalyticsUsages(context.Background(), fileKey).GroupBy(groupBy).Cursor(cursor).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LibraryAnalyticsAPI.GetLibraryAnalyticsUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLibraryAnalyticsUsages`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LibraryAnalyticsAPI.GetLibraryAnalyticsUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File key of the library to fetch analytics data for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLibraryAnalyticsUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **groupBy** | **string** | A dimension to group returned analytics data by. Accepts \&quot;component\&quot; or \&quot;file\&quot;. | 
 **cursor** | **string** | Cursor indicating what page of data to fetch. Obtained from prior API call. | 
 **order** | **string** | How to order response rows by number of instances. This param can be either \&quot;asc\&quot; or \&quot;desc\&quot; (default). | [default to &quot;desc&quot;]

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

