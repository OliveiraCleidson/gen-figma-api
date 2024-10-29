# WebhookFileCommentPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passcode** | **string** | The passcode specified when the webhook was created, should match what was initially provided | 
**Timestamp** | **time.Time** | UTC ISO 8601 timestamp of when the event was triggered. | 
**WebhookId** | **string** | The id of the webhook that caused the callback | 
**EventType** | **string** |  | 
**Comment** | [**[]CommentFragment**](CommentFragment.md) | Contents of the comment itself | 
**CommentId** | **string** | Unique identifier for comment | 
**CreatedAt** | **time.Time** | The UTC ISO 8601 time at which the comment was left | 
**FileKey** | **string** | The key of the file that was commented on | 
**FileName** | **string** | The name of the file that was commented on | 
**Mentions** | Pointer to [**[]User**](User.md) | Users that were mentioned in the comment | [optional] 
**TriggeredBy** | [**User**](User.md) | The user that made the comment and triggered this event | 

## Methods

### NewWebhookFileCommentPayload

`func NewWebhookFileCommentPayload(passcode string, timestamp time.Time, webhookId string, eventType string, comment []CommentFragment, commentId string, createdAt time.Time, fileKey string, fileName string, triggeredBy User, ) *WebhookFileCommentPayload`

NewWebhookFileCommentPayload instantiates a new WebhookFileCommentPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookFileCommentPayloadWithDefaults

`func NewWebhookFileCommentPayloadWithDefaults() *WebhookFileCommentPayload`

NewWebhookFileCommentPayloadWithDefaults instantiates a new WebhookFileCommentPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasscode

`func (o *WebhookFileCommentPayload) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *WebhookFileCommentPayload) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *WebhookFileCommentPayload) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.


### GetTimestamp

`func (o *WebhookFileCommentPayload) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookFileCommentPayload) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookFileCommentPayload) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetWebhookId

`func (o *WebhookFileCommentPayload) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookFileCommentPayload) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookFileCommentPayload) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetEventType

`func (o *WebhookFileCommentPayload) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookFileCommentPayload) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookFileCommentPayload) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetComment

`func (o *WebhookFileCommentPayload) GetComment() []CommentFragment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *WebhookFileCommentPayload) GetCommentOk() (*[]CommentFragment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *WebhookFileCommentPayload) SetComment(v []CommentFragment)`

SetComment sets Comment field to given value.


### GetCommentId

`func (o *WebhookFileCommentPayload) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *WebhookFileCommentPayload) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *WebhookFileCommentPayload) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.


### GetCreatedAt

`func (o *WebhookFileCommentPayload) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookFileCommentPayload) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookFileCommentPayload) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFileKey

`func (o *WebhookFileCommentPayload) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *WebhookFileCommentPayload) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *WebhookFileCommentPayload) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.


### GetFileName

`func (o *WebhookFileCommentPayload) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *WebhookFileCommentPayload) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *WebhookFileCommentPayload) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetMentions

`func (o *WebhookFileCommentPayload) GetMentions() []User`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *WebhookFileCommentPayload) GetMentionsOk() (*[]User, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *WebhookFileCommentPayload) SetMentions(v []User)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *WebhookFileCommentPayload) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *WebhookFileCommentPayload) GetTriggeredBy() User`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *WebhookFileCommentPayload) GetTriggeredByOk() (*User, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *WebhookFileCommentPayload) SetTriggeredBy(v User)`

SetTriggeredBy sets TriggeredBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


