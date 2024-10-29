# Measurement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Start** | [**MeasurementStartEnd**](MeasurementStartEnd.md) |  | 
**End** | [**MeasurementStartEnd**](MeasurementStartEnd.md) |  | 
**Offset** | [**MeasurementOffset**](MeasurementOffset.md) |  | 
**FreeText** | Pointer to **string** | When manually overridden, the displayed value of the measurement | [optional] 

## Methods

### NewMeasurement

`func NewMeasurement(id string, start MeasurementStartEnd, end MeasurementStartEnd, offset MeasurementOffset, ) *Measurement`

NewMeasurement instantiates a new Measurement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeasurementWithDefaults

`func NewMeasurementWithDefaults() *Measurement`

NewMeasurementWithDefaults instantiates a new Measurement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Measurement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Measurement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Measurement) SetId(v string)`

SetId sets Id field to given value.


### GetStart

`func (o *Measurement) GetStart() MeasurementStartEnd`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Measurement) GetStartOk() (*MeasurementStartEnd, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Measurement) SetStart(v MeasurementStartEnd)`

SetStart sets Start field to given value.


### GetEnd

`func (o *Measurement) GetEnd() MeasurementStartEnd`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Measurement) GetEndOk() (*MeasurementStartEnd, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Measurement) SetEnd(v MeasurementStartEnd)`

SetEnd sets End field to given value.


### GetOffset

`func (o *Measurement) GetOffset() MeasurementOffset`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *Measurement) GetOffsetOk() (*MeasurementOffset, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *Measurement) SetOffset(v MeasurementOffset)`

SetOffset sets Offset field to given value.


### GetFreeText

`func (o *Measurement) GetFreeText() string`

GetFreeText returns the FreeText field if non-nil, zero value otherwise.

### GetFreeTextOk

`func (o *Measurement) GetFreeTextOk() (*string, bool)`

GetFreeTextOk returns a tuple with the FreeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeText

`func (o *Measurement) SetFreeText(v string)`

SetFreeText sets FreeText field to given value.

### HasFreeText

`func (o *Measurement) HasFreeText() bool`

HasFreeText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


