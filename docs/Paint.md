# Paint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The string literal \&quot;SOLID\&quot; representing the paint&#39;s type. Always check the &#x60;type&#x60; before reading other properties. | 
**Color** | [**RGBA**](RGBA.md) | Solid color of the paint | 
**BoundVariables** | Pointer to [**SolidPaintAllOfBoundVariables**](SolidPaintAllOfBoundVariables.md) |  | [optional] 
**Visible** | Pointer to **bool** | Is the paint enabled? | [optional] [default to true]
**Opacity** | Pointer to **float32** | Overall opacity of paint (colors within the paint can also have opacity values which would blend with this) | [optional] [default to 1]
**BlendMode** | [**BlendMode**](BlendMode.md) | How this node blends with nodes behind it in the scene | 
**GradientHandlePositions** | [**[]Vector**](Vector.md) | This field contains three vectors, each of which are a position in normalized object space (normalized object space is if the top left corner of the bounding box of the object is (0, 0) and the bottom right is (1,1)). The first position corresponds to the start of the gradient (value 0 for the purposes of calculating gradient stops), the second position is the end of the gradient (value 1), and the third handle position determines the width of the gradient. | 
**GradientStops** | [**[]ColorStop**](ColorStop.md) | Positions of key points along the gradient axis with the colors anchored there. Colors along the gradient are interpolated smoothly between neighboring gradient stops. | 
**ScaleMode** | **string** | Image scaling mode. | 
**ImageRef** | **string** | A reference to an image embedded in this node. To download the image using this reference, use the &#x60;GET file images&#x60; endpoint to retrieve the mapping from image references to image URLs. | 
**ImageTransform** | Pointer to **[][]float32** | A transformation matrix is standard way in computer graphics to represent translation and rotation. These are the top two rows of a 3x3 matrix. The bottom row of the matrix is assumed to be [0, 0, 1]. This is known as an affine transform and is enough to represent translation, rotation, and skew.  The identity transform is [[1, 0, 0], [0, 1, 0]].  A translation matrix will typically look like:  &#x60;&#x60;&#x60; [[1, 0, tx],   [0, 1, ty]] &#x60;&#x60;&#x60;  and a rotation matrix will typically look like:  &#x60;&#x60;&#x60; [[cos(angle), sin(angle), 0],   [-sin(angle), cos(angle), 0]] &#x60;&#x60;&#x60;  Another way to think about this transform is as three vectors:  - The x axis (t[0][0], t[1][0]) - The y axis (t[0][1], t[1][1]) - The translation offset (t[0][2], t[1][2])  The most common usage of the Transform matrix is the &#x60;relativeTransform property&#x60;. This particular usage of the matrix has a few additional restrictions. The translation offset can take on any value but we do enforce that the axis vectors are unit vectors (i.e. have length 1). The axes are not required to be at 90° angles to each other. | [optional] 
**ScalingFactor** | Pointer to **float32** | Amount image is scaled by in tiling, only present if scaleMode is &#x60;TILE&#x60;. | [optional] 
**Filters** | Pointer to [**ImageFilters**](ImageFilters.md) | Defines what image filters have been applied to this paint, if any. If this property is not defined, no filters have been applied. | [optional] 
**Rotation** | Pointer to **float32** | Image rotation, in degrees. | [optional] [default to 0]
**GifRef** | Pointer to **string** | A reference to an animated GIF embedded in this node. To download the image using this reference, use the &#x60;GET file images&#x60; endpoint to retrieve the mapping from image references to image URLs. | [optional] 

## Methods

### NewPaint

`func NewPaint(type_ string, color RGBA, blendMode BlendMode, gradientHandlePositions []Vector, gradientStops []ColorStop, scaleMode string, imageRef string, ) *Paint`

NewPaint instantiates a new Paint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaintWithDefaults

`func NewPaintWithDefaults() *Paint`

NewPaintWithDefaults instantiates a new Paint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Paint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Paint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Paint) SetType(v string)`

SetType sets Type field to given value.


### GetColor

`func (o *Paint) GetColor() RGBA`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Paint) GetColorOk() (*RGBA, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Paint) SetColor(v RGBA)`

SetColor sets Color field to given value.


### GetBoundVariables

`func (o *Paint) GetBoundVariables() SolidPaintAllOfBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *Paint) GetBoundVariablesOk() (*SolidPaintAllOfBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *Paint) SetBoundVariables(v SolidPaintAllOfBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *Paint) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.

### GetVisible

`func (o *Paint) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *Paint) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *Paint) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *Paint) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetOpacity

