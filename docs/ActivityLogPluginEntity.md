# ActivityLogPluginEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of entity. | 
**Id** | **string** | Unique identifier of the plugin. | 
**Name** | **string** | Name of the plugin. | 
**EditorType** | **string** | Indicates if the object is a plugin is available on Figma Design or FigJam. | 

## Methods

### NewActivityLogPluginEntity

`func NewActivityLogPluginEntity(type_ string, id string, name string, editorType string, ) *ActivityLogPluginEntity`

NewActivityLogPluginEntity instantiates a new ActivityLogPluginEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogPluginEntityWithDefaults

`func NewActivityLogPluginEntityWithDefaults() *ActivityLogPluginEntity`

NewActivityLogPluginEntityWithDefaults instantiates a new ActivityLogPluginEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActivityLogPluginEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityLogPluginEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityLogPluginEntity) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ActivityLogPluginEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityLogPluginEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityLogPluginEntity) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ActivityLogPluginEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivityLogPluginEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivityLogPluginEntity) SetName(v string)`

SetName sets Name field to given value.


### GetEditorType

`func (o *ActivityLogPluginEntity) GetEditorType() string`

GetEditorType returns the EditorType field if non-nil, zero value otherwise.

### GetEditorTypeOk

`func (o *ActivityLogPluginEntity) GetEditorTypeOk() (*string, bool)`

GetEditorTypeOk returns a tuple with the EditorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditorType

`func (o *ActivityLogPluginEntity) SetEditorType(v string)`

SetEditorType sets EditorType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


