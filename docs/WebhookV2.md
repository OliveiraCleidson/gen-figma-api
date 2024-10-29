# WebhookV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the webhook | 
**EventType** | [**WebhookV2Event**](WebhookV2Event.md) | The event this webhook triggers on | 
**TeamId** | **string** | The team id you are subscribed to for updates | 
**Status** | [**WebhookV2Status**](WebhookV2Status.md) | The current status of the webhook | 
**ClientId** | **NullableString** | The client ID of the OAuth application that registered this webhook, if any | 
**Passcode** | **string** | The passcode that will be passed back to the webhook endpoint | 
**Endpoint** | **string** | The endpoint that will be hit when the webhook is triggered | 
**Description** | **NullableString** | Optional user-provided description or name for the webhook. This is provided to help make maintaining a number of webhooks more convenient. Max length 140 characters. | 

## Methods

### NewWebhookV2

`func NewWebhookV2(id string, eventType WebhookV2Event, teamId string, status WebhookV2Status, clientId NullableString, passcode string, endpoint string, description NullableString, ) *WebhookV2`

NewWebhookV2 instantiates a new WebhookV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV2WithDefaults

`func NewWebhookV2WithDefaults() *WebhookV2`

NewWebhookV2WithDefaults instantiates a new WebhookV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookV2) SetId(v string)`

SetId sets Id field to given value.


### GetEventType

`func (o *WebhookV2) GetEventType() WebhookV2Event`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookV2) GetEventTypeOk() (*WebhookV2Event, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookV2) SetEventType(v WebhookV2Event)`

SetEventType sets EventType field to given value.


### GetTeamId

`func (o *WebhookV2) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *WebhookV2) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *WebhookV2) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetStatus

`func (o *WebhookV2) GetStatus() WebhookV2Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookV2) GetStatusOk() (*WebhookV2Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookV2) SetStatus(v WebhookV2Status)`

SetStatus sets Status field to given value.


### GetClientId

`func (o *WebhookV2) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *WebhookV2) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *WebhookV2) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### SetClientIdNil

`func (o *WebhookV2) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *WebhookV2) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetPasscode

`func (o *WebhookV2) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *WebhookV2) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *WebhookV2) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.


### GetEndpoint

`func (o *WebhookV2) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *WebhookV2) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *WebhookV2) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetDescription

`func (o *WebhookV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookV2) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *WebhookV2) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WebhookV2) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


