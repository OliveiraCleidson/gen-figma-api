# LibraryAnalyticsUsagesByFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** | The name of the file using the library. | 
**TeamName** | **string** | The name of the team the file belongs to. | 
**WorkspaceName** | Pointer to **string** | The name of the workspace that the file belongs to. | [optional] 
**NumInstances** | **float32** | The number of component instances from the library used within the file. | 

## Methods

### NewLibraryAnalyticsUsagesByFile

`func NewLibraryAnalyticsUsagesByFile(fileName string, teamName string, numInstances float32, ) *LibraryAnalyticsUsagesByFile`

NewLibraryAnalyticsUsagesByFile instantiates a new LibraryAnalyticsUsagesByFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryAnalyticsUsagesByFileWithDefaults

`func NewLibraryAnalyticsUsagesByFileWithDefaults() *LibraryAnalyticsUsagesByFile`

NewLibraryAnalyticsUsagesByFileWithDefaults instantiates a new LibraryAnalyticsUsagesByFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *LibraryAnalyticsUsagesByFile) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *LibraryAnalyticsUsagesByFile) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *LibraryAnalyticsUsagesByFile) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetTeamName

`func (o *LibraryAnalyticsUsagesByFile) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *LibraryAnalyticsUsagesByFile) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *LibraryAnalyticsUsagesByFile) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.


### GetWorkspaceName

`func (o *LibraryAnalyticsUsagesByFile) GetWorkspaceName() string`

GetWorkspaceName returns the WorkspaceName field if non-nil, zero value otherwise.

### GetWorkspaceNameOk

`func (o *LibraryAnalyticsUsagesByFile) GetWorkspaceNameOk() (*string, bool)`

GetWorkspaceNameOk returns a tuple with the WorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceName

`func (o *LibraryAnalyticsUsagesByFile) SetWorkspaceName(v string)`

SetWorkspaceName sets WorkspaceName field to given value.

### HasWorkspaceName

`func (o *LibraryAnalyticsUsagesByFile) HasWorkspaceName() bool`

HasWorkspaceName returns a boolean if a field has been set.

### GetNumInstances

`func (o *LibraryAnalyticsUsagesByFile) GetNumInstances() float32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *LibraryAnalyticsUsagesByFile) GetNumInstancesOk() (*float32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *LibraryAnalyticsUsagesByFile) SetNumInstances(v float32)`

SetNumInstances sets NumInstances field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


