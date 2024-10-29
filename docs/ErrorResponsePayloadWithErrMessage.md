# ErrorResponsePayloadWithErrMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **float32** | Status code | 
**Err** | **string** | A string describing the error | 

## Methods

### NewErrorResponsePayloadWithErrMessage

`func NewErrorResponsePayloadWithErrMessage(status float32, err string, ) *ErrorResponsePayloadWithErrMessage`

NewErrorResponsePayloadWithErrMessage instantiates a new ErrorResponsePayloadWithErrMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponsePayloadWithErrMessageWithDefaults

`func NewErrorResponsePayloadWithErrMessageWithDefaults() *ErrorResponsePayloadWithErrMessage`

NewErrorResponsePayloadWithErrMessageWithDefaults instantiates a new ErrorResponsePayloadWithErrMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ErrorResponsePayloadWithErrMessage) GetStatus() float32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorResponsePayloadWithErrMessage) GetStatusOk() (*float32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorResponsePayloadWithErrMessage) SetStatus(v float32)`

SetStatus sets Status field to given value.


### GetErr

`func (o *ErrorResponsePayloadWithErrMessage) GetErr() string`

GetErr returns the Err field if non-nil, zero value otherwise.

### GetErrOk

`func (o *ErrorResponsePayloadWithErrMessage) GetErrOk() (*string, bool)`

GetErrOk returns a tuple with the Err field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErr

`func (o *ErrorResponsePayloadWithErrMessage) SetErr(v string)`

SetErr sets Err field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


