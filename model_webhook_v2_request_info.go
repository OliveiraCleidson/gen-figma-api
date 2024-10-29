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

// checks if the WebhookV2RequestInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookV2RequestInfo{}

// WebhookV2RequestInfo Information regarding the request sent to a webhook endpoint
type WebhookV2RequestInfo struct {
	// The ID of the webhook
	Id string `json:"id"`
	// The actual endpoint the request was sent to
	Endpoint string `json:"endpoint"`
	// The contents of the request that was sent to the endpoint
	Payload map[string]interface{} `json:"payload"`
	// UTC ISO 8601 timestamp of when the request was sent
	SentAt time.Time `json:"sent_at"`
}

type _WebhookV2RequestInfo WebhookV2RequestInfo

// NewWebhookV2RequestInfo instantiates a new WebhookV2RequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookV2RequestInfo(id string, endpoint string, payload map[string]interface{}, sentAt time.Time) *WebhookV2RequestInfo {
	this := WebhookV2RequestInfo{}
	this.Id = id
	this.Endpoint = endpoint
	this.Payload = payload
	this.SentAt = sentAt
	return &this
}

// NewWebhookV2RequestInfoWithDefaults instantiates a new WebhookV2RequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookV2RequestInfoWithDefaults() *WebhookV2RequestInfo {
	this := WebhookV2RequestInfo{}
	return &this
}

// GetId returns the Id field value
func (o *WebhookV2RequestInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WebhookV2RequestInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WebhookV2RequestInfo) SetId(v string) {
	o.Id = v
}

// GetEndpoint returns the Endpoint field value
func (o *WebhookV2RequestInfo) GetEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value
// and a boolean to check if the value has been set.
func (o *WebhookV2RequestInfo) GetEndpointOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Endpoint, true
}

// SetEndpoint sets field value
func (o *WebhookV2RequestInfo) SetEndpoint(v string) {
	o.Endpoint = v
}

// GetPayload returns the Payload field value
func (o *WebhookV2RequestInfo) GetPayload() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *WebhookV2RequestInfo) GetPayloadOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Payload, true
}

// SetPayload sets field value
func (o *WebhookV2RequestInfo) SetPayload(v map[string]interface{}) {
	o.Payload = v
}

// GetSentAt returns the SentAt field value
func (o *WebhookV2RequestInfo) GetSentAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.SentAt
}

// GetSentAtOk returns a tuple with the SentAt field value
// and a boolean to check if the value has been set.
func (o *WebhookV2RequestInfo) GetSentAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SentAt, true
}

// SetSentAt sets field value
func (o *WebhookV2RequestInfo) SetSentAt(v time.Time) {
	o.SentAt = v
}

func (o WebhookV2RequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookV2RequestInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["endpoint"] = o.Endpoint
	toSerialize["payload"] = o.Payload
	toSerialize["sent_at"] = o.SentAt
	return toSerialize, nil
}

func (o *WebhookV2RequestInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"endpoint",
		"payload",
		"sent_at",
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

	varWebhookV2RequestInfo := _WebhookV2RequestInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookV2RequestInfo)

	if err != nil {
		return err
	}

	*o = WebhookV2RequestInfo(varWebhookV2RequestInfo)

	return err
}

type NullableWebhookV2RequestInfo struct {
	value *WebhookV2RequestInfo
	isSet bool
}

func (v NullableWebhookV2RequestInfo) Get() *WebhookV2RequestInfo {
	return v.value
}

func (v *NullableWebhookV2RequestInfo) Set(val *WebhookV2RequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookV2RequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookV2RequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookV2RequestInfo(val *WebhookV2RequestInfo) *NullableWebhookV2RequestInfo {
	return &NullableWebhookV2RequestInfo{value: val, isSet: true}
}

func (v NullableWebhookV2RequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookV2RequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


