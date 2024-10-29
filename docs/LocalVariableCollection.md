# LocalVariableCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of this variable collection. | 
**Name** | **string** | The name of this variable collection. | 
**Key** | **string** | The key of this variable collection. | 
**Modes** | [**[]LocalVariableCollectionModesInner**](LocalVariableCollectionModesInner.md) | The modes of this variable collection. | 
**DefaultModeId** | **string** | The id of the default mode. | 
**Remote** | **bool** | Whether this variable collection is remote. | 
**HiddenFromPublishing** | **bool** | Whether this variable collection is hidden when publishing the current file as a library. | [default to false]
**VariableIds** | **[]string** | The ids of the variables in the collection. Note that the order of these variables is roughly the same as what is shown in Figma Design, however it does not account for groups. As a result, the order of these variables may not exactly reflect the exact ordering and grouping shown in the authoring UI. | 

## Methods

### NewLocalVariableCollection

`func NewLocalVariableCollection(id string, name string, key string, modes []LocalVariableCollectionModesInner, defaultModeId string, remote bool, hiddenFromPublishing bool, variableIds []string, ) *LocalVariableCollection`

NewLocalVariableCollection instantiates a new LocalVariableCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalVariableCollectionWithDefaults

`func NewLocalVariableCollectionWithDefaults() *LocalVariableCollection`

NewLocalVariableCollectionWithDefaults instantiates a new LocalVariableCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocalVariableCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalVariableCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalVariableCollection) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LocalVariableCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocalVariableCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocalVariableCollection) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *LocalVariableCollection) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LocalVariableCollection) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LocalVariableCollection) SetKey(v string)`

SetKey sets Key field to given value.


### GetModes

`func (o *LocalVariableCollection) GetModes() []LocalVariableCollectionModesInner`

GetModes returns the Modes field if non-nil, zero value otherwise.

### GetModesOk

`func (o *LocalVariableCollection) GetModesOk() (*[]LocalVariableCollectionModesInner, bool)`

GetModesOk returns a tuple with the Modes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModes

`func (o *LocalVariableCollection) SetModes(v []LocalVariableCollectionModesInner)`

SetModes sets Modes field to given value.


### GetDefaultModeId

`func (o *LocalVariableCollection) GetDefaultModeId() string`

GetDefaultModeId returns the DefaultModeId field if non-nil, zero value otherwise.

### GetDefaultModeIdOk

`func (o *LocalVariableCollection) GetDefaultModeIdOk() (*string, bool)`

GetDefaultModeIdOk returns a tuple with the DefaultModeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultModeId

`func (o *LocalVariableCollection) SetDefaultModeId(v string)`

SetDefaultModeId sets DefaultModeId field to given value.


### GetRemote

`func (o *LocalVariableCollection) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *LocalVariableCollection) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *LocalVariableCollection) SetRemote(v bool)`

SetRemote sets Remote field to given value.


### GetHiddenFromPublishing

`func (o *LocalVariableCollection) GetHiddenFromPublishing() bool`

GetHiddenFromPublishing returns the HiddenFromPublishing field if non-nil, zero value otherwise.

### GetHiddenFromPublishingOk

`func (o *LocalVariableCollection) GetHiddenFromPublishingOk() (*bool, bool)`

GetHiddenFromPublishingOk returns a tuple with the HiddenFromPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromPublishing

`func (o *LocalVariableCollection) SetHiddenFromPublishing(v bool)`

SetHiddenFromPublishing sets HiddenFromPublishing field to given value.


### GetVariableIds

`func (o *LocalVariableCollection) GetVariableIds() []string`

GetVariableIds returns the VariableIds field if non-nil, zero value otherwise.

### GetVariableIdsOk

`func (o *LocalVariableCollection) GetVariableIdsOk() (*[]string, bool)`

GetVariableIdsOk returns a tuple with the VariableIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableIds

`func (o *LocalVariableCollection) SetVariableIds(v []string)`

SetVariableIds sets VariableIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


