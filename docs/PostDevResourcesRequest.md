# PostDevResourcesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevResources** | [**[]PostDevResourcesRequestDevResourcesInner**](PostDevResourcesRequestDevResourcesInner.md) | An array of dev resources. | 

## Methods

### NewPostDevResourcesRequest

`func NewPostDevResourcesRequest(devResources []PostDevResourcesRequestDevResourcesInner, ) *PostDevResourcesRequest`

NewPostDevResourcesRequest instantiates a new PostDevResourcesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDevResourcesRequestWithDefaults

`func NewPostDevResourcesRequestWithDefaults() *PostDevResourcesRequest`

NewPostDevResourcesRequestWithDefaults instantiates a new PostDevResourcesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevResources

`func (o *PostDevResourcesRequest) GetDevResources() []PostDevResourcesRequestDevResourcesInner`

GetDevResources returns the DevResources field if non-nil, zero value otherwise.

### GetDevResourcesOk

`func (o *PostDevResourcesRequest) GetDevResourcesOk() (*[]PostDevResourcesRequestDevResourcesInner, bool)`

GetDevResourcesOk returns a tuple with the DevResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevResources

`func (o *PostDevResourcesRequest) SetDevResources(v []PostDevResourcesRequestDevResourcesInner)`

SetDevResources sets DevResources field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


