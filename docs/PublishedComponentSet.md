# PublishedComponentSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The unique identifier for the component set. | 
**FileKey** | **string** | The unique identifier of the Figma file that contains the component set. | 
**NodeId** | **string** | The unique identifier of the component set node within the Figma file. | 
**ThumbnailUrl** | Pointer to **string** | A URL to a thumbnail image of the component set. | [optional] 
**Name** | **string** | The name of the component set. | 
**Description** | **string** | The description of the component set as entered by the publisher. | 
**CreatedAt** | **time.Time** | The UTC ISO 8601 time when the component set was created. | 
**UpdatedAt** | **time.Time** | The UTC ISO 8601 time when the component set was last updated. | 
**User** | [**User**](User.md) | The user who last updated the component set. | 
**ContainingFrame** | Pointer to [**FrameInfo**](FrameInfo.md) | The containing frame of the component set. | [optional] 

## Methods

### NewPublishedComponentSet

`func NewPublishedComponentSet(key string, fileKey string, nodeId string, name string, description string, createdAt time.Time, updatedAt time.Time, user User, ) *PublishedComponentSet`

NewPublishedComponentSet instantiates a new PublishedComponentSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedComponentSetWithDefaults

`func NewPublishedComponentSetWithDefaults() *PublishedComponentSet`

NewPublishedComponentSetWithDefaults instantiates a new PublishedComponentSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PublishedComponentSet) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PublishedComponentSet) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PublishedComponentSet) SetKey(v string)`

SetKey sets Key field to given value.


### GetFileKey

`func (o *PublishedComponentSet) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *PublishedComponentSet) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *PublishedComponentSet) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.


### GetNodeId

`func (o *PublishedComponentSet) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PublishedComponentSet) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PublishedComponentSet) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetThumbnailUrl

`func (o *PublishedComponentSet) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *PublishedComponentSet) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *PublishedComponentSet) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *PublishedComponentSet) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetName

`func (o *PublishedComponentSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublishedComponentSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublishedComponentSet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PublishedComponentSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublishedComponentSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublishedComponentSet) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *PublishedComponentSet) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublishedComponentSet) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublishedComponentSet) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PublishedComponentSet) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublishedComponentSet) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublishedComponentSet) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUser

`func (o *PublishedComponentSet) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PublishedComponentSet) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PublishedComponentSet) SetUser(v User)`

SetUser sets User field to given value.


### GetContainingFrame

`func (o *PublishedComponentSet) GetContainingFrame() FrameInfo`

GetContainingFrame returns the ContainingFrame field if non-nil, zero value otherwise.

### GetContainingFrameOk

`func (o *PublishedComponentSet) GetContainingFrameOk() (*FrameInfo, bool)`

GetContainingFrameOk returns a tuple with the ContainingFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainingFrame

`func (o *PublishedComponentSet) SetContainingFrame(v FrameInfo)`

SetContainingFrame sets ContainingFrame field to given value.

### HasContainingFrame

`func (o *PublishedComponentSet) HasContainingFrame() bool`

HasContainingFrame returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


