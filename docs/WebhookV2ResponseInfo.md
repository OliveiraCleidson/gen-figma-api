# WebhookV2ResponseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | HTTP status code of the response | 
**ReceivedAt** | **time.Time** | UTC ISO 8601 timestamp of when the response was received | 

## Methods

### NewWebhookV2ResponseInfo

`func NewWebhookV2ResponseInfo(status string, receivedAt time.Time, ) *WebhookV2ResponseInfo`

NewWebhookV2ResponseInfo instantiates a new WebhookV2ResponseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV2ResponseInfoWithDefaults

`func NewWebhookV2ResponseInfoWithDefaults() *WebhookV2ResponseInfo`

NewWebhookV2ResponseInfoWithDefaults instantiates a new WebhookV2ResponseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *WebhookV2ResponseInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookV2ResponseInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookV2ResponseInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReceivedAt

`func (o *WebhookV2ResponseInfo) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *WebhookV2ResponseInfo) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *WebhookV2ResponseInfo) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


