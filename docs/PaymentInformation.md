# PaymentInformation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The ID of the user whose payment information was queried. Can be used to verify the validity of a response. | 
**ResourceId** | **string** | The ID of the plugin, widget, or Community file that was queried. Can be used to verify the validity of a response. | 
**ResourceType** | **string** | The type of the resource. | 
**PaymentStatus** | [**PaymentStatus**](PaymentStatus.md) |  | 
**DateOfPurchase** | Pointer to **time.Time** | The UTC ISO 8601 timestamp indicating when the user purchased the resource. No value is given if the user has never purchased the resource.    Note that a value will still be returned if the user had purchased the resource, but no longer has active access to it (e.g. purchase refunded, subscription ended). | [optional] 

## Methods

### NewPaymentInformation

`func NewPaymentInformation(userId string, resourceId string, resourceType string, paymentStatus PaymentStatus, ) *PaymentInformation`

NewPaymentInformation instantiates a new PaymentInformation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInformationWithDefaults

`func NewPaymentInformationWithDefaults() *PaymentInformation`

NewPaymentInformationWithDefaults instantiates a new PaymentInformation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PaymentInformation) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PaymentInformation) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PaymentInformation) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetResourceId

`func (o *PaymentInformation) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *PaymentInformation) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *PaymentInformation) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *PaymentInformation) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *PaymentInformation) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *PaymentInformation) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetPaymentStatus

`func (o *PaymentInformation) GetPaymentStatus() PaymentStatus`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *PaymentInformation) GetPaymentStatusOk() (*PaymentStatus, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *PaymentInformation) SetPaymentStatus(v PaymentStatus)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetDateOfPurchase

`func (o *PaymentInformation) GetDateOfPurchase() time.Time`

GetDateOfPurchase returns the DateOfPurchase field if non-nil, zero value otherwise.

### GetDateOfPurchaseOk

`func (o *PaymentInformation) GetDateOfPurchaseOk() (*time.Time, bool)`

GetDateOfPurchaseOk returns a tuple with the DateOfPurchase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfPurchase

`func (o *PaymentInformation) SetDateOfPurchase(v time.Time)`

SetDateOfPurchase sets DateOfPurchase field to given value.

### HasDateOfPurchase

`func (o *PaymentInformation) HasDateOfPurchase() bool`

HasDateOfPurchase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


