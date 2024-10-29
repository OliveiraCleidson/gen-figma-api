# GradientPaint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visible** | Pointer to **bool** | Is the paint enabled? | [optional] [default to true]
**Opacity** | Pointer to **float32** | Overall opacity of paint (colors within the paint can also have opacity values which would blend with this) | [optional] [default to 1]
**BlendMode** | [**BlendMode**](BlendMode.md) | How this node blends with nodes behind it in the scene | 
**Type** | **string** | The string literal representing the paint&#39;s type. Always check the &#x60;type&#x60; before reading other properties. | 
**GradientHandlePositions** | [**[]Vector**](Vector.md) | This field contains three vectors, each of which are a position in normalized object space (normalized object space is if the top left corner of the bounding box of the object is (0, 0) and the bottom right is (1,1)). The first position corresponds to the start of the gradient (value 0 for the purposes of calculating gradient stops), the second position is the end of the gradient (value 1), and the third handle position determines the width of the gradient. | 
**GradientStops** | [**[]ColorStop**](ColorStop.md) | Positions of key points along the gradient axis with the colors anchored there. Colors along the gradient are interpolated smoothly between neighboring gradient stops. | 

## Methods

### NewGradientPaint

`func NewGradientPaint(blendMode BlendMode, type_ string, gradientHandlePositions []Vector, gradientStops []ColorStop, ) *GradientPaint`

NewGradientPaint instantiates a new GradientPaint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGradientPaintWithDefaults

`func NewGradientPaintWithDefaults() *GradientPaint`

NewGradientPaintWithDefaults instantiates a new GradientPaint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisible

`func (o *GradientPaint) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *GradientPaint) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *GradientPaint) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *GradientPaint) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetOpacity

`func (o *GradientPaint) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *GradientPaint) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *GradientPaint) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *GradientPaint) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.

### GetBlendMode

`func (o *GradientPaint) GetBlendMode() BlendMode`

GetBlendMode returns the BlendMode field if non-nil, zero value otherwise.

### GetBlendModeOk

`func (o *GradientPaint) GetBlendModeOk() (*BlendMode, bool)`

GetBlendModeOk returns a tuple with the BlendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlendMode

`func (o *GradientPaint) SetBlendMode(v BlendMode)`

SetBlendMode sets BlendMode field to given value.


### GetType

`func (o *GradientPaint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GradientPaint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GradientPaint) SetType(v string)`

SetType sets Type field to given value.


### GetGradientHandlePositions

`func (o *GradientPaint) GetGradientHandlePositions() []Vector`

GetGradientHandlePositions returns the GradientHandlePositions field if non-nil, zero value otherwise.

### GetGradientHandlePositionsOk

`func (o *GradientPaint) GetGradientHandlePositionsOk() (*[]Vector, bool)`

GetGradientHandlePositionsOk returns a tuple with the GradientHandlePositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradientHandlePositions

`func (o *GradientPaint) SetGradientHandlePositions(v []Vector)`

SetGradientHandlePositions sets GradientHandlePositions field to given value.


### GetGradientStops

`func (o *GradientPaint) GetGradientStops() []ColorStop`

GetGradientStops returns the GradientStops field if non-nil, zero value otherwise.

### GetGradientStopsOk

`func (o *GradientPaint) GetGradientStopsOk() (*[]ColorStop, bool)`

GetGradientStopsOk returns a tuple with the GradientStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradientStops

`func (o *GradientPaint) SetGradientStops(v []ColorStop)`

SetGradientStops sets GradientStops field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


