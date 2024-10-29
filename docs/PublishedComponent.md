# PublishedComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The unique identifier for the component. | 
**FileKey** | **string** | The unique identifier of the Figma file that contains the component. | 
**NodeId** | **string** | The unique identifier of the component node within the Figma file. | 
**ThumbnailUrl** | Pointer to **string** | A URL to a thumbnail image of the component. | [optional] 
**Name** | **string** | The name of the component. | 
**Description** | **string** | The description of the component as entered by the publisher. | 
**CreatedAt** | **time.Time** | The UTC ISO 8601 time when the component was created. | 
**UpdatedAt** | **time.Time** | The UTC ISO 8601 time when the component was last updated. | 
**User** | [**User**](User.md) | The user who last updated the component. | 
**ContainingFrame** | Pointer to [**FrameInfo**](FrameInfo.md) | The containing frame of the component. | [optional] 

## Methods

### NewPublishedComponent

`func NewPublishedComponent(key string, fileKey string, nodeId string, name string, description string, createdAt time.Time, updatedAt time.Time, user User, ) *PublishedComponent`

NewPublishedComponent instantiates a new PublishedComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedComponentWithDefaults

`func NewPublishedComponentWithDefaults() *PublishedComponent`

NewPublishedComponentWithDefaults instantiates a new PublishedComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PublishedComponent) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PublishedComponent) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PublishedComponent) SetKey(v string)`

SetKey sets Key field to given value.


### GetFileKey

`func (o *PublishedComponent) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *PublishedComponent) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *PublishedComponent) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.


### GetNodeId

`func (o *PublishedComponent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PublishedComponent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PublishedComponent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetThumbnailUrl

`func (o *PublishedComponent) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *PublishedComponent) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *PublishedComponent) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *PublishedComponent) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetName

`func (o *PublishedComponent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublishedComponent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublishedComponent) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PublishedComponent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublishedComponent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublishedComponent) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *PublishedComponent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublishedComponent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublishedComponent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PublishedComponent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublishedComponent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublishedComponent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUser

`func (o *PublishedComponent) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PublishedComponent) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PublishedComponent) SetUser(v User)`

SetUser sets User field to given value.


### GetContainingFrame

`func (o *PublishedComponent) GetContainingFrame() FrameInfo`

GetContainingFrame returns the ContainingFrame field if non-nil, zero value otherwise.

### GetContainingFrameOk

`func (o *PublishedComponent) GetContainingFrameOk() (*FrameInfo, bool)`

GetContainingFrameOk returns a tuple with the ContainingFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainingFrame

`func (o *PublishedComponent) SetContainingFrame(v FrameInfo)`

SetContainingFrame sets ContainingFrame field to given value.

### HasContainingFrame

`func (o *PublishedComponent) HasContainingFrame() bool`

HasContainingFrame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


