# ActivityLogFileEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of entity. | 
**Key** | **string** | Unique identifier of the file. | 
**Name** | **string** | Name of the file. | 
**EditorType** | **string** | Indicates if the object is a file on Figma Design or FigJam. | 
**LinkAccess** | **string** | Access policy for users who have the link to the file. | 
**ProtoLinkAccess** | **string** | Access policy for users who have the link to the file&#39;s prototype. | 

## Methods

### NewActivityLogFileEntity

`func NewActivityLogFileEntity(type_ string, key string, name string, editorType string, linkAccess string, protoLinkAccess string, ) *ActivityLogFileEntity`

NewActivityLogFileEntity instantiates a new ActivityLogFileEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogFileEntityWithDefaults

`func NewActivityLogFileEntityWithDefaults() *ActivityLogFileEntity`

NewActivityLogFileEntityWithDefaults instantiates a new ActivityLogFileEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActivityLogFileEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityLogFileEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityLogFileEntity) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *ActivityLogFileEntity) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ActivityLogFileEntity) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ActivityLogFileEntity) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ActivityLogFileEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivityLogFileEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivityLogFileEntity) SetName(v string)`

SetName sets Name field to given value.


### GetEditorType

`func (o *ActivityLogFileEntity) GetEditorType() string`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *ActivityLogFileEntity) GetEditorTypeOk() (*string, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *ActivityLogFileEntity) SetEditorType(v string)`

SetEditorType sets EditorType field to given value.


### GetLinkAccess

`func (o *ActivityLogFileEntity) GetLinkAccess() string`

GetLinkAccess returns the LinkAccess field if non-nil, zero value otherwise.

### GetLinkAccessOk

`func (o *ActivityLogFileEntity) GetLinkAccessOk() (*string, bool)`

GetLinkAccessOk returns a tuple with the LinkAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAccess

`func (o *ActivityLogFileEntity) SetLinkAccess(v string)`

SetLinkAccess sets LinkAccess field to given value.


### GetProtoLinkAccess

`func (o *ActivityLogFileEntity) GetProtoLinkAccess() string`

GetProtoLinkAccess returns the ProtoLinkAccess field if non-nil, zero value otherwise.

### GetProtoLinkAccessOk

`func (o *ActivityLogFileEntity) GetProtoLinkAccessOk() (*string, bool)`

GetProtoLinkAccessOk returns a tuple with the ProtoLinkAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoLinkAccess

`func (o *ActivityLogFileEntity) SetProtoLinkAccess(v string)`

SetProtoLinkAccess sets ProtoLinkAccess field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


