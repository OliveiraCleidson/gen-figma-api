# Easing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EasingType**](EasingType.md) | The type of easing curve. | 
**EasingFunctionCubicBezier** | Pointer to [**EasingEasingFunctionCubicBezier**](EasingEasingFunctionCubicBezier.md) |  | [optional] 
**EasingFunctionSpring** | Pointer to [**EasingEasingFunctionSpring**](EasingEasingFunctionSpring.md) |  | [optional] 

## Methods

### NewEasing

`func NewEasing(type_ EasingType, ) *Easing`

NewEasing instantiates a new Easing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEasingWithDefaults

`func NewEasingWithDefaults() *Easing`

NewEasingWithDefaults instantiates a new Easing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Easing) GetType() EasingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Easing) GetTypeOk() (*EasingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Easing) SetType(v EasingType)`

SetType sets Type field to given value.


### GetEasingFunctionCubicBezier

`func (o *Easing) GetEasingFunctionCubicBezier() EasingEasingFunctionCubicBezier`

GetEasingFunctionCubicBezier returns the EasingFunctionCubicBezier field if non-nil, zero value otherwise.

### GetEasingFunctionCubicBezierOk

`func (o *Easing) GetEasingFunctionCubicBezierOk() (*EasingEasingFunctionCubicBezier, bool)`

GetEasingFunctionCubicBezierOk returns a tuple with the EasingFunctionCubicBezier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasingFunctionCubicBezier

`func (o *Easing) SetEasingFunctionCubicBezier(v EasingEasingFunctionCubicBezier)`

SetEasingFunctionCubicBezier sets EasingFunctionCubicBezier field to given value.

### HasEasingFunctionCubicBezier

`func (o *Easing) HasEasingFunctionCubicBezier() bool`

HasEasingFunctionCubicBezier returns a boolean if a field has been set.

### GetEasingFunctionSpring

`func (o *Easing) GetEasingFunctionSpring() EasingEasingFunctionSpring`

GetEasingFunctionSpring returns the EasingFunctionSpring field if non-nil, zero value otherwise.

### GetEasingFunctionSpringOk

`func (o *Easing) GetEasingFunctionSpringOk() (*EasingEasingFunctionSpring, bool)`

GetEasingFunctionSpringOk returns a tuple with the EasingFunctionSpring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasingFunctionSpring

`func (o *Easing) SetEasingFunctionSpring(v EasingEasingFunctionSpring)`

SetEasingFunctionSpring sets EasingFunctionSpring field to given value.

### HasEasingFunctionSpring

`func (o *Easing) HasEasingFunctionSpring() bool`

HasEasingFunctionSpring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


