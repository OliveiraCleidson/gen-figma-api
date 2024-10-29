# Transition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Duration** | **float32** | The duration of the transition in milliseconds. | 
**Easing** | [**Easing**](Easing.md) | The easing curve of the transition. | 
**Direction** | **string** |  | 
**MatchLayers** | Pointer to **bool** | When the transition &#x60;type&#x60; is &#x60;\&quot;SMART_ANIMATE\&quot;&#x60; or when &#x60;matchLayers&#x60; is &#x60;true&#x60;, then the transition will be performed using smart animate, which attempts to match corresponding layers an interpolate other properties during the animation. | [optional] 

## Methods

### NewTransition

`func NewTransition(type_ string, duration float32, easing Easing, direction string, ) *Transition`

NewTransition instantiates a new Transition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitionWithDefaults

`func NewTransitionWithDefaults() *Transition`

NewTransitionWithDefaults instantiates a new Transition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Transition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transition) SetType(v string)`

SetType sets Type field to given value.


### GetDuration

`func (o *Transition) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Transition) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Transition) SetDuration(v float32)`

SetDuration sets Duration field to given value.


### GetEasing

`func (o *Transition) GetEasing() Easing`

GetEasing returns the Easing field if non-nil, zero value otherwise.

### GetEasingOk

`func (o *Transition) GetEasingOk() (*Easing, bool)`

GetEasingOk returns a tuple with the Easing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasing

`func (o *Transition) SetEasing(v Easing)`

SetEasing sets Easing field to given value.


### GetDirection

`func (o *Transition) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Transition) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Transition) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetMatchLayers

`func (o *Transition) GetMatchLayers() bool`

GetMatchLayers returns the MatchLayers field if non-nil, zero value otherwise.

### GetMatchLayersOk

`func (o *Transition) GetMatchLayersOk() (*bool, bool)`

GetMatchLayersOk returns a tuple with the MatchLayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchLayers

`func (o *Transition) SetMatchLayers(v bool)`

SetMatchLayers sets MatchLayers field to given value.

### HasMatchLayers

`func (o *Transition) HasMatchLayers() bool`

HasMatchLayers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


