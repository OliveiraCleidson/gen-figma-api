# ResponseCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Before** | Pointer to **float32** |  | [optional] 
**After** | Pointer to **float32** |  | [optional] 

## Methods

### NewResponseCursor

`func NewResponseCursor() *ResponseCursor`

NewResponseCursor instantiates a new ResponseCursor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseCursorWithDefaults

`func NewResponseCursorWithDefaults() *ResponseCursor`

NewResponseCursorWithDefaults instantiates a new ResponseCursor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBefore

`func (o *ResponseCursor) GetBefore() float32`

GetBefore returns the Before field if non-nil, zero value otherwise.

### GetBeforeOk

`func (o *ResponseCursor) GetBeforeOk() (*float32, bool)`

GetBeforeOk returns a tuple with the Before field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBefore

`func (o *ResponseCursor) SetBefore(v float32)`

SetBefore sets Before field to given value.

### HasBefore

`func (o *ResponseCursor) HasBefore() bool`

HasBefore returns a boolean if a field has been set.

### GetAfter

`func (o *ResponseCursor) GetAfter() float32`

GetAfter returns the After field if non-nil, zero value otherwise.

### GetAfterOk

`func (o *ResponseCursor) GetAfterOk() (*float32, bool)`

GetAfterOk returns a tuple with the After field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfter

`func (o *ResponseCursor) SetAfter(v float32)`

SetAfter sets After field to given value.

### HasAfter

`func (o *ResponseCursor) HasAfter() bool`

HasAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


