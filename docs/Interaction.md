# Interaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trigger** | [**NullableTrigger**](Trigger.md) |  | 
**Actions** | Pointer to [**[]Action**](Action.md) | The actions that are performed when the trigger is activated. | [optional] 

## Methods

### NewInteraction

`func NewInteraction(trigger NullableTrigger, ) *Interaction`

NewInteraction instantiates a new Interaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInteractionWithDefaults

`func NewInteractionWithDefaults() *Interaction`

NewInteractionWithDefaults instantiates a new Interaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrigger

`func (o *Interaction) GetTrigger() Trigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *Interaction) GetTriggerOk() (*Trigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *Interaction) SetTrigger(v Trigger)`

SetTrigger sets Trigger field to given value.


### SetTriggerNil

`func (o *Interaction) SetTriggerNil(b bool)`

 SetTriggerNil sets the value for Trigger to be an explicit nil

### UnsetTrigger
`func (o *Interaction) UnsetTrigger()`

UnsetTrigger ensures that no value is present for Trigger, not even an explicit nil
### GetActions

`func (o *Interaction) GetActions() []Action`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Interaction) GetActionsOk() (*[]Action, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Interaction) SetActions(v []Action)`

SetActions sets Actions field to given value.

### HasActions

`func (o *Interaction) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


