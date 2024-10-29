# SolidPaint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visible** | Pointer to **bool** | Is the paint enabled? | [optional] [default to true]
**Opacity** | Pointer to **float32** | Overall opacity of paint (colors within the paint can also have opacity values which would blend with this) | [optional] [default to 1]
**BlendMode** | [**BlendMode**](BlendMode.md) | How this node blends with nodes behind it in the scene | 
**Type** | **string** | The string literal \&quot;SOLID\&quot; representing the paint&#39;s type. Always check the &#x60;type&#x60; before reading other properties. | 
**Color** | [**RGBA**](RGBA.md) | Solid color of the paint | 
**BoundVariables** | Pointer to [**SolidPaintAllOfBoundVariables**](SolidPaintAllOfBoundVariables.md) |  | [optional] 

## Methods

### NewSolidPaint

`func NewSolidPaint(blendMode BlendMode, type_ string, color RGBA, ) *SolidPaint`

NewSolidPaint instantiates a new SolidPaint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolidPaintWithDefaults

`func NewSolidPaintWithDefaults() *SolidPaint`

NewSolidPaintWithDefaults instantiates a new SolidPaint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisible

`func (o *SolidPaint) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *SolidPaint) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *SolidPaint) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *SolidPaint) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetOpacity

`func (o *SolidPaint) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *SolidPaint) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *SolidPaint) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *SolidPaint) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.

### GetBlendMode

`func (o *SolidPaint) GetBlendMode() BlendMode`

GetBlendMode returns the BlendMode field if non-nil, zero value otherwise.

### GetBlendModeOk

`func (o *SolidPaint) GetBlendModeOk() (*BlendMode, bool)`

GetBlendModeOk returns a tuple with the BlendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlendMode

`func (o *SolidPaint) SetBlendMode(v BlendMode)`

SetBlendMode sets BlendMode field to given value.


### GetType

`func (o *SolidPaint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SolidPaint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SolidPaint) SetType(v string)`

SetType sets Type field to given value.


### GetColor

`func (o *SolidPaint) GetColor() RGBA`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *SolidPaint) GetColorOk() (*RGBA, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *SolidPaint) SetColor(v RGBA)`

SetColor sets Color field to given value.


### GetBoundVariables

`func (o *SolidPaint) GetBoundVariables() SolidPaintAllOfBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *SolidPaint) GetBoundVariablesOk() (*SolidPaintAllOfBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *SolidPaint) SetBoundVariables(v SolidPaintAllOfBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *SolidPaint) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


