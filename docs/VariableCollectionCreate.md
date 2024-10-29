# VariableCollectionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to perform for the variable collection. | 
**Id** | Pointer to **string** | A temporary id for this variable collection. | [optional] 
**Name** | **string** | The name of this variable collection. | 
**InitialModeId** | Pointer to **string** | The initial mode refers to the mode that is created by default. You can set a temporary id here, in order to reference this mode later in this request. | [optional] 
**HiddenFromPublishing** | Pointer to **bool** | Whether this variable collection is hidden when publishing the current file as a library. | [optional] [default to false]

## Methods

### NewVariableCollectionCreate

`func NewVariableCollectionCreate(action string, name string, ) *VariableCollectionCreate`

NewVariableCollectionCreate instantiates a new VariableCollectionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableCollectionCreateWithDefaults

`func NewVariableCollectionCreateWithDefaults() *VariableCollectionCreate`

NewVariableCollectionCreateWithDefaults instantiates a new VariableCollectionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VariableCollectionCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VariableCollectionCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VariableCollectionCreate) SetAction(v string)`

SetAction sets Action field to given value.


### GetId

`func (o *VariableCollectionCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariableCollectionCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariableCollectionCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VariableCollectionCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VariableCollectionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableCollectionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableCollectionCreate) SetName(v string)`

SetName sets Name field to given value.


### GetInitialModeId

`func (o *VariableCollectionCreate) GetInitialModeId() string`

GetInitialModeId returns the InitialModeId field if non-nil, zero value otherwise.

### GetInitialModeIdOk

`func (o *VariableCollectionCreate) GetInitialModeIdOk() (*string, bool)`

GetInitialModeIdOk returns a tuple with the InitialModeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialModeId

`func (o *VariableCollectionCreate) SetInitialModeId(v string)`

SetInitialModeId sets InitialModeId field to given value.

### HasInitialModeId

`func (o *VariableCollectionCreate) HasInitialModeId() bool`

HasInitialModeId returns a boolean if a field has been set.

### GetHiddenFromPublishing

`func (o *VariableCollectionCreate) GetHiddenFromPublishing() bool`

GetHiddenFromPublishing returns the HiddenFromPublishing field if non-nil, zero value otherwise.

### GetHiddenFromPublishingOk

`func (o *VariableCollectionCreate) GetHiddenFromPublishingOk() (*bool, bool)`

GetHiddenFromPublishingOk returns a tuple with the HiddenFromPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromPublishing

`func (o *VariableCollectionCreate) SetHiddenFromPublishing(v bool)`

SetHiddenFromPublishing sets HiddenFromPublishing field to given value.

### HasHiddenFromPublishing

`func (o *VariableCollectionCreate) HasHiddenFromPublishing() bool`

HasHiddenFromPublishing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


