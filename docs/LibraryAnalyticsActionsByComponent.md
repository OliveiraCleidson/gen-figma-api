# LibraryAnalyticsActionsByComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Week** | **string** | The date in ISO 8601 format. e.g. 2023-12-13 | 
**ComponentKey** | **string** | Unique, stable id of the component. | 
**ComponentName** | **string** | Name of the component. | 
**Detachments** | **float32** | The number of detach events for this period. | 
**Insertions** | **float32** | The number of insertion events for this period. | 

## Methods

### NewLibraryAnalyticsActionsByComponent

`func NewLibraryAnalyticsActionsByComponent(week string, componentKey string, componentName string, detachments float32, insertions float32, ) *LibraryAnalyticsActionsByComponent`

NewLibraryAnalyticsActionsByComponent instantiates a new LibraryAnalyticsActionsByComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryAnalyticsActionsByComponentWithDefaults

`func NewLibraryAnalyticsActionsByComponentWithDefaults() *LibraryAnalyticsActionsByComponent`

NewLibraryAnalyticsActionsByComponentWithDefaults instantiates a new LibraryAnalyticsActionsByComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeek

`func (o *LibraryAnalyticsActionsByComponent) GetWeek() string`

GetWeek returns the Week field if non-nil, zero value otherwise.

### GetWeekOk

`func (o *LibraryAnalyticsActionsByComponent) GetWeekOk() (*string, bool)`

GetWeekOk returns a tuple with the Week field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeek

`func (o *LibraryAnalyticsActionsByComponent) SetWeek(v string)`

SetWeek sets Week field to given value.


### GetComponentKey

`func (o *LibraryAnalyticsActionsByComponent) GetComponentKey() string`

GetComponentKey returns the ComponentKey field if non-nil, zero value otherwise.

### GetComponentKeyOk

`func (o *LibraryAnalyticsActionsByComponent) GetComponentKeyOk() (*string, bool)`

GetComponentKeyOk returns a tuple with the ComponentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentKey

`func (o *LibraryAnalyticsActionsByComponent) SetComponentKey(v string)`

SetComponentKey sets ComponentKey field to given value.


### GetComponentName

`func (o *LibraryAnalyticsActionsByComponent) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *LibraryAnalyticsActionsByComponent) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *LibraryAnalyticsActionsByComponent) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.


### GetDetachments

`func (o *LibraryAnalyticsActionsByComponent) GetDetachments() float32`

GetDetachments returns the Detachments field if non-nil, zero value otherwise.

### GetDetachmentsOk

`func (o *LibraryAnalyticsActionsByComponent) GetDetachmentsOk() (*float32, bool)`

GetDetachmentsOk returns a tuple with the Detachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetachments

`func (o *LibraryAnalyticsActionsByComponent) SetDetachments(v float32)`

SetDetachments sets Detachments field to given value.


### GetInsertions

`func (o *LibraryAnalyticsActionsByComponent) GetInsertions() float32`

GetInsertions returns the Insertions field if non-nil, zero value otherwise.

### GetInsertionsOk

`func (o *LibraryAnalyticsActionsByComponent) GetInsertionsOk() (*float32, bool)`

GetInsertionsOk returns a tuple with the Insertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertions

`func (o *LibraryAnalyticsActionsByComponent) SetInsertions(v float32)`

SetInsertions sets Insertions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


