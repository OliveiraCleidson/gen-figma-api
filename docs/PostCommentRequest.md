# PostCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The text contents of the comment to post. | 
**CommentId** | Pointer to **string** | The ID of the comment to reply to, if any. This must be a root comment. You cannot reply to other replies (a comment that has a parent_id). | [optional] 
**ClientMeta** | Pointer to [**PostCommentRequestClientMeta**](PostCommentRequestClientMeta.md) |  | [optional] 

## Methods

### NewPostCommentRequest

`func NewPostCommentRequest(message string, ) *PostCommentRequest`

NewPostCommentRequest instantiates a new PostCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCommentRequestWithDefaults

`func NewPostCommentRequestWithDefaults() *PostCommentRequest`

NewPostCommentRequestWithDefaults instantiates a new PostCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PostCommentRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostCommentRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostCommentRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCommentId

`func (o *PostCommentRequest) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *PostCommentRequest) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *PostCommentRequest) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *PostCommentRequest) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetClientMeta

`func (o *PostCommentRequest) GetClientMeta() PostCommentRequestClientMeta`

GetClientMeta returns the ClientMeta field if non-nil, zero value otherwise.

### GetClientMetaOk

`func (o *PostCommentRequest) GetClientMetaOk() (*PostCommentRequestClientMeta, bool)`

GetClientMetaOk returns a tuple with the ClientMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMeta

`func (o *PostCommentRequest) SetClientMeta(v PostCommentRequestClientMeta)`

SetClientMeta sets ClientMeta field to given value.

### HasClientMeta

`func (o *PostCommentRequest) HasClientMeta() bool`

HasClientMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


