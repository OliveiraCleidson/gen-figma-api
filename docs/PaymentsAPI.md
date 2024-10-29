# \PaymentsAPI

All URIs are relative to *https://api.figma.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayments**](PaymentsAPI.md#GetPayments) | **Get** /v1/payments | Get payments



## GetPayments

> map[string]interface{} GetPayments(ctx).PluginPaymentToken(pluginPaymentToken).UserId(userId).CommunityFileId(communityFileId).PluginId(pluginId).WidgetId(widgetId).Execute()

Get payments



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
	pluginPaymentToken := "pluginPaymentToken_example" // string | Short-lived token returned from \"getPluginPaymentTokenAsync\" in the plugin payments API and used to authenticate to this endpoint. Read more about generating this token through \"Calling the Payments REST API from a plugin or widget\" below. (optional)
	userId := float32(8.14) // float32 | The ID of the user to query payment information about. You can get the user ID by having the user OAuth2 to the Figma REST API. (optional)
	communityFileId := float32(8.14) // float32 | The ID of the Community file to query a user's payment information on. You can get the Community file ID from the file's Community page (look for the number after \"file/\" in the URL). Provide exactly one of \"community_file_id\", \"plugin_id\", or \"widget_id\". (optional)
	pluginId := float32(8.14) // float32 | The ID of the plugin to query a user's payment information on. You can get the plugin ID from the plugin's manifest, or from the plugin's Community page (look for the number after \"plugin/\" in the URL). Provide exactly one of \"community_file_id\", \"plugin_id\", or \"widget_id\". (optional)
	widgetId := float32(8.14) // float32 | The ID of the widget to query a user's payment information on. You can get the widget ID from the widget's manifest, or from the widget's Community page (look for the number after \"widget/\" in the URL). Provide exactly one of \"community_file_id\", \"plugin_id\", or \"widget_id\". (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentsAPI.GetPayments(context.Background()).PluginPaymentToken(pluginPaymentToken).UserId(userId).CommunityFileId(communityFileId).PluginId(pluginId).WidgetId(widgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentsAPI.GetPayments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPayments`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PaymentsAPI.GetPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pluginPaymentToken** | **string** | Short-lived token returned from \&quot;getPluginPaymentTokenAsync\&quot; in the plugin payments API and used to authenticate to this endpoint. Read more about generating this token through \&quot;Calling the Payments REST API from a plugin or widget\&quot; below. | 
 **userId** | **float32** | The ID of the user to query payment information about. You can get the user ID by having the user OAuth2 to the Figma REST API. | 
 **communityFileId** | **float32** | The ID of the Community file to query a user&#39;s payment information on. You can get the Community file ID from the file&#39;s Community page (look for the number after \&quot;file/\&quot; in the URL). Provide exactly one of \&quot;community_file_id\&quot;, \&quot;plugin_id\&quot;, or \&quot;widget_id\&quot;. | 
 **pluginId** | **float32** | The ID of the plugin to query a user&#39;s payment information on. You can get the plugin ID from the plugin&#39;s manifest, or from the plugin&#39;s Community page (look for the number after \&quot;plugin/\&quot; in the URL). Provide exactly one of \&quot;community_file_id\&quot;, \&quot;plugin_id\&quot;, or \&quot;widget_id\&quot;. | 
 **widgetId** | **float32** | The ID of the widget to query a user&#39;s payment information on. You can get the widget ID from the widget&#39;s manifest, or from the widget&#39;s Community page (look for the number after \&quot;widget/\&quot; in the URL). Provide exactly one of \&quot;community_file_id\&quot;, \&quot;plugin_id\&quot;, or \&quot;widget_id\&quot;. | 

### Return type

**map[string]interface{}**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

