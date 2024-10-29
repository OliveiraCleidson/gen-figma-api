# VariableModeUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to perform for the variable mode. | 
**Id** | **string** | The id of the variable mode to update. | 
**Name** | Pointer to **string** | The name of this variable mode. | [optional] 
**VariableCollectionId** | **string** | The variable collection that contains the mode. | 

## Methods

### NewVariableModeUpdate

`func NewVariableModeUpdate(action string, id string, variableCollectionId string, ) *VariableModeUpdate`

NewVariableModeUpdate instantiates a new VariableModeUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableModeUpdateWithDefaults

`func NewVariableModeUpdateWithDefaults() *VariableModeUpdate`

NewVariableModeUpdateWithDefaults instantiates a new VariableModeUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VariableModeUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VariableModeUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VariableModeUpdate) SetAction(v string)`

SetAction sets Action field to given value.


### GetId

`func (o *VariableModeUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariableModeUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariableModeUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VariableModeUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableModeUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableModeUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VariableModeUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVariableCollectionId

`func (o *VariableModeUpdate) GetVariableCollectionId() string`

GetVariableCollectionId returns the VariableCollectionId field if non-nil, zero value otherwise.

### GetVariableCollectionIdOk

`func (o *VariableModeUpdate) GetVariableCollectionIdOk() (*string, bool)`

GetVariableCollectionIdOk returns a tuple with the VariableCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableCollectionId

`func (o *VariableModeUpdate) SetVariableCollectionId(v string)`

SetVariableCollectionId sets VariableCollectionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


