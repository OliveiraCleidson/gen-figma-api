# CommentFragment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Comment text that is set if a fragment is text based | [optional] 
**Mention** | Pointer to **string** | User id that is set if a fragment refers to a user mention | [optional] 

## Methods

### NewCommentFragment

`func NewCommentFragment() *CommentFragment`

NewCommentFragment instantiates a new CommentFragment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentFragmentWithDefaults

`func NewCommentFragmentWithDefaults() *CommentFragment`

NewCommentFragmentWithDefaults instantiates a new CommentFragment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *CommentFragment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CommentFragment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CommentFragment) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CommentFragment) HasText() bool`

HasText returns a boolean if a field has been set.

### GetMention

`func (o *CommentFragment) GetMention() string`

GetMention returns the Mention field if non-nil, zero value otherwise.

### GetMentionOk

`func (o *CommentFragment) GetMentionOk() (*string, bool)`

GetMentionOk returns a tuple with the Mention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMention

`func (o *CommentFragment) SetMention(v string)`

SetMention sets Mention field to given value.

### HasMention

`func (o *CommentFragment) HasMention() bool`

HasMention returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


