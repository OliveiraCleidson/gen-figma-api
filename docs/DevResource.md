# DevResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the dev resource | 
**Name** | **string** | The name of the dev resource. | 
**Url** | **string** | The URL of the dev resource. | 
**FileKey** | **string** | The file key where the dev resource belongs. | 
**NodeId** | **string** | The target node to attach the dev resource to. | 

## Methods

### NewDevResource

`func NewDevResource(id string, name string, url string, fileKey string, nodeId string, ) *DevResource`

NewDevResource instantiates a new DevResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevResourceWithDefaults

`func NewDevResourceWithDefaults() *DevResource`

NewDevResourceWithDefaults instantiates a new DevResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DevResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DevResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DevResource) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DevResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevResource) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *DevResource) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DevResource) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DevResource) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFileKey

`func (o *DevResource) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *DevResource) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *DevResource) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.


### GetNodeId

`func (o *DevResource) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *DevResource) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *DevResource) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


