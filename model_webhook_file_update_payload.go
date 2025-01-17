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

// checks if the WebhookFileUpdatePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookFileUpdatePayload{}

// WebhookFileUpdatePayload struct for WebhookFileUpdatePayload
type WebhookFileUpdatePayload struct {
	// The passcode specified when the webhook was created, should match what was initially provided
	Passcode string `json:"passcode"`
	// UTC ISO 8601 timestamp of when the event was triggered.
	Timestamp time.Time `json:"timestamp"`
	// The id of the webhook that caused the callback
	WebhookId string `json:"webhook_id"`
	EventType string `json:"event_type"`
	// The key of the file that was updated
	FileKey string `json:"file_key"`
	// The name of the file that was updated
	FileName string `json:"file_name"`
}

type _WebhookFileUpdatePayload WebhookFileUpdatePayload

// NewWebhookFileUpdatePayload instantiates a new WebhookFileUpdatePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookFileUpdatePayload(passcode string, timestamp time.Time, webhookId string, eventType string, fileKey string, fileName string) *WebhookFileUpdatePayload {
	this := WebhookFileUpdatePayload{}
	this.Passcode = passcode
	this.Timestamp = timestamp
	this.WebhookId = webhookId
	this.EventType = eventType
	this.FileKey = fileKey
	this.FileName = fileName
	return &this
}

// NewWebhookFileUpdatePayloadWithDefaults instantiates a new WebhookFileUpdatePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookFileUpdatePayloadWithDefaults() *WebhookFileUpdatePayload {
	this := WebhookFileUpdatePayload{}
	return &this
}

// GetPasscode returns the Passcode field value
func (o *WebhookFileUpdatePayload) GetPasscode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Passcode
}

// GetPasscodeOk returns a tuple with the Passcode field value
// and a boolean to check if the value has been set.
func (o *WebhookFileUpdatePayload) GetPasscodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Passcode, true
}

// SetPasscode sets field value
func (o *WebhookFileUpdatePayload) SetPasscode(v string) {
	o.Passcode = v
}

// GetTimestamp returns the Timestamp field value
func (o *WebhookFileUpdatePayload) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *WebhookFileUpdatePayload) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *WebhookFileUpdatePayload) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetWebhookId returns the WebhookId field value
func (o *WebhookFileUpdatePayload) GetWebhookId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value
// and a boolean to check if the value has been set.
func (o *WebhookFileUpdatePayload) GetWebhookIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebhookId, true
}

// SetWebhookId sets field value
func (o *WebhookFileUpdatePayload) SetWebhookId(v string) {
	o.WebhookId = v
}

// GetEventType returns the EventType field value
func (o *WebhookFileUpdatePayload) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *WebhookFileUpdatePayload) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *WebhookFileUpdatePayload) SetEventType(v string) {
	o.EventType = v
}

// GetFileKey returns the FileKey field value
func (o *WebhookFileUpdatePayload) GetFileKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileKey
}

// GetFileKeyOk returns a tuple with the FileKey field value
// and a boolean to check if the value has been set.
func (o *WebhookFileUpdatePayload) GetFileKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileKey, true
}

// SetFileKey sets field value
func (o *WebhookFileUpdatePayload) SetFileKey(v string) {
	o.FileKey = v
}

// GetFileName returns the FileName field value
func (o *WebhookFileUpdatePayload) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *WebhookFileUpdatePayload) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *WebhookFileUpdatePayload) SetFileName(v string) {
	o.FileName = v
}

func (o WebhookFileUpdatePayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookFileUpdatePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["passcode"] = o.Passcode
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["webhook_id"] = o.WebhookId
	toSerialize["event_type"] = o.EventType
	toSerialize["file_key"] = o.FileKey
	toSerialize["file_name"] = o.FileName
	return toSerialize, nil
}

func (o *WebhookFileUpdatePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"passcode",
		"timestamp",
		"webhook_id",
		"event_type",
		"file_key",
		"file_name",
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

	varWebhookFileUpdatePayload := _WebhookFileUpdatePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookFileUpdatePayload)

	if err != nil {
		return err
	}

	*o = WebhookFileUpdatePayload(varWebhookFileUpdatePayload)

	return err
}

type NullableWebhookFileUpdatePayload struct {
	value *WebhookFileUpdatePayload
	isSet bool
}

func (v NullableWebhookFileUpdatePayload) Get() *WebhookFileUpdatePayload {
	return v.value
}

func (v *NullableWebhookFileUpdatePayload) Set(val *WebhookFileUpdatePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookFileUpdatePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookFileUpdatePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookFileUpdatePayload(val *WebhookFileUpdatePayload) *NullableWebhookFileUpdatePayload {
	return &NullableWebhookFileUpdatePayload{value: val, isSet: true}
}

func (v NullableWebhookFileUpdatePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookFileUpdatePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


