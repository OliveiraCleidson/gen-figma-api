# ErrorResponsePayloadWithErrorBoolean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **bool** | For erroneous requests, this value is always &#x60;true&#x60;. | 
**Status** | **float32** | Status code | 
**Message** | **string** | A string describing the error | 

## Methods

### NewErrorResponsePayloadWithErrorBoolean

`func NewErrorResponsePayloadWithErrorBoolean(error_ bool, status float32, message string, ) *ErrorResponsePayloadWithErrorBoolean`

NewErrorResponsePayloadWithErrorBoolean instantiates a new ErrorResponsePayloadWithErrorBoolean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponsePayloadWithErrorBooleanWithDefaults

`func NewErrorResponsePayloadWithErrorBooleanWithDefaults() *ErrorResponsePayloadWithErrorBoolean`

NewErrorResponsePayloadWithErrorBooleanWithDefaults instantiates a new ErrorResponsePayloadWithErrorBoolean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorResponsePayloadWithErrorBoolean) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorResponsePayloadWithErrorBoolean) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorResponsePayloadWithErrorBoolean) SetError(v bool)`

SetError sets Error field to given value.


### GetStatus

`func (o *ErrorResponsePayloadWithErrorBoolean) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorResponsePayloadWithErrorBoolean) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorResponsePayloadWithErrorBoolean) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetMessage

`func (o *ErrorResponsePayloadWithErrorBoolean) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponsePayloadWithErrorBoolean) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponsePayloadWithErrorBoolean) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


