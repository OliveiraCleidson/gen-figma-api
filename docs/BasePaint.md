# BasePaint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Visible** | Pointer to **bool** | Is the paint enabled? | [optional] [default to true]
**Opacity** | Pointer to **float32** | Overall opacity of paint (colors within the paint can also have opacity values which would blend with this) | [optional] [default to 1]
**BlendMode** | [**BlendMode**](BlendMode.md) | How this node blends with nodes behind it in the scene | 

## Methods

### NewBasePaint

`func NewBasePaint(blendMode BlendMode, ) *BasePaint`

NewBasePaint instantiates a new BasePaint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBasePaintWithDefaults

`func NewBasePaintWithDefaults() *BasePaint`

NewBasePaintWithDefaults instantiates a new BasePaint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisible

`func (o *BasePaint) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *BasePaint) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *BasePaint) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *BasePaint) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetOpacity

`func (o *BasePaint) GetOpacity() float32`

GetOpacity returns the Opacity field if non-nil, zero value otherwise.

### GetOpacityOk

`func (o *BasePaint) GetOpacityOk() (*float32, bool)`

GetOpacityOk returns a tuple with the Opacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpacity

`func (o *BasePaint) SetOpacity(v float32)`

SetOpacity sets Opacity field to given value.

### HasOpacity

`func (o *BasePaint) HasOpacity() bool`

HasOpacity returns a boolean if a field has been set.

### GetBlendMode

`func (o *BasePaint) GetBlendMode() BlendMode`

GetBlendMode returns the BlendMode field if non-nil, zero value otherwise.

### GetBlendModeOk

`func (o *BasePaint) GetBlendModeOk() (*BlendMode, bool)`

GetBlendModeOk returns a tuple with the BlendMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlendMode

`func (o *BasePaint) SetBlendMode(v BlendMode)`

SetBlendMode sets BlendMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


