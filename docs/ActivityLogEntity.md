# ActivityLogEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of entity. | 
**Id** | **string** | Unique identifier of the widget. | 
**Name** | **string** | Name of the widget. | 
**Email** | **string** | Email associated with the user&#39;s account. | 
**Key** | **string** | Unique identifier of the file. | 
**EditorType** | **string** | Indicates if the object is a widget available on Figma Design or FigJam. | 
**LinkAccess** | **string** | Access policy for users who have the link to the file. | 
**ProtoLinkAccess** | **string** | Access policy for users who have the link to the file&#39;s prototype. | 
**MainFileKey** | **string** | Key of the main file. | 

## Methods

### NewActivityLogEntity

`func NewActivityLogEntity(type_ string, id string, name string, email string, key string, editorType string, linkAccess string, protoLinkAccess string, mainFileKey string, ) *ActivityLogEntity`

NewActivityLogEntity instantiates a new ActivityLogEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogEntityWithDefaults

`func NewActivityLogEntityWithDefaults() *ActivityLogEntity`

NewActivityLogEntityWithDefaults instantiates a new ActivityLogEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActivityLogEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityLogEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityLogEntity) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ActivityLogEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityLogEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityLogEntity) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ActivityLogEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivityLogEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivityLogEntity) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *ActivityLogEntity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ActivityLogEntity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ActivityLogEntity) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetKey

`func (o *ActivityLogEntity) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ActivityLogEntity) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ActivityLogEntity) SetKey(v string)`

SetKey sets Key field to given value.


### GetEditorType

`func (o *ActivityLogEntity) GetEditorType() string`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *ActivityLogEntity) GetEditorTypeOk() (*string, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *ActivityLogEntity) SetEditorType(v string)`

SetEditorType sets EditorType field to given value.


### GetLinkAccess

`func (o *ActivityLogEntity) GetLinkAccess() string`

GetLinkAccess returns the LinkAccess field if non-nil, zero value otherwise.

### GetLinkAccessOk

`func (o *ActivityLogEntity) GetLinkAccessOk() (*string, bool)`

GetLinkAccessOk returns a tuple with the LinkAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAccess

`func (o *ActivityLogEntity) SetLinkAccess(v string)`

SetLinkAccess sets LinkAccess field to given value.


### GetProtoLinkAccess

`func (o *ActivityLogEntity) GetProtoLinkAccess() string`

GetProtoLinkAccess returns the ProtoLinkAccess field if non-nil, zero value otherwise.

### GetProtoLinkAccessOk

`func (o *ActivityLogEntity) GetProtoLinkAccessOk() (*string, bool)`

GetProtoLinkAccessOk returns a tuple with the ProtoLinkAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoLinkAccess

`func (o *ActivityLogEntity) SetProtoLinkAccess(v string)`

SetProtoLinkAccess sets ProtoLinkAccess field to given value.


### GetMainFileKey

`func (o *ActivityLogEntity) GetMainFileKey() string`

GetMainFileKey returns the MainFileKey field if non-nil, zero value otherwise.

### GetMainFileKeyOk

`func (o *ActivityLogEntity) GetMainFileKeyOk() (*string, bool)`

GetMainFileKeyOk returns a tuple with the MainFileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainFileKey

`func (o *ActivityLogEntity) SetMainFileKey(v string)`

SetMainFileKey sets MainFileKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


