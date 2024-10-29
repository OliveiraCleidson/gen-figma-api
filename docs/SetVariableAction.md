# SetVariableAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**VariableId** | **NullableString** |  | 
**VariableValue** | Pointer to [**VariableData**](VariableData.md) |  | [optional] 

## Methods

### NewSetVariableAction

`func NewSetVariableAction(type_ string, variableId NullableString, ) *SetVariableAction`

NewSetVariableAction instantiates a new SetVariableAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetVariableActionWithDefaults

`func NewSetVariableActionWithDefaults() *SetVariableAction`

NewSetVariableActionWithDefaults instantiates a new SetVariableAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SetVariableAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SetVariableAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SetVariableAction) SetType(v string)`

SetType sets Type field to given value.


### GetVariableId

`func (o *SetVariableAction) GetVariableId() string`

GetVariableId returns the VariableId field if non-nil, zero value otherwise.

### GetVariableIdOk

`func (o *SetVariableAction) GetVariableIdOk() (*string, bool)`

GetVariableIdOk returns a tuple with the VariableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableId

`func (o *SetVariableAction) SetVariableId(v string)`

SetVariableId sets VariableId field to given value.


### SetVariableIdNil

`func (o *SetVariableAction) SetVariableIdNil(b bool)`

 SetVariableIdNil sets the value for VariableId to be an explicit nil

### UnsetVariableId
`func (o *SetVariableAction) UnsetVariableId()`

UnsetVariableId ensures that no value is present for VariableId, not even an explicit nil
### GetVariableValue

`func (o *SetVariableAction) GetVariableValue() VariableData`

GetVariableValue returns the VariableValue field if non-nil, zero value otherwise.

### GetVariableValueOk

`func (o *SetVariableAction) GetVariableValueOk() (*VariableData, bool)`

GetVariableValueOk returns a tuple with the VariableValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValue

`func (o *SetVariableAction) SetVariableValue(v VariableData)`

SetVariableValue sets VariableValue field to given value.

### HasVariableValue

`func (o *SetVariableAction) HasVariableValue() bool`

HasVariableValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


