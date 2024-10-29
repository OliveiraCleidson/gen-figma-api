# \ProjectsAPI

All URIs are relative to *https://api.figma.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectFiles**](ProjectsAPI.md#GetProjectFiles) | **Get** /v1/projects/{project_id}/files | Get files in a project
[**GetTeamProjects**](ProjectsAPI.md#GetTeamProjects) | **Get** /v1/teams/{team_id}/projects | Get projects in a team



## GetProjectFiles

> map[string]interface{} GetProjectFiles(ctx, projectId).BranchData(branchData).Execute()

Get files in a project



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
	projectId := "projectId_example" // string | ID of the project to list files from
	branchData := true // bool | Returns branch metadata in the response for each main file with a branch inside the project. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetProjectFiles(context.Background(), projectId).BranchData(branchData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetProjectFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectFiles`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetProjectFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | ID of the project to list files from | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **branchData** | **bool** | Returns branch metadata in the response for each main file with a branch inside the project. | [default to false]

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


## GetTeamProjects

> map[string]interface{} GetTeamProjects(ctx, teamId).Execute()

Get projects in a team



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
	teamId := "teamId_example" // string | ID of the team to list projects from

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.GetTeamProjects(context.Background(), teamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.GetTeamProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeamProjects`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.GetTeamProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**teamId** | **string** | ID of the team to list projects from | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamProjectsRequest struct via the builder pattern


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

