# ConditionalBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**VariableData**](VariableData.md) |  | [optional] 
**Actions** | [**[]Action**](Action.md) |  | 

## Methods

### NewConditionalBlock

`func NewConditionalBlock(actions []Action, ) *ConditionalBlock`

NewConditionalBlock instantiates a new ConditionalBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionalBlockWithDefaults

`func NewConditionalBlockWithDefaults() *ConditionalBlock`

NewConditionalBlockWithDefaults instantiates a new ConditionalBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *ConditionalBlock) GetCondition() VariableData`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ConditionalBlock) GetConditionOk() (*VariableData, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ConditionalBlock) SetCondition(v VariableData)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ConditionalBlock) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetActions

`func (o *ConditionalBlock) GetActions() []Action`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ConditionalBlock) GetActionsOk() (*[]Action, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ConditionalBlock) SetActions(v []Action)`

SetActions sets Actions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


