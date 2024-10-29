# PaymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The current payment status of the user on the resource, as a string enum:    - &#x60;UNPAID&#x60;: user has not paid for the resource - &#x60;PAID&#x60;: user has an active purchase on the resource - &#x60;TRIAL&#x60;: user is in the trial period for a subscription resource | [optional] 

## Methods

### NewPaymentStatus

`func NewPaymentStatus() *PaymentStatus`

NewPaymentStatus instantiates a new PaymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentStatusWithDefaults

`func NewPaymentStatusWithDefaults() *PaymentStatus`

NewPaymentStatusWithDefaults instantiates a new PaymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PaymentStatus) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentStatus) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentStatus) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentStatus) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


