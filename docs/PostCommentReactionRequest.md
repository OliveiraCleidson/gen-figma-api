# PostCommentReactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emoji** | **string** | The emoji type of reaction as shortcode (e.g. &#x60;:heart:&#x60;, &#x60;:+1::skin-tone-2:&#x60;). The list of accepted emoji shortcodes can be found in [this file](https://raw.githubusercontent.com/missive/emoji-mart/main/packages/emoji-mart-data/sets/14/native.json) under the top-level emojis and aliases fields, with optional skin tone modifiers when applicable. | 

## Methods

### NewPostCommentReactionRequest

`func NewPostCommentReactionRequest(emoji string, ) *PostCommentReactionRequest`

NewPostCommentReactionRequest instantiates a new PostCommentReactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCommentReactionRequestWithDefaults

`func NewPostCommentReactionRequestWithDefaults() *PostCommentReactionRequest`

NewPostCommentReactionRequestWithDefaults instantiates a new PostCommentReactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmoji

`func (o *PostCommentReactionRequest) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *PostCommentReactionRequest) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *PostCommentReactionRequest) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


