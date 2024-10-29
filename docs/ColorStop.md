# ColorStop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Position** | **float32** | Value between 0 and 1 representing position along gradient axis. | 
**Color** | [**RGBA**](RGBA.md) | Color attached to corresponding position. | 
**BoundVariables** | Pointer to [**ColorStopBoundVariables**](ColorStopBoundVariables.md) |  | [optional] 

## Methods

### NewColorStop

`func NewColorStop(position float32, color RGBA, ) *ColorStop`

NewColorStop instantiates a new ColorStop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColorStopWithDefaults

`func NewColorStopWithDefaults() *ColorStop`

NewColorStopWithDefaults instantiates a new ColorStop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPosition

`func (o *ColorStop) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ColorStop) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ColorStop) SetPosition(v float32)`

SetPosition sets Position field to given value.


### GetColor

`func (o *ColorStop) GetColor() RGBA`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ColorStop) GetColorOk() (*RGBA, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ColorStop) SetColor(v RGBA)`

SetColor sets Color field to given value.


### GetBoundVariables

`func (o *ColorStop) GetBoundVariables() ColorStopBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *ColorStop) GetBoundVariablesOk() (*ColorStopBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *ColorStop) SetBoundVariables(v ColorStopBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *ColorStop) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