`func (o *Paint) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *Paint) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *Paint) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *Paint) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.

### GetBlendMode

`func (o *Paint) GetBlendMode() BlendMode`

GetBlendMode returns the BlendMode field if non-nil, zero value otherwise.

### GetBlendModeOk

`func (o *Paint) GetBlendModeOk() (*BlendMode, bool)`

GetBlendModeOk returns a tuple with the BlendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlendMode

`func (o *Paint) SetBlendMode(v BlendMode)`

SetBlendMode sets BlendMode field to given value.


### GetGradientHandlePositions

`func (o *Paint) GetGradientHandlePositions() []Vector`

GetGradientHandlePositions returns the GradientHandlePositions field if non-nil, zero value otherwise.

### GetGradientHandlePositionsOk

`func (o *Paint) GetGradientHandlePositionsOk() (*[]Vector, bool)`

GetGradientHandlePositionsOk returns a tuple with the GradientHandlePositions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradientHandlePositions

`func (o *Paint) SetGradientHandlePositions(v []Vector)`

SetGradientHandlePositions sets GradientHandlePositions field to given value.


### GetGradientStops

`func (o *Paint) GetGradientStops() []ColorStop`

GetGradientStops returns the GradientStops field if non-nil, zero value otherwise.

### GetGradientStopsOk

`func (o *Paint) GetGradientStopsOk() (*[]ColorStop, bool)`

GetGradientStopsOk returns a tuple with the GradientStops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradientStops

`func (o *Paint) SetGradientStops(v []ColorStop)`

SetGradientStops sets GradientStops field to given value.


### GetScaleMode

`func (o *Paint) GetScaleMode() string`

GetScaleMode returns the ScaleMode field if non-nil, zero value otherwise.

### GetScaleModeOk

`func (o *Paint) GetScaleModeOk() (*string, bool)`

GetScaleModeOk returns a tuple with the ScaleMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScaleMode

`func (o *Paint) SetScaleMode(v string)`

SetScaleMode sets ScaleMode field to given value.


### GetImageRef

`func (o *Paint) GetImageRef() string`

GetImageRef returns the ImageRef field if non-nil, zero value otherwise.

### GetImageRefOk

`func (o *Paint) GetImageRefOk() (*string, bool)`

GetImageRefOk returns a tuple with the ImageRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageRef

`func (o *Paint) SetImageRef(v string)`

SetImageRef sets ImageRef field to given value.


### GetImageTransform

`func (o *Paint) GetImageTransform() [][]float32`

GetImageTransform returns the ImageTransform field if non-nil, zero value otherwise.

### GetImageTransformOk

`func (o *Paint) GetImageTransformOk() (*[][]float32, bool)`

GetImageTransformOk returns a tuple with the ImageTransform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTransform

`func (o *Paint) SetImageTransform(v [][]float32)`

SetImageTransform sets ImageTransform field to given value.

### HasImageTransform

`func (o *Paint) HasImageTransform() bool`

HasImageTransform returns a boolean if a field has been set.

### GetScalingFactor

`func (o *Paint) GetScalingFactor() float32`

GetScalingFactor returns the ScalingFactor field if non-nil, zero value otherwise.

### GetScalingFactorOk

`func (o *Paint) GetScalingFactorOk() (*float32, bool)`

GetScalingFactorOk returns a tuple with the ScalingFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalingFactor

`func (o *Paint) SetScalingFactor(v float32)`

SetScalingFactor sets ScalingFactor field to given value.

### HasScalingFactor

`func (o *Paint) HasScalingFactor() bool`

HasScalingFactor returns a boolean if a field has been set.

### GetFilters

`func (o *Paint) GetFilters() ImageFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *Paint) GetFiltersOk() (*ImageFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *Paint) SetFilters(v ImageFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *Paint) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetRotation

`func (o *Paint) GetRotation() float32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *Paint) GetRotationOk() (*float32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *Paint) SetRotation(v float32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *Paint) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetGifRef

`func (o *Paint) GetGifRef() string`

GetGifRef returns the GifRef field if non-nil, zero value otherwise.

### GetGifRefOk

`func (o *Paint) GetGifRefOk() (*string, bool)`

GetGifRefOk returns a tuple with the GifRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifRef

`func (o *Paint) SetGifRef(v string)`

SetGifRef sets GifRef field to given value.

### HasGifRef

`func (o *Paint) HasGifRef() bool`

HasGifRef returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


