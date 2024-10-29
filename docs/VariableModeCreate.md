# VariableModeCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to perform for the variable mode. | 
**Id** | Pointer to **string** | A temporary id for this variable mode. | [optional] 
**Name** | **string** | The name of this variable mode. | 
**VariableCollectionId** | **string** | The variable collection that will contain the mode. You can use the temporary id of a variable collection. | 

## Methods

### NewVariableModeCreate

`func NewVariableModeCreate(action string, name string, variableCollectionId string, ) *VariableModeCreate`

NewVariableModeCreate instantiates a new VariableModeCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableModeCreateWithDefaults

`func NewVariableModeCreateWithDefaults() *VariableModeCreate`

NewVariableModeCreateWithDefaults instantiates a new VariableModeCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VariableModeCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VariableModeCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VariableModeCreate) SetAction(v string)`

SetAction sets Action field to given value.


### GetId

`func (o *VariableModeCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariableModeCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariableModeCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VariableModeCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VariableModeCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableModeCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableModeCreate) SetName(v string)`

SetName sets Name field to given value.


### GetVariableCollectionId

`func (o *VariableModeCreate) GetVariableCollectionId() string`

GetVariableCollectionId returns the VariableCollectionId field if non-nil, zero value otherwise.

### GetVariableCollectionIdOk

`func (o *VariableModeCreate) GetVariableCollectionIdOk() (*string, bool)`

GetVariableCollectionIdOk returns a tuple with the VariableCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableCollectionId

`func (o *VariableModeCreate) SetVariableCollectionId(v string)`

SetVariableCollectionId sets VariableCollectionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


