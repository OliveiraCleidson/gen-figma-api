# PublishedStyle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The unique identifier for the style | 
**FileKey** | **string** | The unique identifier of the Figma file that contains the style. | 
**NodeId** | **string** | ID of the style node within the figma file | 
**StyleType** | [**StyleType**](StyleType.md) |  | 
**ThumbnailUrl** | Pointer to **string** | A URL to a thumbnail image of the style. | [optional] 
**Name** | **string** | The name of the style. | 
**Description** | **string** | The description of the style as entered by the publisher. | 
**CreatedAt** | **time.Time** | The UTC ISO 8601 time when the style was created. | 
**UpdatedAt** | **time.Time** | The UTC ISO 8601 time when the style was last updated. | 
**User** | [**User**](User.md) | The user who last updated the style. | 
**SortPosition** | **string** | A user specified order number by which the style can be sorted. | 

## Methods

### NewPublishedStyle

`func NewPublishedStyle(key string, fileKey string, nodeId string, styleType StyleType, name string, description string, createdAt time.Time, updatedAt time.Time, user User, sortPosition string, ) *PublishedStyle`

NewPublishedStyle instantiates a new PublishedStyle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedStyleWithDefaults

`func NewPublishedStyleWithDefaults() *PublishedStyle`

NewPublishedStyleWithDefaults instantiates a new PublishedStyle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PublishedStyle) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PublishedStyle) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PublishedStyle) SetKey(v string)`

SetKey sets Key field to given value.


### GetFileKey

`func (o *PublishedStyle) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *PublishedStyle) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *PublishedStyle) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.


### GetNodeId

`func (o *PublishedStyle) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *PublishedStyle) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *PublishedStyle) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetStyleType

`func (o *PublishedStyle) GetStyleType() StyleType`

GetStyleType returns the StyleType field if non-nil, zero value otherwise.

### GetStyleTypeOk

`func (o *PublishedStyle) GetStyleTypeOk() (*StyleType, bool)`

GetStyleTypeOk returns a tuple with the StyleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleType

`func (o *PublishedStyle) SetStyleType(v StyleType)`

SetStyleType sets StyleType field to given value.


### GetThumbnailUrl

`func (o *PublishedStyle) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *PublishedStyle) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *PublishedStyle) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *PublishedStyle) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetName

`func (o *PublishedStyle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublishedStyle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublishedStyle) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PublishedStyle) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublishedStyle) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublishedStyle) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *PublishedStyle) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PublishedStyle) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PublishedStyle) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PublishedStyle) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublishedStyle) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublishedStyle) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUser

`func (o *PublishedStyle) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PublishedStyle) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PublishedStyle) SetUser(v User)`

SetUser sets User field to given value.


### GetSortPosition

`func (o *PublishedStyle) GetSortPosition() string`

GetSortPosition returns the SortPosition field if non-nil, zero value otherwise.

### GetSortPositionOk

`func (o *PublishedStyle) GetSortPositionOk() (*string, bool)`

GetSortPositionOk returns a tuple with the SortPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPosition

`func (o *PublishedStyle) SetSortPosition(v string)`

SetSortPosition sets SortPosition field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


