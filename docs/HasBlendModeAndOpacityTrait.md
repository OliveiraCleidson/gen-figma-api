# HasBlendModeAndOpacityTrait

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlendMode** | [**BlendMode**](BlendMode.md) | How this node blends with nodes behind it in the scene (see blend mode section for more details) | 
**Opacity** | Pointer to **float32** | Opacity of the node | [optional] [default to 1]

## Methods

### NewHasBlendModeAndOpacityTrait

`func NewHasBlendModeAndOpacityTrait(blendMode BlendMode, ) *HasBlendModeAndOpacityTrait`

NewHasBlendModeAndOpacityTrait instantiates a new HasBlendModeAndOpacityTrait object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHasBlendModeAndOpacityTraitWithDefaults

`func NewHasBlendModeAndOpacityTraitWithDefaults() *HasBlendModeAndOpacityTrait`

NewHasBlendModeAndOpacityTraitWithDefaults instantiates a new HasBlendModeAndOpacityTrait object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlendMode

`func (o *HasBlendModeAndOpacityTrait) GetBlendMode() BlendMode`

GetBlendMode returns the BlendMode field if non-nil, zero value otherwise.

### GetBlendModeOk

`func (o *HasBlendModeAndOpacityTrait) GetBlendModeOk() (*BlendMode, bool)`

GetBlendModeOk returns a tuple with the BlendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlendMode

`func (o *HasBlendModeAndOpacityTrait) SetBlendMode(v BlendMode)`

SetBlendMode sets BlendMode field to given value.


### GetOpacity

`func (o *HasBlendModeAndOpacityTrait) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *HasBlendModeAndOpacityTrait) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *HasBlendModeAndOpacityTrait) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *HasBlendModeAndOpacityTrait) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


