# WebhookFileDeletePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passcode** | **string** | The passcode specified when the webhook was created, should match what was initially provided | 
**Timestamp** | **time.Time** | UTC ISO 8601 timestamp of when the event was triggered. | 
**WebhookId** | **string** | The id of the webhook that caused the callback | 
**EventType** | **string** |  | 
**FileKey** | **string** | The key of the file that was deleted | 
**FileName** | **string** | The name of the file that was deleted | 
**TriggeredBy** | [**User**](User.md) | The user that deleted the file and triggered this event | 

## Methods

### NewWebhookFileDeletePayload

`func NewWebhookFileDeletePayload(passcode string, timestamp time.Time, webhookId string, eventType string, fileKey string, fileName string, triggeredBy User, ) *WebhookFileDeletePayload`

NewWebhookFileDeletePayload instantiates a new WebhookFileDeletePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookFileDeletePayloadWithDefaults

`func NewWebhookFileDeletePayloadWithDefaults() *WebhookFileDeletePayload`

NewWebhookFileDeletePayloadWithDefaults instantiates a new WebhookFileDeletePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasscode

`func (o *WebhookFileDeletePayload) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *WebhookFileDeletePayload) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *WebhookFileDeletePayload) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.


### GetTimestamp

`func (o *WebhookFileDeletePayload) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookFileDeletePayload) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookFileDeletePayload) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetWebhookId

`func (o *WebhookFileDeletePayload) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookFileDeletePayload) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookFileDeletePayload) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetEventType

`func (o *WebhookFileDeletePayload) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookFileDeletePayload) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookFileDeletePayload) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetFileKey

`func (o *WebhookFileDeletePayload) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *WebhookFileDeletePayload) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *WebhookFileDeletePayload) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.


### GetFileName

`func (o *WebhookFileDeletePayload) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *WebhookFileDeletePayload) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *WebhookFileDeletePayload) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetTriggeredBy

`func (o *WebhookFileDeletePayload) GetTriggeredBy() User`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *WebhookFileDeletePayload) GetTriggeredByOk() (*User, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *WebhookFileDeletePayload) SetTriggeredBy(v User)`

SetTriggeredBy sets TriggeredBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


