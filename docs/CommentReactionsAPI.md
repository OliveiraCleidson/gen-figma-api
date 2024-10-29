# \CommentReactionsAPI

All URIs are relative to *https://api.figma.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCommentReaction**](CommentReactionsAPI.md#DeleteCommentReaction) | **Delete** /v1/files/{file_key}/comments/{comment_id}/reactions | Delete a reaction
[**GetCommentReactions**](CommentReactionsAPI.md#GetCommentReactions) | **Get** /v1/files/{file_key}/comments/{comment_id}/reactions | Get reactions for a comment
[**PostCommentReaction**](CommentReactionsAPI.md#PostCommentReaction) | **Post** /v1/files/{file_key}/comments/{comment_id}/reactions | Add a reaction to a comment



## DeleteCommentReaction

> map[string]interface{} DeleteCommentReaction(ctx, fileKey, commentId).Emoji(emoji).Execute()

Delete a reaction



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
	fileKey := "fileKey_example" // string | File to delete comment reaction from. This can be a file key or branch key. Use `GET /v1/files/:key` with the `branch_data` query param to get the branch key.
	commentId := "commentId_example" // string | ID of comment to delete reaction from.
	emoji := "emoji_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentReactionsAPI.DeleteCommentReaction(context.Background(), fileKey, commentId).Emoji(emoji).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentReactionsAPI.DeleteCommentReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCommentReaction`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommentReactionsAPI.DeleteCommentReaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File to delete comment reaction from. This can be a file key or branch key. Use &#x60;GET /v1/files/:key&#x60; with the &#x60;branch_data&#x60; query param to get the branch key. | 
**commentId** | **string** | ID of comment to delete reaction from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **emoji** | **string** |  | 

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


## GetCommentReactions

> map[string]interface{} GetCommentReactions(ctx, fileKey, commentId).Cursor(cursor).Execute()

Get reactions for a comment



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
	fileKey := "fileKey_example" // string | File to get comment containing reactions from. This can be a file key or branch key. Use `GET /v1/files/:key` with the `branch_data` query param to get the branch key.
	commentId := "commentId_example" // string | ID of comment to get reactions from.
	cursor := "cursor_example" // string | Cursor for pagination, retrieved from the response of the previous call. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentReactionsAPI.GetCommentReactions(context.Background(), fileKey, commentId).Cursor(cursor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentReactionsAPI.GetCommentReactions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentReactions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommentReactionsAPI.GetCommentReactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File to get comment containing reactions from. This can be a file key or branch key. Use &#x60;GET /v1/files/:key&#x60; with the &#x60;branch_data&#x60; query param to get the branch key. | 
**commentId** | **string** | ID of comment to get reactions from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentReactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **cursor** | **string** | Cursor for pagination, retrieved from the response of the previous call. | 

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


## PostCommentReaction

> map[string]interface{} PostCommentReaction(ctx, fileKey, commentId).PostCommentReactionRequest(postCommentReactionRequest).Execute()

Add a reaction to a comment



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
	fileKey := "fileKey_example" // string | File to post comment reactions to. This can be a file key or branch key. Use `GET /v1/files/:key` with the `branch_data` query param to get the branch key.
	commentId := "commentId_example" // string | ID of comment to react to.
	postCommentReactionRequest := *openapiclient.NewPostCommentReactionRequest("Emoji_example") // PostCommentReactionRequest | Reaction to post.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CommentReactionsAPI.PostCommentReaction(context.Background(), fileKey, commentId).PostCommentReactionRequest(postCommentReactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CommentReactionsAPI.PostCommentReaction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCommentReaction`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CommentReactionsAPI.PostCommentReaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File to post comment reactions to. This can be a file key or branch key. Use &#x60;GET /v1/files/:key&#x60; with the &#x60;branch_data&#x60; query param to get the branch key. | 
**commentId** | **string** | ID of comment to react to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCommentReactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postCommentReactionRequest** | [**PostCommentReactionRequest**](PostCommentReactionRequest.md) | Reaction to post. | 

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

