# PostWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**WebhookV2Event**](WebhookV2Event.md) |  | 
**TeamId** | **string** | Team id to receive updates about | 
**Endpoint** | **string** | The HTTP endpoint that will receive a POST request when the event triggers. Max length 2048 characters. | 
**Passcode** | **string** | String that will be passed back to your webhook endpoint to verify that it is being called by Figma. Max length 100 characters. | 
**Status** | Pointer to [**WebhookV2Status**](WebhookV2Status.md) | State of the webhook, including any error state it may be in | [optional] 
**Description** | Pointer to **string** | User provided description or name for the webhook. Max length 150 characters. | [optional] 

## Methods

### NewPostWebhookRequest

`func NewPostWebhookRequest(eventType WebhookV2Event, teamId string, endpoint string, passcode string, ) *PostWebhookRequest`

NewPostWebhookRequest instantiates a new PostWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWebhookRequestWithDefaults

`func NewPostWebhookRequestWithDefaults() *PostWebhookRequest`

NewPostWebhookRequestWithDefaults instantiates a new PostWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *PostWebhookRequest) GetEventType() WebhookV2Event`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PostWebhookRequest) GetEventTypeOk() (*WebhookV2Event, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PostWebhookRequest) SetEventType(v WebhookV2Event)`

SetEventType sets EventType field to given value.


### GetTeamId

`func (o *PostWebhookRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *PostWebhookRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *PostWebhookRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetEndpoint

`func (o *PostWebhookRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PostWebhookRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PostWebhookRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetPasscode

`func (o *PostWebhookRequest) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *PostWebhookRequest) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *PostWebhookRequest) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.


### GetStatus

`func (o *PostWebhookRequest) GetStatus() WebhookV2Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostWebhookRequest) GetStatusOk() (*WebhookV2Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostWebhookRequest) SetStatus(v WebhookV2Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostWebhookRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *PostWebhookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostWebhookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostWebhookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostWebhookRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


