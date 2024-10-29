# LibraryAnalyticsActionsByTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Week** | **string** | The date in ISO 8601 format. e.g. 2023-12-13 | 
**TeamName** | **string** | The name of the team using the library. | 
**WorkspaceName** | Pointer to **string** | The name of the workspace that the team belongs to. | [optional] 
**Detachments** | **float32** | The number of detach events for this period. | 
**Insertions** | **float32** | The number of insertion events for this period. | 

## Methods

### NewLibraryAnalyticsActionsByTeam

`func NewLibraryAnalyticsActionsByTeam(week string, teamName string, detachments float32, insertions float32, ) *LibraryAnalyticsActionsByTeam`

NewLibraryAnalyticsActionsByTeam instantiates a new LibraryAnalyticsActionsByTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryAnalyticsActionsByTeamWithDefaults

`func NewLibraryAnalyticsActionsByTeamWithDefaults() *LibraryAnalyticsActionsByTeam`

NewLibraryAnalyticsActionsByTeamWithDefaults instantiates a new LibraryAnalyticsActionsByTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeek

`func (o *LibraryAnalyticsActionsByTeam) GetWeek() string`

GetWeek returns the Week field if non-nil, zero value otherwise.

### GetWeekOk

`func (o *LibraryAnalyticsActionsByTeam) GetWeekOk() (*string, bool)`

GetWeekOk returns a tuple with the Week field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeek

`func (o *LibraryAnalyticsActionsByTeam) SetWeek(v string)`

SetWeek sets Week field to given value.


### GetTeamName

`func (o *LibraryAnalyticsActionsByTeam) GetTeamName() string`

GetTeamName returns the TeamName field if non-nil, zero value otherwise.

### GetTeamNameOk

`func (o *LibraryAnalyticsActionsByTeam) GetTeamNameOk() (*string, bool)`

GetTeamNameOk returns a tuple with the TeamName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamName

`func (o *LibraryAnalyticsActionsByTeam) SetTeamName(v string)`

SetTeamName sets TeamName field to given value.


### GetWorkspaceName

`func (o *LibraryAnalyticsActionsByTeam) GetWorkspaceName() string`

GetWorkspaceName returns the WorkspaceName field if non-nil, zero value otherwise.

### GetWorkspaceNameOk

`func (o *LibraryAnalyticsActionsByTeam) GetWorkspaceNameOk() (*string, bool)`

GetWorkspaceNameOk returns a tuple with the WorkspaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceName

`func (o *LibraryAnalyticsActionsByTeam) SetWorkspaceName(v string)`

SetWorkspaceName sets WorkspaceName field to given value.

### HasWorkspaceName

`func (o *LibraryAnalyticsActionsByTeam) HasWorkspaceName() bool`

HasWorkspaceName returns a boolean if a field has been set.

### GetDetachments

`func (o *LibraryAnalyticsActionsByTeam) GetDetachments() float32`

GetDetachments returns the Detachments field if non-nil, zero value otherwise.

### GetDetachmentsOk

`func (o *LibraryAnalyticsActionsByTeam) GetDetachmentsOk() (*float32, bool)`

GetDetachmentsOk returns a tuple with the Detachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachments

`func (o *LibraryAnalyticsActionsByTeam) SetDetachments(v float32)`

SetDetachments sets Detachments field to given value.


### GetInsertions

`func (o *LibraryAnalyticsActionsByTeam) GetInsertions() float32`

GetInsertions returns the Insertions field if non-nil, zero value otherwise.

### GetInsertionsOk

`func (o *LibraryAnalyticsActionsByTeam) GetInsertionsOk() (*float32, bool)`

GetInsertionsOk returns a tuple with the Insertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertions

`func (o *LibraryAnalyticsActionsByTeam) SetInsertions(v float32)`

SetInsertions sets Insertions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


