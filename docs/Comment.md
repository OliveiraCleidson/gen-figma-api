# Comment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for comment. | 
**ClientMeta** | [**CommentClientMeta**](CommentClientMeta.md) |  | 
**FileKey** | **string** | The file in which the comment lives | 
**ParentId** | Pointer to **string** | If present, the id of the comment to which this is the reply | [optional] 
**User** | [**User**](User.md) | The user who left the comment | 
**CreatedAt** | **time.Time** | The UTC ISO 8601 time at which the comment was left | 
**ResolvedAt** | Pointer to **NullableTime** | If set, the UTC ISO 8601 time the comment was resolved | [optional] 
**Message** | **string** | The content of the comment | 
**OrderId** | **NullableString** | Only set for top level comments. The number displayed with the comment in the UI | 
**Reactions** | [**[]Reaction**](Reaction.md) | An array of reactions to the comment | 

## Methods

### NewComment

`func NewComment(id string, clientMeta CommentClientMeta, fileKey string, user User, createdAt time.Time, message string, orderId NullableString, reactions []Reaction, ) *Comment`

NewComment instantiates a new Comment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentWithDefaults

`func NewCommentWithDefaults() *Comment`

NewCommentWithDefaults instantiates a new Comment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Comment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Comment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Comment) SetId(v string)`

SetId sets Id field to given value.


### GetClientMeta

`func (o *Comment) GetClientMeta() CommentClientMeta`

GetClientMeta returns the ClientMeta field if non-nil, zero value otherwise.

### GetClientMetaOk

`func (o *Comment) GetClientMetaOk() (*CommentClientMeta, bool)`

GetClientMetaOk returns a tuple with the ClientMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMeta

`func (o *Comment) SetClientMeta(v CommentClientMeta)`

SetClientMeta sets ClientMeta field to given value.


### GetFileKey

`func (o *Comment) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *Comment) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *Comment) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.


### GetParentId

`func (o *Comment) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Comment) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Comment) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Comment) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetUser

`func (o *Comment) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Comment) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Comment) SetUser(v User)`

SetUser sets User field to given value.


### GetCreatedAt

`func (o *Comment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Comment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Comment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetResolvedAt

`func (o *Comment) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *Comment) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *Comment) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *Comment) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### SetResolvedAtNil

`func (o *Comment) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *Comment) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetMessage

`func (o *Comment) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Comment) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Comment) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetOrderId

`func (o *Comment) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Comment) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Comment) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### SetOrderIdNil

`func (o *Comment) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *Comment) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetReactions

`func (o *Comment) GetReactions() []Reaction`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *Comment) GetReactionsOk() (*[]Reaction, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *Comment) SetReactions(v []Reaction)`

SetReactions sets Reactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


