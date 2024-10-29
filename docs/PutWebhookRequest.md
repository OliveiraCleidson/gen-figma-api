# PutWebhookRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | [**WebhookV2Event**](WebhookV2Event.md) |  | 
**Endpoint** | **string** | The HTTP endpoint that will receive a POST request when the event triggers. Max length 2048 characters. | 
**Passcode** | **string** | String that will be passed back to your webhook endpoint to verify that it is being called by Figma. Max length 100 characters. | 
**Status** | Pointer to [**WebhookV2Status**](WebhookV2Status.md) | State of the webhook, including any error state it may be in | [optional] 
**Description** | Pointer to **string** | User provided description or name for the webhook. Max length 150 characters. | [optional] 

## Methods

### NewPutWebhookRequest

`func NewPutWebhookRequest(eventType WebhookV2Event, endpoint string, passcode string, ) *PutWebhookRequest`

NewPutWebhookRequest instantiates a new PutWebhookRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutWebhookRequestWithDefaults

`func NewPutWebhookRequestWithDefaults() *PutWebhookRequest`

NewPutWebhookRequestWithDefaults instantiates a new PutWebhookRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *PutWebhookRequest) GetEventType() WebhookV2Event`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *PutWebhookRequest) GetEventTypeOk() (*WebhookV2Event, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *PutWebhookRequest) SetEventType(v WebhookV2Event)`

SetEventType sets EventType field to given value.


### GetEndpoint

`func (o *PutWebhookRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *PutWebhookRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *PutWebhookRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetPasscode

`func (o *PutWebhookRequest) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *PutWebhookRequest) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *PutWebhookRequest) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.


### GetStatus

`func (o *PutWebhookRequest) GetStatus() WebhookV2Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PutWebhookRequest) GetStatusOk() (*WebhookV2Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PutWebhookRequest) SetStatus(v WebhookV2Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PutWebhookRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDescription

`func (o *PutWebhookRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutWebhookRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutWebhookRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PutWebhookRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


