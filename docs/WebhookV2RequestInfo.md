# WebhookV2RequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the webhook | 
**Endpoint** | **string** | The actual endpoint the request was sent to | 
**Payload** | **map[string]interface{}** | The contents of the request that was sent to the endpoint | 
**SentAt** | **time.Time** | UTC ISO 8601 timestamp of when the request was sent | 

## Methods

### NewWebhookV2RequestInfo

`func NewWebhookV2RequestInfo(id string, endpoint string, payload map[string]interface{}, sentAt time.Time, ) *WebhookV2RequestInfo`

NewWebhookV2RequestInfo instantiates a new WebhookV2RequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV2RequestInfoWithDefaults

`func NewWebhookV2RequestInfoWithDefaults() *WebhookV2RequestInfo`

NewWebhookV2RequestInfoWithDefaults instantiates a new WebhookV2RequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookV2RequestInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookV2RequestInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookV2RequestInfo) SetId(v string)`

SetId sets Id field to given value.


### GetEndpoint

`func (o *WebhookV2RequestInfo) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *WebhookV2RequestInfo) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *WebhookV2RequestInfo) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetPayload

`func (o *WebhookV2RequestInfo) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhookV2RequestInfo) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebhookV2RequestInfo) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.


### GetSentAt

`func (o *WebhookV2RequestInfo) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *WebhookV2RequestInfo) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *WebhookV2RequestInfo) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


