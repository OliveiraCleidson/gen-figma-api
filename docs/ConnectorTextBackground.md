# ConnectorTextBackground

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CornerRadius** | Pointer to **float32** | Radius of each corner if a single radius is set for all corners | [optional] [default to 0]
**CornerSmoothing** | Pointer to **float32** | A value that lets you control how \&quot;smooth\&quot; the corners are. Ranges from 0 to 1. 0 is the default and means that the corner is perfectly circular. A value of 0.6 means the corner matches the iOS 7 \&quot;squircle\&quot; icon shape. Other values produce various other curves. | [optional] 
**RectangleCornerRadii** | Pointer to **[]float32** | Array of length 4 of the radius of each corner of the frame, starting in the top left and proceeding clockwise.  Values are given in the order top-left, top-right, bottom-right, bottom-left. | [optional] 
**Fills** | [**[]Paint**](Paint.md) | An array of fill paints applied to the node. | 
**Styles** | Pointer to **map[string]string** | A mapping of a StyleType to style ID (see Style) of styles present on this node. The style ID can be used to look up more information about the style in the top-level styles field. | [optional] 

## Methods

### NewConnectorTextBackground

`func NewConnectorTextBackground(fills []Paint, ) *ConnectorTextBackground`

NewConnectorTextBackground instantiates a new ConnectorTextBackground object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorTextBackgroundWithDefaults

`func NewConnectorTextBackgroundWithDefaults() *ConnectorTextBackground`

NewConnectorTextBackgroundWithDefaults instantiates a new ConnectorTextBackground object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCornerRadius

`func (o *ConnectorTextBackground) GetCornerRadius() float32`

GetCornerRadius returns the CornerRadius field if non-nil, zero value otherwise.

### GetCornerRadiusOk

`func (o *ConnectorTextBackground) GetCornerRadiusOk() (*float32, bool)`

GetCornerRadiusOk returns a tuple with the CornerRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerRadius

`func (o *ConnectorTextBackground) SetCornerRadius(v float32)`

SetCornerRadius sets CornerRadius field to given value.

### HasCornerRadius

`func (o *ConnectorTextBackground) HasCornerRadius() bool`

HasCornerRadius returns a boolean if a field has been set.

### GetCornerSmoothing

`func (o *ConnectorTextBackground) GetCornerSmoothing() float32`

GetCornerSmoothing returns the CornerSmoothing field if non-nil, zero value otherwise.

### GetCornerSmoothingOk

`func (o *ConnectorTextBackground) GetCornerSmoothingOk() (*float32, bool)`

GetCornerSmoothingOk returns a tuple with the CornerSmoothing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCornerSmoothing

`func (o *ConnectorTextBackground) SetCornerSmoothing(v float32)`

SetCornerSmoothing sets CornerSmoothing field to given value.

### HasCornerSmoothing

`func (o *ConnectorTextBackground) HasCornerSmoothing() bool`

HasCornerSmoothing returns a boolean if a field has been set.

### GetRectangleCornerRadii

`func (o *ConnectorTextBackground) GetRectangleCornerRadii() []float32`

GetRectangleCornerRadii returns the RectangleCornerRadii field if non-nil, zero value otherwise.

### GetRectangleCornerRadiiOk

`func (o *ConnectorTextBackground) GetRectangleCornerRadiiOk() (*[]float32, bool)`

GetRectangleCornerRadiiOk returns a tuple with the RectangleCornerRadii field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRectangleCornerRadii

`func (o *ConnectorTextBackground) SetRectangleCornerRadii(v []float32)`

SetRectangleCornerRadii sets RectangleCornerRadii field to given value.

### HasRectangleCornerRadii

`func (o *ConnectorTextBackground) HasRectangleCornerRadii() bool`

HasRectangleCornerRadii returns a boolean if a field has been set.

### GetFills

`func (o *ConnectorTextBackground) GetFills() []Paint`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *ConnectorTextBackground) GetFillsOk() (*[]Paint, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *ConnectorTextBackground) SetFills(v []Paint)`

SetFills sets Fills field to given value.


### GetStyles

`func (o *ConnectorTextBackground) GetStyles() map[string]string`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *ConnectorTextBackground) GetStylesOk() (*map[string]string, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *ConnectorTextBackground) SetStyles(v map[string]string)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *ConnectorTextBackground) HasStyles() bool`

HasStyles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


