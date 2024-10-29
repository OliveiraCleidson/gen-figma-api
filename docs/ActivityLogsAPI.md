# \ActivityLogsAPI

All URIs are relative to *https://api.figma.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetActivityLogs**](ActivityLogsAPI.md#GetActivityLogs) | **Get** /v1/activity_logs | Get activity logs



## GetActivityLogs

> map[string]interface{} GetActivityLogs(ctx).Events(events).StartTime(startTime).EndTime(endTime).Limit(limit).Order(order).Execute()

Get activity logs



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
	events := "events_example" // string | Event type(s) to include in the response. Can have multiple values separated by comma. All events are returned by default. (optional)
	startTime := float32(8.14) // float32 | Unix timestamp of the least recent event to include. This param defaults to one year ago if unspecified. Events prior to one year ago are not available. (optional)
	endTime := float32(8.14) // float32 | Unix timestamp of the most recent event to include. This param defaults to the current timestamp if unspecified. (optional)
	limit := float32(8.14) // float32 | Maximum number of events to return. This param defaults to 1000 if unspecified. (optional)
	order := "order_example" // string | Event order by timestamp. This param can be either \"asc\" (default) or \"desc\". (optional) (default to "asc")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ActivityLogsAPI.GetActivityLogs(context.Background()).Events(events).StartTime(startTime).EndTime(endTime).Limit(limit).Order(order).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ActivityLogsAPI.GetActivityLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetActivityLogs`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ActivityLogsAPI.GetActivityLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetActivityLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **events** | **string** | Event type(s) to include in the response. Can have multiple values separated by comma. All events are returned by default. | 
 **startTime** | **float32** | Unix timestamp of the least recent event to include. This param defaults to one year ago if unspecified. Events prior to one year ago are not available. | 
 **endTime** | **float32** | Unix timestamp of the most recent event to include. This param defaults to the current timestamp if unspecified. | 
 **limit** | **float32** | Maximum number of events to return. This param defaults to 1000 if unspecified. | 
 **order** | **string** | Event order by timestamp. This param can be either \&quot;asc\&quot; (default) or \&quot;desc\&quot;. | [default to &quot;asc&quot;]

### Return type

**map[string]interface{}**

### Authorization

[OrgOAuth2](../README.md#OrgOAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

