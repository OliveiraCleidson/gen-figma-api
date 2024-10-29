# LibraryAnalyticsUsagesByComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentKey** | **string** | Unique, stable id of the component. | 
**ComponentName** | **string** | Name of the component. | 
**NumInstances** | **float32** | The number of instances of the component within the organization. | 
**NumTeamsUsing** | **float32** | The number of teams using the component within the organization. | 
**NumFilesUsing** | **float32** | The number of files using the component within the organization. | 

## Methods

### NewLibraryAnalyticsUsagesByComponent

`func NewLibraryAnalyticsUsagesByComponent(componentKey string, componentName string, numInstances float32, numTeamsUsing float32, numFilesUsing float32, ) *LibraryAnalyticsUsagesByComponent`

NewLibraryAnalyticsUsagesByComponent instantiates a new LibraryAnalyticsUsagesByComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLibraryAnalyticsUsagesByComponentWithDefaults

`func NewLibraryAnalyticsUsagesByComponentWithDefaults() *LibraryAnalyticsUsagesByComponent`

NewLibraryAnalyticsUsagesByComponentWithDefaults instantiates a new LibraryAnalyticsUsagesByComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentKey

`func (o *LibraryAnalyticsUsagesByComponent) GetComponentKey() string`

GetComponentKey returns the ComponentKey field if non-nil, zero value otherwise.

### GetComponentKeyOk

`func (o *LibraryAnalyticsUsagesByComponent) GetComponentKeyOk() (*string, bool)`

GetComponentKeyOk returns a tuple with the ComponentKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentKey

`func (o *LibraryAnalyticsUsagesByComponent) SetComponentKey(v string)`

SetComponentKey sets ComponentKey field to given value.


### GetComponentName

`func (o *LibraryAnalyticsUsagesByComponent) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *LibraryAnalyticsUsagesByComponent) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *LibraryAnalyticsUsagesByComponent) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.


### GetNumInstances

`func (o *LibraryAnalyticsUsagesByComponent) GetNumInstances() float32`

GetNumInstances returns the NumInstances field if non-nil, zero value otherwise.

### GetNumInstancesOk

`func (o *LibraryAnalyticsUsagesByComponent) GetNumInstancesOk() (*float32, bool)`

GetNumInstancesOk returns a tuple with the NumInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumInstances

`func (o *LibraryAnalyticsUsagesByComponent) SetNumInstances(v float32)`

SetNumInstances sets NumInstances field to given value.


### GetNumTeamsUsing

`func (o *LibraryAnalyticsUsagesByComponent) GetNumTeamsUsing() float32`

GetNumTeamsUsing returns the NumTeamsUsing field if non-nil, zero value otherwise.

### GetNumTeamsUsingOk

`func (o *LibraryAnalyticsUsagesByComponent) GetNumTeamsUsingOk() (*float32, bool)`

GetNumTeamsUsingOk returns a tuple with the NumTeamsUsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumTeamsUsing

`func (o *LibraryAnalyticsUsagesByComponent) SetNumTeamsUsing(v float32)`

SetNumTeamsUsing sets NumTeamsUsing field to given value.


### GetNumFilesUsing

`func (o *LibraryAnalyticsUsagesByComponent) GetNumFilesUsing() float32`

GetNumFilesUsing returns the NumFilesUsing field if non-nil, zero value otherwise.

### GetNumFilesUsingOk

`func (o *LibraryAnalyticsUsagesByComponent) GetNumFilesUsingOk() (*float32, bool)`

GetNumFilesUsingOk returns a tuple with the NumFilesUsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumFilesUsing

`func (o *LibraryAnalyticsUsagesByComponent) SetNumFilesUsing(v float32)`

SetNumFilesUsing sets NumFilesUsing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


