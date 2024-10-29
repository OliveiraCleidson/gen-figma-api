# SimpleTransition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Duration** | **float32** | The duration of the transition in milliseconds. | 
**Easing** | [**Easing**](Easing.md) | The easing curve of the transition. | 

## Methods

### NewSimpleTransition

`func NewSimpleTransition(type_ string, duration float32, easing Easing, ) *SimpleTransition`

NewSimpleTransition instantiates a new SimpleTransition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleTransitionWithDefaults

`func NewSimpleTransitionWithDefaults() *SimpleTransition`

NewSimpleTransitionWithDefaults instantiates a new SimpleTransition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SimpleTransition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SimpleTransition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SimpleTransition) SetType(v string)`

SetType sets Type field to given value.


### GetDuration

`func (o *SimpleTransition) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SimpleTransition) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SimpleTransition) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetEasing

`func (o *SimpleTransition) GetEasing() Easing`

GetEasing returns the Easing field if non-nil, zero value otherwise.

### GetEasingOk

`func (o *SimpleTransition) GetEasingOk() (*Easing, bool)`

GetEasingOk returns a tuple with the Easing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasing

`func (o *SimpleTransition) SetEasing(v Easing)`

SetEasing sets Easing field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


