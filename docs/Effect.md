# Effect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A string literal representing the effect&#39;s type. Always check the type before reading other properties. | 
**ShowShadowBehindNode** | **bool** | Whether to show the shadow behind translucent or transparent pixels | [default to false]
**Color** | [**RGBA**](RGBA.md) | The color of the shadow | 
**BlendMode** | [**BlendMode**](BlendMode.md) | Blend mode of the shadow | 
**Offset** | [**Vector**](Vector.md) | How far the shadow is projected in the x and y directions | 
**Radius** | **float32** | Radius of the blur effect | 
**Spread** | Pointer to **float32** | The distance by which to expand (or contract) the shadow.  For drop shadows, a positive &#x60;spread&#x60; value creates a shadow larger than the node, whereas a negative value creates a shadow smaller than the node.  For inner shadows, a positive &#x60;spread&#x60; value contracts the shadow. Spread values are only accepted on rectangles and ellipses, or on frames, components, and instances with visible fill paints and &#x60;clipsContent&#x60; enabled. When left unspecified, the default value is 0. | [optional] [default to 0]
**Visible** | **bool** | Whether this blur is active. | 
**BoundVariables** | Pointer to [**BlurEffectBoundVariables**](BlurEffectBoundVariables.md) |  | [optional] 

## Methods

### NewEffect

`func NewEffect(type_ string, showShadowBehindNode bool, color RGBA, blendMode BlendMode, offset Vector, radius float32, visible bool, ) *Effect`

NewEffect instantiates a new Effect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEffectWithDefaults

`func NewEffectWithDefaults() *Effect`

NewEffectWithDefaults instantiates a new Effect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Effect) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Effect) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Effect) SetType(v string)`

SetType sets Type field to given value.


### GetShowShadowBehindNode

`func (o *Effect) GetShowShadowBehindNode() bool`

GetShowShadowBehindNode returns the ShowShadowBehindNode field if non-nil, zero value otherwise.

### GetShowShadowBehindNodeOk

`func (o *Effect) GetShowShadowBehindNodeOk() (*bool, bool)`

GetShowShadowBehindNodeOk returns a tuple with the ShowShadowBehindNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowShadowBehindNode

`func (o *Effect) SetShowShadowBehindNode(v bool)`

SetShowShadowBehindNode sets ShowShadowBehindNode field to given value.


### GetColor

`func (o *Effect) GetColor() RGBA`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Effect) GetColorOk() (*RGBA, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Effect) SetColor(v RGBA)`

SetColor sets Color field to given value.


### GetBlendMode

`func (o *Effect) GetBlendMode() BlendMode`

GetBlendMode returns the BlendMode field if non-nil, zero value otherwise.

### GetBlendModeOk

`func (o *Effect) GetBlendModeOk() (*BlendMode, bool)`

GetBlendModeOk returns a tuple with the BlendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlendMode

`func (o *Effect) SetBlendMode(v BlendMode)`

SetBlendMode sets BlendMode field to given value.


### GetOffset

`func (o *Effect) GetOffset() Vector`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Effect) GetOffsetOk() (*Vector, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Effect) SetOffset(v Vector)`

SetOffset sets Offset field to given value.


### GetRadius

`func (o *Effect) GetRadius() float32`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *Effect) GetRadiusOk() (*float32, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *Effect) SetRadius(v float32)`

SetRadius sets Radius field to given value.


### GetSpread

`func (o *Effect) GetSpread() float32`

GetSpread returns the Spread field if non-nil, zero value otherwise.

### GetSpreadOk

`func (o *Effect) GetSpreadOk() (*float32, bool)`

GetSpreadOk returns a tuple with the Spread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpread

`func (o *Effect) SetSpread(v float32)`

SetSpread sets Spread field to given value.

### HasSpread

`func (o *Effect) HasSpread() bool`

HasSpread returns a boolean if a field has been set.

### GetVisible

`func (o *Effect) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *Effect) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *Effect) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetBoundVariables

`func (o *Effect) GetBoundVariables() BlurEffectBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *Effect) GetBoundVariablesOk() (*BlurEffectBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *Effect) SetBoundVariables(v BlurEffectBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *Effect) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


