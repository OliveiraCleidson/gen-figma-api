# WebhookPingPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passcode** | **string** | The passcode specified when the webhook was created, should match what was initially provided | 
**Timestamp** | **time.Time** | UTC ISO 8601 timestamp of when the event was triggered. | 
**WebhookId** | **string** | The id of the webhook that caused the callback | 
**EventType** | **string** |  | 

## Methods

### NewWebhookPingPayload

`func NewWebhookPingPayload(passcode string, timestamp time.Time, webhookId string, eventType string, ) *WebhookPingPayload`

NewWebhookPingPayload instantiates a new WebhookPingPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPingPayloadWithDefaults

`func NewWebhookPingPayloadWithDefaults() *WebhookPingPayload`

NewWebhookPingPayloadWithDefaults instantiates a new WebhookPingPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasscode

`func (o *WebhookPingPayload) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *WebhookPingPayload) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *WebhookPingPayload) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.


### GetTimestamp

`func (o *WebhookPingPayload) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPingPayload) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPingPayload) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetWebhookId

`func (o *WebhookPingPayload) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookPingPayload) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookPingPayload) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetEventType

`func (o *WebhookPingPayload) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookPingPayload) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookPingPayload) SetEventType(v string)`

SetEventType sets EventType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


