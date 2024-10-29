# PaintOverride

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fills** | Pointer to [**[]Paint**](Paint.md) | Paints applied to characters. | [optional] 
**InheritFillStyleId** | Pointer to **string** | ID of style node, if any, that this inherits fill data from. | [optional] 

## Methods

### NewPaintOverride

`func NewPaintOverride() *PaintOverride`

NewPaintOverride instantiates a new PaintOverride object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaintOverrideWithDefaults

`func NewPaintOverrideWithDefaults() *PaintOverride`

NewPaintOverrideWithDefaults instantiates a new PaintOverride object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFills

`func (o *PaintOverride) GetFills() []Paint`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *PaintOverride) GetFillsOk() (*[]Paint, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *PaintOverride) SetFills(v []Paint)`

SetFills sets Fills field to given value.

### HasFills

`func (o *PaintOverride) HasFills() bool`

HasFills returns a boolean if a field has been set.

### GetInheritFillStyleId

`func (o *PaintOverride) GetInheritFillStyleId() string`

GetInheritFillStyleId returns the InheritFillStyleId field if non-nil, zero value otherwise.

### GetInheritFillStyleIdOk

`func (o *PaintOverride) GetInheritFillStyleIdOk() (*string, bool)`

GetInheritFillStyleIdOk returns a tuple with the InheritFillStyleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritFillStyleId

`func (o *PaintOverride) SetInheritFillStyleId(v string)`

SetInheritFillStyleId sets InheritFillStyleId field to given value.

### HasInheritFillStyleId

`func (o *PaintOverride) HasInheritFillStyleId() bool`

HasInheritFillStyleId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


