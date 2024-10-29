# InnerShadowEffect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | [**RGBA**](RGBA.md) | The color of the shadow | 
**BlendMode** | [**BlendMode**](BlendMode.md) | Blend mode of the shadow | 
**Offset** | [**Vector**](Vector.md) | How far the shadow is projected in the x and y directions | 
**Radius** | **float32** | Radius of the blur effect (applies to shadows as well) | 
**Spread** | Pointer to **float32** | The distance by which to expand (or contract) the shadow.  For drop shadows, a positive &#x60;spread&#x60; value creates a shadow larger than the node, whereas a negative value creates a shadow smaller than the node.  For inner shadows, a positive &#x60;spread&#x60; value contracts the shadow. Spread values are only accepted on rectangles and ellipses, or on frames, components, and instances with visible fill paints and &#x60;clipsContent&#x60; enabled. When left unspecified, the default value is 0. | [optional] [default to 0]
**Visible** | **bool** | Whether this shadow is visible. | 
**BoundVariables** | Pointer to [**BaseShadowEffectBoundVariables**](BaseShadowEffectBoundVariables.md) |  | [optional] 
**Type** | Pointer to **string** | A string literal representing the effect&#39;s type. Always check the type before reading other properties. | [optional] 

## Methods

### NewInnerShadowEffect

`func NewInnerShadowEffect(color RGBA, blendMode BlendMode, offset Vector, radius float32, visible bool, ) *InnerShadowEffect`

NewInnerShadowEffect instantiates a new InnerShadowEffect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInnerShadowEffectWithDefaults

`func NewInnerShadowEffectWithDefaults() *InnerShadowEffect`

NewInnerShadowEffectWithDefaults instantiates a new InnerShadowEffect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *InnerShadowEffect) GetColor() RGBA`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *InnerShadowEffect) GetColorOk() (*RGBA, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *InnerShadowEffect) SetColor(v RGBA)`

SetColor sets Color field to given value.


### GetBlendMode

`func (o *InnerShadowEffect) GetBlendMode() BlendMode`

GetBlendMode returns the BlendMode field if non-nil, zero value otherwise.

### GetBlendModeOk

`func (o *InnerShadowEffect) GetBlendModeOk() (*BlendMode, bool)`

GetBlendModeOk returns a tuple with the BlendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlendMode

`func (o *InnerShadowEffect) SetBlendMode(v BlendMode)`

SetBlendMode sets BlendMode field to given value.


### GetOffset

`func (o *InnerShadowEffect) GetOffset() Vector`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *InnerShadowEffect) GetOffsetOk() (*Vector, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *InnerShadowEffect) SetOffset(v Vector)`

SetOffset sets Offset field to given value.


### GetRadius

`func (o *InnerShadowEffect) GetRadius() float32`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *InnerShadowEffect) GetRadiusOk() (*float32, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *InnerShadowEffect) SetRadius(v float32)`

SetRadius sets Radius field to given value.


### GetSpread

`func (o *InnerShadowEffect) GetSpread() float32`

GetSpread returns the Spread field if non-nil, zero value otherwise.

### GetSpreadOk

`func (o *InnerShadowEffect) GetSpreadOk() (*float32, bool)`

GetSpreadOk returns a tuple with the Spread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpread

`func (o *InnerShadowEffect) SetSpread(v float32)`

SetSpread sets Spread field to given value.

### HasSpread

`func (o *InnerShadowEffect) HasSpread() bool`

HasSpread returns a boolean if a field has been set.

### GetVisible

`func (o *InnerShadowEffect) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *InnerShadowEffect) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *InnerShadowEffect) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetBoundVariables

`func (o *InnerShadowEffect) GetBoundVariables() BaseShadowEffectBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *InnerShadowEffect) GetBoundVariablesOk() (*BaseShadowEffectBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *InnerShadowEffect) SetBoundVariables(v BaseShadowEffectBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *InnerShadowEffect) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.

### GetType

`func (o *InnerShadowEffect) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InnerShadowEffect) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InnerShadowEffect) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InnerShadowEffect) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


