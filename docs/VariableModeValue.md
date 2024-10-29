# VariableModeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariableId** | **string** | The target variable. You can use the temporary id of a variable. | 
**ModeId** | **string** | Must correspond to a mode in the variable collection that contains the target variable. | 
**Value** | [**VariableValue**](VariableValue.md) |  | 

## Methods

### NewVariableModeValue

`func NewVariableModeValue(variableId string, modeId string, value VariableValue, ) *VariableModeValue`

NewVariableModeValue instantiates a new VariableModeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableModeValueWithDefaults

`func NewVariableModeValueWithDefaults() *VariableModeValue`

NewVariableModeValueWithDefaults instantiates a new VariableModeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariableId

`func (o *VariableModeValue) GetVariableId() string`

GetVariableId returns the VariableId field if non-nil, zero value otherwise.

### GetVariableIdOk

`func (o *VariableModeValue) GetVariableIdOk() (*string, bool)`

GetVariableIdOk returns a tuple with the VariableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableId

`func (o *VariableModeValue) SetVariableId(v string)`

SetVariableId sets VariableId field to given value.


### GetModeId

`func (o *VariableModeValue) GetModeId() string`

GetModeId returns the ModeId field if non-nil, zero value otherwise.

### GetModeIdOk

`func (o *VariableModeValue) GetModeIdOk() (*string, bool)`

GetModeIdOk returns a tuple with the ModeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeId

`func (o *VariableModeValue) SetModeId(v string)`

SetModeId sets ModeId field to given value.


### GetValue

`func (o *VariableModeValue) GetValue() VariableValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableModeValue) GetValueOk() (*VariableValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableModeValue) SetValue(v VariableValue)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


