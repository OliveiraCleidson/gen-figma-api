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
	"bytes"
	"fmt"
)

// checks if the UpdateMediaRuntimeActionOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMediaRuntimeActionOneOf2{}

// UpdateMediaRuntimeActionOneOf2 An action that updates the runtime of a media node by skipping to a specific time.  The `destinationId` is the node ID of the media node to update. If `destinationId` is `null`, the action will  update the media node that contains the action.  The `mediaAction` is the action to perform on the media node.  The `newTimestamp` is the new time to skip to in seconds.
type UpdateMediaRuntimeActionOneOf2 struct {
	Type string `json:"type"`
	DestinationId NullableString `json:"destinationId,omitempty"`
	MediaAction string `json:"mediaAction"`
	NewTimestamp float32 `json:"newTimestamp"`
}

type _UpdateMediaRuntimeActionOneOf2 UpdateMediaRuntimeActionOneOf2

// NewUpdateMediaRuntimeActionOneOf2 instantiates a new UpdateMediaRuntimeActionOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMediaRuntimeActionOneOf2(type_ string, mediaAction string, newTimestamp float32) *UpdateMediaRuntimeActionOneOf2 {
	this := UpdateMediaRuntimeActionOneOf2{}
	this.Type = type_
	this.MediaAction = mediaAction
	this.NewTimestamp = newTimestamp
	return &this
}

// NewUpdateMediaRuntimeActionOneOf2WithDefaults instantiates a new UpdateMediaRuntimeActionOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMediaRuntimeActionOneOf2WithDefaults() *UpdateMediaRuntimeActionOneOf2 {
	this := UpdateMediaRuntimeActionOneOf2{}
	return &this
}

// GetType returns the Type field value
func (o *UpdateMediaRuntimeActionOneOf2) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateMediaRuntimeActionOneOf2) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateMediaRuntimeActionOneOf2) SetType(v string) {
	o.Type = v
}

// GetDestinationId returns the DestinationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMediaRuntimeActionOneOf2) GetDestinationId() string {
	if o == nil || IsNil(o.DestinationId.Get()) {
		var ret string
		return ret
	}
	return *o.DestinationId.Get()
}

// GetDestinationIdOk returns a tuple with the DestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMediaRuntimeActionOneOf2) GetDestinationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestinationId.Get(), o.DestinationId.IsSet()
}

// HasDestinationId returns a boolean if a field has been set.
func (o *UpdateMediaRuntimeActionOneOf2) HasDestinationId() bool {
	if o != nil && o.DestinationId.IsSet() {
		return true
	}

	return false
}

// SetDestinationId gets a reference to the given NullableString and assigns it to the DestinationId field.
func (o *UpdateMediaRuntimeActionOneOf2) SetDestinationId(v string) {
	o.DestinationId.Set(&v)
}
// SetDestinationIdNil sets the value for DestinationId to be an explicit nil
func (o *UpdateMediaRuntimeActionOneOf2) SetDestinationIdNil() {
	o.DestinationId.Set(nil)
}

// UnsetDestinationId ensures that no value is present for DestinationId, not even an explicit nil
func (o *UpdateMediaRuntimeActionOneOf2) UnsetDestinationId() {
	o.DestinationId.Unset()
}

// GetMediaAction returns the MediaAction field value
func (o *UpdateMediaRuntimeActionOneOf2) GetMediaAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaAction
}

// GetMediaActionOk returns a tuple with the MediaAction field value
// and a boolean to check if the value has been set.
func (o *UpdateMediaRuntimeActionOneOf2) GetMediaActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaAction, true
}

// SetMediaAction sets field value
func (o *UpdateMediaRuntimeActionOneOf2) SetMediaAction(v string) {
	o.MediaAction = v
}

// GetNewTimestamp returns the NewTimestamp field value
func (o *UpdateMediaRuntimeActionOneOf2) GetNewTimestamp() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NewTimestamp
}

// GetNewTimestampOk returns a tuple with the NewTimestamp field value
// and a boolean to check if the value has been set.
func (o *UpdateMediaRuntimeActionOneOf2) GetNewTimestampOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewTimestamp, true
}

// SetNewTimestamp sets field value
func (o *UpdateMediaRuntimeActionOneOf2) SetNewTimestamp(v float32) {
	o.NewTimestamp = v
}

func (o UpdateMediaRuntimeActionOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMediaRuntimeActionOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.DestinationId.IsSet() {
		toSerialize["destinationId"] = o.DestinationId.Get()
	}
	toSerialize["mediaAction"] = o.MediaAction
	toSerialize["newTimestamp"] = o.NewTimestamp
	return toSerialize, nil
}

func (o *UpdateMediaRuntimeActionOneOf2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"mediaAction",
		"newTimestamp",
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

	varUpdateMediaRuntimeActionOneOf2 := _UpdateMediaRuntimeActionOneOf2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateMediaRuntimeActionOneOf2)

	if err != nil {
		return err
	}

	*o = UpdateMediaRuntimeActionOneOf2(varUpdateMediaRuntimeActionOneOf2)

	return err
}

type NullableUpdateMediaRuntimeActionOneOf2 struct {
	value *UpdateMediaRuntimeActionOneOf2
	isSet bool
}

func (v NullableUpdateMediaRuntimeActionOneOf2) Get() *UpdateMediaRuntimeActionOneOf2 {
	return v.value
}

func (v *NullableUpdateMediaRuntimeActionOneOf2) Set(val *UpdateMediaRuntimeActionOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMediaRuntimeActionOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMediaRuntimeActionOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMediaRuntimeActionOneOf2(val *UpdateMediaRuntimeActionOneOf2) *NullableUpdateMediaRuntimeActionOneOf2 {
	return &NullableUpdateMediaRuntimeActionOneOf2{value: val, isSet: true}
}

func (v NullableUpdateMediaRuntimeActionOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMediaRuntimeActionOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


