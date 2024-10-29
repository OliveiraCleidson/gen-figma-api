# ArcData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartingAngle** | **float32** | Start of the sweep in radians. | [default to 0]
**EndingAngle** | **float32** | End of the sweep in radians. | [default to 0]
**InnerRadius** | **float32** | Inner radius value between 0 and 1 | [default to 0]

## Methods

### NewArcData

`func NewArcData(startingAngle float32, endingAngle float32, innerRadius float32, ) *ArcData`

NewArcData instantiates a new ArcData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArcDataWithDefaults

`func NewArcDataWithDefaults() *ArcData`

NewArcDataWithDefaults instantiates a new ArcData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStartingAngle

`func (o *ArcData) GetStartingAngle() float32`

GetStartingAngle returns the StartingAngle field if non-nil, zero value otherwise.

### GetStartingAngleOk

`func (o *ArcData) GetStartingAngleOk() (*float32, bool)`

GetStartingAngleOk returns a tuple with the StartingAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingAngle

`func (o *ArcData) SetStartingAngle(v float32)`

SetStartingAngle sets StartingAngle field to given value.


### GetEndingAngle

`func (o *ArcData) GetEndingAngle() float32`

GetEndingAngle returns the EndingAngle field if non-nil, zero value otherwise.

### GetEndingAngleOk

`func (o *ArcData) GetEndingAngleOk() (*float32, bool)`

GetEndingAngleOk returns a tuple with the EndingAngle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingAngle

`func (o *ArcData) SetEndingAngle(v float32)`

SetEndingAngle sets EndingAngle field to given value.


### GetInnerRadius

`func (o *ArcData) GetInnerRadius() float32`

GetInnerRadius returns the InnerRadius field if non-nil, zero value otherwise.

### GetInnerRadiusOk

`func (o *ArcData) GetInnerRadiusOk() (*float32, bool)`

GetInnerRadiusOk returns a tuple with the InnerRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerRadius

`func (o *ArcData) SetInnerRadius(v float32)`

SetInnerRadius sets InnerRadius field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


