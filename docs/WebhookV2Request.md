# WebhookV2Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookId** | **string** | The ID of the webhook the requests were sent to | 
**RequestInfo** | [**WebhookV2RequestInfo**](WebhookV2RequestInfo.md) |  | 
**ResponseInfo** | [**WebhookV2ResponseInfo**](WebhookV2ResponseInfo.md) |  | 
**ErrorMsg** | **NullableString** | Error message for this request. NULL if no error occurred | 

## Methods

### NewWebhookV2Request

`func NewWebhookV2Request(webhookId string, requestInfo WebhookV2RequestInfo, responseInfo WebhookV2ResponseInfo, errorMsg NullableString, ) *WebhookV2Request`

NewWebhookV2Request instantiates a new WebhookV2Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV2RequestWithDefaults

`func NewWebhookV2RequestWithDefaults() *WebhookV2Request`

NewWebhookV2RequestWithDefaults instantiates a new WebhookV2Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookId

`func (o *WebhookV2Request) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookV2Request) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookV2Request) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetRequestInfo

`func (o *WebhookV2Request) GetRequestInfo() WebhookV2RequestInfo`

GetRequestInfo returns the RequestInfo field if non-nil, zero value otherwise.

### GetRequestInfoOk

`func (o *WebhookV2Request) GetRequestInfoOk() (*WebhookV2RequestInfo, bool)`

GetRequestInfoOk returns a tuple with the RequestInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestInfo

`func (o *WebhookV2Request) SetRequestInfo(v WebhookV2RequestInfo)`

SetRequestInfo sets RequestInfo field to given value.


### GetResponseInfo

`func (o *WebhookV2Request) GetResponseInfo() WebhookV2ResponseInfo`

GetResponseInfo returns the ResponseInfo field if non-nil, zero value otherwise.

### GetResponseInfoOk

`func (o *WebhookV2Request) GetResponseInfoOk() (*WebhookV2ResponseInfo, bool)`

GetResponseInfoOk returns a tuple with the ResponseInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseInfo

`func (o *WebhookV2Request) SetResponseInfo(v WebhookV2ResponseInfo)`

SetResponseInfo sets ResponseInfo field to given value.


### GetErrorMsg

`func (o *WebhookV2Request) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *WebhookV2Request) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *WebhookV2Request) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.


### SetErrorMsgNil

`func (o *WebhookV2Request) SetErrorMsgNil(b bool)`

 SetErrorMsgNil sets the value for ErrorMsg to be an explicit nil

### UnsetErrorMsg
`func (o *WebhookV2Request) UnsetErrorMsg()`

UnsetErrorMsg ensures that no value is present for ErrorMsg, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


