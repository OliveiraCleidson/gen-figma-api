/*
Figma API

This is the OpenAPI specification for the [Figma REST API](https://www.figma.com/developers/api).  Note: we are releasing the OpenAPI specification as a beta given the large surface area and complexity of the REST API. If you notice any inaccuracies with the specification, please [file an issue](https://github.com/figma/rest-api-spec/issues).

API version: 0.20.0
Contact: support@figma.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the WebhookFileCommentPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookFileCommentPayload{}

// WebhookFileCommentPayload struct for WebhookFileCommentPayload
type WebhookFileCommentPayload struct {
	// The passcode specified when the webhook was created, should match what was initially provided
	Passcode string `json:"passcode"`
	// UTC ISO 8601 timestamp of when the event was triggered.
	Timestamp time.Time `json:"timestamp"`
	// The id of the webhook that caused the callback
	WebhookId string `json:"webhook_id"`
	EventType string `json:"event_type"`
	// Contents of the comment itself
	Comment []CommentFragment `json:"comment"`
	// Unique identifier for comment
	CommentId string `json:"comment_id"`
	// The UTC ISO 8601 time at which the comment was left
	CreatedAt time.Time `json:"created_at"`
	// The key of the file that was commented on
	FileKey string `json:"file_key"`
	// The name of the file that was commented on
	FileName string `json:"file_name"`
	// Users that were mentioned in the comment
	Mentions []User `json:"mentions,omitempty"`
	// The user that made the comment and triggered this event
	TriggeredBy User `json:"triggered_by"`
}

type _WebhookFileCommentPayload WebhookFileCommentPayload

// NewWebhookFileCommentPayload instantiates a new WebhookFileCommentPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookFileCommentPayload(passcode string, timestamp time.Time, webhookId string, eventType string, comment []CommentFragment, commentId string, createdAt time.Time, fileKey string, fileName string, triggeredBy User) *WebhookFileCommentPayload {
	this := WebhookFileCommentPayload{}
	this.Passcode = passcode
	this.Timestamp = timestamp
	this.WebhookId = webhookId
	this.EventType = eventType
	this.Comment = comment
	this.CommentId = commentId
	this.CreatedAt = createdAt
	this.FileKey = fileKey
	this.FileName = fileName
	this.TriggeredBy = triggeredBy
	return &this
}

// NewWebhookFileCommentPayloadWithDefaults instantiates a new WebhookFileCommentPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookFileCommentPayloadWithDefaults() *WebhookFileCommentPayload {
	this := WebhookFileCommentPayload{}
	return &this
}

// GetPasscode returns the Passcode field value
func (o *WebhookFileCommentPayload) GetPasscode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Passcode
}

// GetPasscodeOk returns a tuple with the Passcode field value
// and a boolean to check if the value has been set.
func (o *WebhookFileCommentPayload) GetPasscodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Passcode, true
}

// SetPasscode sets field value
func (o *WebhookFileCommentPayload) SetPasscode(v string) {
	o.Passcode = v
}

// GetTimestamp returns the Timestamp field value
func (o *WebhookFileCommentPayload) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *WebhookFileCommentPayload) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *WebhookFileCommentPayload) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetWebhookId returns the WebhookId field value
func (o *WebhookFileCommentPayload) GetWebhookId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value
// and a boolean to check if the value has been set.
func (o *WebhookFileCommentPayload) GetWebhookIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhookId, true
}

// SetWebhookId sets field value
func (o *WebhookFileCommentPayload) SetWebhookId(v string) {
	o.WebhookId = v
}

// GetEventType returns the EventType field value
func (o *WebhookFileCommentPayload) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *WebhookFileCommentPayload) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *WebhookFileCommentPayload) SetEventType(v string) {
	o.EventType = v
}

// GetComment returns the Comment field value
func (o *WebhookFileCommentPayload) GetComment() []CommentFragment {
	if o == nil {
		var ret []CommentFragment
		return ret
	}

	return o.Comment
}

// GetCommentOk returns a tuple with the Comment field value
// and a boolean to check if the value has been set.
func (o *WebhookFileCommentPayload) GetCommentOk() ([]CommentFragment, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comment, true
}

// SetComment sets field value
func (o *WebhookFileCommentPayload) SetComment(v []CommentFragment) {
	o.Comment = v
}

// GetCommentId returns the CommentId field value
func (o *WebhookFileCommentPayload) GetCommentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CommentId
}

// GetCommentIdOk returns a tuple with the CommentId field value
// and a boolean to check if the value has been set.
func (o *WebhookFileCommentPayload) GetCommentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommentId, true
}

// SetCommentId sets field value
func (o *WebhookFileCommentPayload) SetCommentId(v string) {
	o.CommentId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *WebhookFileCommentPayload) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *WebhookFileCommentPayload) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *WebhookFileCommentPayload) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetFileKey returns the FileKey field value
func (o *WebhookFileCommentPayload) GetFileKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileKey
}

// GetFileKeyOk returns a tuple with the FileKey field value
// and a boolean to check if the value has been set.
func (o *WebhookFileCommentPayload) GetFileKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileKey, true
}

// SetFileKey sets field value
func (o *WebhookFileCommentPayload) SetFileKey(v string) {
	o.FileKey = v
}

// GetFileName returns the FileName field value
func (o *WebhookFileCommentPayload) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *WebhookFileCommentPayload) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *WebhookFileCommentPayload) SetFileName(v string) {
	o.FileName = v
}

// GetMentions returns the Mentions field value if set, zero value otherwise.
func (o *WebhookFileCommentPayload) GetMentions() []User {
	if o == nil || IsNil(o.Mentions) {
		var ret []User
		return ret
	}
	return o.Mentions
}

// GetMentionsOk returns a tuple with the Mentions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookFileCommentPayload) GetMentionsOk() ([]User, bool) {
	if o == nil || IsNil(o.Mentions) {
		return nil, false
	}
	return o.Mentions, true
}

// HasMentions returns a boolean if a field has been set.
func (o *WebhookFileCommentPayload) HasMentions() bool {
	if o != nil && !IsNil(o.Mentions) {
		return true
	}

	return false
}

// SetMentions gets a reference to the given []User and assigns it to the Mentions field.
func (o *WebhookFileCommentPayload) SetMentions(v []User) {
	o.Mentions = v
}

// GetTriggeredBy returns the TriggeredBy field value
func (o *WebhookFileCommentPayload) GetTriggeredBy() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.TriggeredBy
}

// GetTriggeredByOk returns a tuple with the TriggeredBy field value
// and a boolean to check if the value has been set.
func (o *WebhookFileCommentPayload) GetTriggeredByOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggeredBy, true
}

// SetTriggeredBy sets field value
func (o *WebhookFileCommentPayload) SetTriggeredBy(v User) {
	o.TriggeredBy = v
}

func (o WebhookFileCommentPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookFileCommentPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["passcode"] = o.Passcode
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["webhook_id"] = o.WebhookId
	toSerialize["event_type"] = o.EventType
	toSerialize["comment"] = o.Comment
	toSerialize["comment_id"] = o.CommentId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["file_key"] = o.FileKey
	toSerialize["file_name"] = o.FileName
	if !IsNil(o.Mentions) {
		toSerialize["mentions"] = o.Mentions
	}
	toSerialize["triggered_by"] = o.TriggeredBy
	return toSerialize, nil
}

func (o *WebhookFileCommentPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"passcode",
		"timestamp",
		"webhook_id",
		"event_type",
		"comment",
		"comment_id",
		"created_at",
		"file_key",
		"file_name",
		"triggered_by",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWebhookFileCommentPayload := _WebhookFileCommentPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookFileCommentPayload)

	if err != nil {
		return err
	}

	*o = WebhookFileCommentPayload(varWebhookFileCommentPayload)

	return err
}

type NullableWebhookFileCommentPayload struct {
	value *WebhookFileCommentPayload
	isSet bool
}

func (v NullableWebhookFileCommentPayload) Get() *WebhookFileCommentPayload {
	return v.value
}

func (v *NullableWebhookFileCommentPayload) Set(val *WebhookFileCommentPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookFileCommentPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookFileCommentPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookFileCommentPayload(val *WebhookFileCommentPayload) *NullableWebhookFileCommentPayload {
	return &NullableWebhookFileCommentPayload{value: val, isSet: true}
}

func (v NullableWebhookFileCommentPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookFileCommentPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


