# WebhookFileVersionUpdatePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passcode** | **string** | The passcode specified when the webhook was created, should match what was initially provided | 
**Timestamp** | **time.Time** | UTC ISO 8601 timestamp of when the event was triggered. | 
**WebhookId** | **string** | The id of the webhook that caused the callback | 
**EventType** | **string** |  | 
**CreatedAt** | **time.Time** | UTC ISO 8601 timestamp of when the version was created | 
**Description** | Pointer to **string** | Description of the version in the version history | [optional] 
**FileKey** | **string** | The key of the file that was updated | 
**FileName** | **string** | The name of the file that was updated | 
**TriggeredBy** | [**User**](User.md) | The user that created the named version and triggered this event | 
**VersionId** | **string** | ID of the published version | 

## Methods

### NewWebhookFileVersionUpdatePayload

`func NewWebhookFileVersionUpdatePayload(passcode string, timestamp time.Time, webhookId string, eventType string, createdAt time.Time, fileKey string, fileName string, triggeredBy User, versionId string, ) *WebhookFileVersionUpdatePayload`

NewWebhookFileVersionUpdatePayload instantiates a new WebhookFileVersionUpdatePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookFileVersionUpdatePayloadWithDefaults

`func NewWebhookFileVersionUpdatePayloadWithDefaults() *WebhookFileVersionUpdatePayload`

NewWebhookFileVersionUpdatePayloadWithDefaults instantiates a new WebhookFileVersionUpdatePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasscode

`func (o *WebhookFileVersionUpdatePayload) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *WebhookFileVersionUpdatePayload) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *WebhookFileVersionUpdatePayload) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.


### GetTimestamp

`func (o *WebhookFileVersionUpdatePayload) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookFileVersionUpdatePayload) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookFileVersionUpdatePayload) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetWebhookId

`func (o *WebhookFileVersionUpdatePayload) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookFileVersionUpdatePayload) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookFileVersionUpdatePayload) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetEventType

`func (o *WebhookFileVersionUpdatePayload) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookFileVersionUpdatePayload) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookFileVersionUpdatePayload) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetCreatedAt

`func (o *WebhookFileVersionUpdatePayload) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookFileVersionUpdatePayload) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookFileVersionUpdatePayload) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *WebhookFileVersionUpdatePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookFileVersionUpdatePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookFileVersionUpdatePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookFileVersionUpdatePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileKey

`func (o *WebhookFileVersionUpdatePayload) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *WebhookFileVersionUpdatePayload) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *WebhookFileVersionUpdatePayload) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.


### GetFileName

`func (o *WebhookFileVersionUpdatePayload) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *WebhookFileVersionUpdatePayload) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *WebhookFileVersionUpdatePayload) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetTriggeredBy

`func (o *WebhookFileVersionUpdatePayload) GetTriggeredBy() User`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *WebhookFileVersionUpdatePayload) GetTriggeredByOk() (*User, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *WebhookFileVersionUpdatePayload) SetTriggeredBy(v User)`

SetTriggeredBy sets TriggeredBy field to given value.


### GetVersionId

`func (o *WebhookFileVersionUpdatePayload) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *WebhookFileVersionUpdatePayload) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *WebhookFileVersionUpdatePayload) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


