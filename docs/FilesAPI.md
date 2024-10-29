# \FilesAPI

All URIs are relative to *https://api.figma.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFile**](FilesAPI.md#GetFile) | **Get** /v1/files/{file_key} | Get file JSON
[**GetFileNodes**](FilesAPI.md#GetFileNodes) | **Get** /v1/files/{file_key}/nodes | Get file JSON for specific nodes
[**GetFileVersions**](FilesAPI.md#GetFileVersions) | **Get** /v1/files/{file_key}/versions | Get versions of a file
[**GetImageFills**](FilesAPI.md#GetImageFills) | **Get** /v1/files/{file_key}/images | Get image fills
[**GetImages**](FilesAPI.md#GetImages) | **Get** /v1/images/{file_key} | Render images of file nodes



## GetFile

> map[string]interface{} GetFile(ctx, fileKey).Version(version).Ids(ids).Depth(depth).Geometry(geometry).PluginData(pluginData).BranchData(branchData).Execute()

Get file JSON



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
	fileKey := "fileKey_example" // string | File to export JSON from. This can be a file key or branch key. Use `GET /v1/files/:key` with the `branch_data` query param to get the branch key.
	version := "version_example" // string | A specific version ID to get. Omitting this will get the current version of the file. (optional)
	ids := "ids_example" // string | Comma separated list of nodes that you care about in the document. If specified, only a subset of the document will be returned corresponding to the nodes listed, their children, and everything between the root node and the listed nodes.  Note: There may be other nodes included in the returned JSON that are outside the ancestor chains of the desired nodes. The response may also include dependencies of anything in the nodes' subtrees. For example, if a node subtree contains an instance of a local component that lives elsewhere in that file, that component and its ancestor chain will also be included.  For historical reasons, top-level canvas nodes are always returned, regardless of whether they are listed in the `ids` parameter. This quirk may be removed in a future version of the API. (optional)
	depth := float32(8.14) // float32 | Positive integer representing how deep into the document tree to traverse. For example, setting this to 1 returns only Pages, setting it to 2 returns Pages and all top level objects on each page. Not setting this parameter returns all nodes. (optional)
	geometry := "geometry_example" // string | Set to \"paths\" to export vector data. (optional)
	pluginData := "pluginData_example" // string | A comma separated list of plugin IDs and/or the string \"shared\". Any data present in the document written by those plugins will be included in the result in the `pluginData` and `sharedPluginData` properties. (optional)
	branchData := true // bool | Returns branch metadata for the requested file. If the file is a branch, the main file's key will be included in the returned response. If the file has branches, their metadata will be included in the returned response. Default: false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFile(context.Background(), fileKey).Version(version).Ids(ids).Depth(depth).Geometry(geometry).PluginData(pluginData).BranchData(branchData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFile`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File to export JSON from. This can be a file key or branch key. Use &#x60;GET /v1/files/:key&#x60; with the &#x60;branch_data&#x60; query param to get the branch key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **string** | A specific version ID to get. Omitting this will get the current version of the file. | 
 **ids** | **string** | Comma separated list of nodes that you care about in the document. If specified, only a subset of the document will be returned corresponding to the nodes listed, their children, and everything between the root node and the listed nodes.  Note: There may be other nodes included in the returned JSON that are outside the ancestor chains of the desired nodes. The response may also include dependencies of anything in the nodes&#39; subtrees. For example, if a node subtree contains an instance of a local component that lives elsewhere in that file, that component and its ancestor chain will also be included.  For historical reasons, top-level canvas nodes are always returned, regardless of whether they are listed in the &#x60;ids&#x60; parameter. This quirk may be removed in a future version of the API. | 
 **depth** | **float32** | Positive integer representing how deep into the document tree to traverse. For example, setting this to 1 returns only Pages, setting it to 2 returns Pages and all top level objects on each page. Not setting this parameter returns all nodes. | 
 **geometry** | **string** | Set to \&quot;paths\&quot; to export vector data. | 
 **pluginData** | **string** | A comma separated list of plugin IDs and/or the string \&quot;shared\&quot;. Any data present in the document written by those plugins will be included in the result in the &#x60;pluginData&#x60; and &#x60;sharedPluginData&#x60; properties. | 
 **branchData** | **bool** | Returns branch metadata for the requested file. If the file is a branch, the main file&#39;s key will be included in the returned response. If the file has branches, their metadata will be included in the returned response. Default: false. | [default to false]

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


## GetFileNodes

> map[string]interface{} GetFileNodes(ctx, fileKey).Ids(ids).Version(version).Depth(depth).Geometry(geometry).PluginData(pluginData).Execute()

Get file JSON for specific nodes



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
	fileKey := "fileKey_example" // string | File to export JSON from. This can be a file key or branch key. Use `GET /v1/files/:key` with the `branch_data` query param to get the branch key.
	ids := "ids_example" // string | A comma separated list of node IDs to retrieve and convert.
	version := "version_example" // string | A specific version ID to get. Omitting this will get the current version of the file. (optional)
	depth := float32(8.14) // float32 | Positive integer representing how deep into the node tree to traverse. For example, setting this to 1 will return only the children directly underneath the desired nodes. Not setting this parameter returns all nodes.  Note: this parameter behaves differently from the same parameter in the `GET /v1/files/:key` endpoint. In this endpoint, the depth will be counted starting from the desired node rather than the document root node. (optional)
	geometry := "geometry_example" // string | Set to \"paths\" to export vector data. (optional)
	pluginData := "pluginData_example" // string | A comma separated list of plugin IDs and/or the string \"shared\". Any data present in the document written by those plugins will be included in the result in the `pluginData` and `sharedPluginData` properties. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFileNodes(context.Background(), fileKey).Ids(ids).Version(version).Depth(depth).Geometry(geometry).PluginData(pluginData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileNodes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File to export JSON from. This can be a file key or branch key. Use &#x60;GET /v1/files/:key&#x60; with the &#x60;branch_data&#x60; query param to get the branch key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **string** | A comma separated list of node IDs to retrieve and convert. | 
 **version** | **string** | A specific version ID to get. Omitting this will get the current version of the file. | 
 **depth** | **float32** | Positive integer representing how deep into the node tree to traverse. For example, setting this to 1 will return only the children directly underneath the desired nodes. Not setting this parameter returns all nodes.  Note: this parameter behaves differently from the same parameter in the &#x60;GET /v1/files/:key&#x60; endpoint. In this endpoint, the depth will be counted starting from the desired node rather than the document root node. | 
 **geometry** | **string** | Set to \&quot;paths\&quot; to export vector data. | 
 **pluginData** | **string** | A comma separated list of plugin IDs and/or the string \&quot;shared\&quot;. Any data present in the document written by those plugins will be included in the result in the &#x60;pluginData&#x60; and &#x60;sharedPluginData&#x60; properties. | 

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


## GetFileVersions

> map[string]interface{} GetFileVersions(ctx, fileKey).PageSize(pageSize).Before(before).After(after).Execute()

Get versions of a file



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
	fileKey := "fileKey_example" // string | File to get version history from. This can be a file key or branch key. Use `GET /v1/files/:key` with the `branch_data` query param to get the branch key.
	pageSize := float32(8.14) // float32 | The number of items returned in a page of the response. If not included, `page_size` is `30`. (optional)
	before := float32(8.14) // float32 | A version ID for one of the versions in the history. Gets versions before this ID. Used for paginating. If the response is not paginated, this link returns the same data in the current response. (optional)
	after := float32(8.14) // float32 | A version ID for one of the versions in the history. Gets versions after this ID. Used for paginating. If the response is not paginated, this property is not included. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFileVersions(context.Background(), fileKey).PageSize(pageSize).Before(before).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileVersions`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File to get version history from. This can be a file key or branch key. Use &#x60;GET /v1/files/:key&#x60; with the &#x60;branch_data&#x60; query param to get the branch key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **float32** | The number of items returned in a page of the response. If not included, &#x60;page_size&#x60; is &#x60;30&#x60;. | 
 **before** | **float32** | A version ID for one of the versions in the history. Gets versions before this ID. Used for paginating. If the response is not paginated, this link returns the same data in the current response. | 
 **after** | **float32** | A version ID for one of the versions in the history. Gets versions after this ID. Used for paginating. If the response is not paginated, this property is not included. | 

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


## GetImageFills

> map[string]interface{} GetImageFills(ctx, fileKey).Execute()

Get image fills



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
	fileKey := "fileKey_example" // string | File to get image URLs from. This can be a file key or branch key. Use `GET /v1/files/:key` with the `branch_data` query param to get the branch key.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetImageFills(context.Background(), fileKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetImageFills``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImageFills`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetImageFills`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File to get image URLs from. This can be a file key or branch key. Use &#x60;GET /v1/files/:key&#x60; with the &#x60;branch_data&#x60; query param to get the branch key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageFillsRequest struct via the builder pattern


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


## GetImages

> map[string]interface{} GetImages(ctx, fileKey).Ids(ids).Version(version).Scale(scale).Format(format).SvgOutlineText(svgOutlineText).SvgIncludeId(svgIncludeId).SvgIncludeNodeId(svgIncludeNodeId).SvgSimplifyStroke(svgSimplifyStroke).ContentsOnly(contentsOnly).UseAbsoluteBounds(useAbsoluteBounds).Execute()

Render images of file nodes



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
	fileKey := "fileKey_example" // string | File to export images from. This can be a file key or branch key. Use `GET /v1/files/:key` with the `branch_data` query param to get the branch key.
	ids := "ids_example" // string | A comma separated list of node IDs to render.
	version := "version_example" // string | A specific version ID to get. Omitting this will get the current version of the file. (optional)
	scale := float32(8.14) // float32 | A number between 0.01 and 4, the image scaling factor. (optional)
	format := "format_example" // string | A string enum for the image output format. (optional) (default to "png")
	svgOutlineText := true // bool | Whether text elements are rendered as outlines (vector paths) or as `<text>` elements in SVGs.  Rendering text elements as outlines guarantees that the text looks exactly the same in the SVG as it does in the browser/inside Figma.  Exporting as `<text>` allows text to be selectable inside SVGs and generally makes the SVG easier to read. However, this relies on the browser's rendering engine which can vary between browsers and/or operating systems. As such, visual accuracy is not guaranteed as the result could look different than in Figma. (optional) (default to true)
	svgIncludeId := true // bool | Whether to include id attributes for all SVG elements. Adds the layer name to the `id` attribute of an svg element. (optional) (default to false)
	svgIncludeNodeId := true // bool | Whether to include node id attributes for all SVG elements. Adds the node id to a `data-node-id` attribute of an svg element. (optional) (default to false)
	svgSimplifyStroke := true // bool | Whether to simplify inside/outside strokes and use stroke attribute if possible instead of `<mask>`. (optional) (default to true)
	contentsOnly := true // bool | Whether content that overlaps the node should be excluded from rendering. Passing false (i.e., rendering overlaps) may increase processing time, since more of the document must be included in rendering. (optional) (default to true)
	useAbsoluteBounds := true // bool | Use the full dimensions of the node regardless of whether or not it is cropped or the space around it is empty. Use this to export text nodes without cropping. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetImages(context.Background(), fileKey).Ids(ids).Version(version).Scale(scale).Format(format).SvgOutlineText(svgOutlineText).SvgIncludeId(svgIncludeId).SvgIncludeNodeId(svgIncludeNodeId).SvgSimplifyStroke(svgSimplifyStroke).ContentsOnly(contentsOnly).UseAbsoluteBounds(useAbsoluteBounds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImages`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetImages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileKey** | **string** | File to export images from. This can be a file key or branch key. Use &#x60;GET /v1/files/:key&#x60; with the &#x60;branch_data&#x60; query param to get the branch key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ids** | **string** | A comma separated list of node IDs to render. | 
 **version** | **string** | A specific version ID to get. Omitting this will get the current version of the file. | 
 **scale** | **float32** | A number between 0.01 and 4, the image scaling factor. | 
 **format** | **string** | A string enum for the image output format. | [default to &quot;png&quot;]
 **svgOutlineText** | **bool** | Whether text elements are rendered as outlines (vector paths) or as &#x60;&lt;text&gt;&#x60; elements in SVGs.  Rendering text elements as outlines guarantees that the text looks exactly the same in the SVG as it does in the browser/inside Figma.  Exporting as &#x60;&lt;text&gt;&#x60; allows text to be selectable inside SVGs and generally makes the SVG easier to read. However, this relies on the browser&#39;s rendering engine which can vary between browsers and/or operating systems. As such, visual accuracy is not guaranteed as the result could look different than in Figma. | [default to true]
 **svgIncludeId** | **bool** | Whether to include id attributes for all SVG elements. Adds the layer name to the &#x60;id&#x60; attribute of an svg element. | [default to false]
 **svgIncludeNodeId** | **bool** | Whether to include node id attributes for all SVG elements. Adds the node id to a &#x60;data-node-id&#x60; attribute of an svg element. | [default to false]
 **svgSimplifyStroke** | **bool** | Whether to simplify inside/outside strokes and use stroke attribute if possible instead of &#x60;&lt;mask&gt;&#x60;. | [default to true]
 **contentsOnly** | **bool** | Whether content that overlaps the node should be excluded from rendering. Passing false (i.e., rendering overlaps) may increase processing time, since more of the document must be included in rendering. | [default to true]
 **useAbsoluteBounds** | **bool** | Use the full dimensions of the node regardless of whether or not it is cropped or the space around it is empty. Use this to export text nodes without cropping. | [default to false]

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

