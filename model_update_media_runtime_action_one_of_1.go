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

// checks if the UpdateMediaRuntimeActionOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMediaRuntimeActionOneOf1{}

// UpdateMediaRuntimeActionOneOf1 An action that updates the runtime of a media node by skipping forward or backward.  The `destinationId` is the node ID of the media node to update. If `destinationId` is `null`, the action will  update the media node that contains the action.  The `mediaAction` is the action to perform on the media node.  The `amountToSkip` is the amount of time to skip in seconds.
type UpdateMediaRuntimeActionOneOf1 struct {
	Type string `json:"type"`
	DestinationId NullableString `json:"destinationId,omitempty"`
	MediaAction string `json:"mediaAction"`
	AmountToSkip float32 `json:"amountToSkip"`
}

type _UpdateMediaRuntimeActionOneOf1 UpdateMediaRuntimeActionOneOf1

// NewUpdateMediaRuntimeActionOneOf1 instantiates a new UpdateMediaRuntimeActionOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMediaRuntimeActionOneOf1(type_ string, mediaAction string, amountToSkip float32) *UpdateMediaRuntimeActionOneOf1 {
	this := UpdateMediaRuntimeActionOneOf1{}
	this.Type = type_
	this.MediaAction = mediaAction
	this.AmountToSkip = amountToSkip
	return &this
}

// NewUpdateMediaRuntimeActionOneOf1WithDefaults instantiates a new UpdateMediaRuntimeActionOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMediaRuntimeActionOneOf1WithDefaults() *UpdateMediaRuntimeActionOneOf1 {
	this := UpdateMediaRuntimeActionOneOf1{}
	return &this
}

// GetType returns the Type field value
func (o *UpdateMediaRuntimeActionOneOf1) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UpdateMediaRuntimeActionOneOf1) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UpdateMediaRuntimeActionOneOf1) SetType(v string) {
	o.Type = v
}

// GetDestinationId returns the DestinationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMediaRuntimeActionOneOf1) GetDestinationId() string {
	if o == nil || IsNil(o.DestinationId.Get()) {
		var ret string
		return ret
	}
	return *o.DestinationId.Get()
}

// GetDestinationIdOk returns a tuple with the DestinationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMediaRuntimeActionOneOf1) GetDestinationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DestinationId.Get(), o.DestinationId.IsSet()
}

// HasDestinationId returns a boolean if a field has been set.
func (o *UpdateMediaRuntimeActionOneOf1) HasDestinationId() bool {
	if o != nil && o.DestinationId.IsSet() {
		return true
	}

	return false
}

// SetDestinationId gets a reference to the given NullableString and assigns it to the DestinationId field.
func (o *UpdateMediaRuntimeActionOneOf1) SetDestinationId(v string) {
	o.DestinationId.Set(&v)
}
// SetDestinationIdNil sets the value for DestinationId to be an explicit nil
func (o *UpdateMediaRuntimeActionOneOf1) SetDestinationIdNil() {
	o.DestinationId.Set(nil)
}

// UnsetDestinationId ensures that no value is present for DestinationId, not even an explicit nil
func (o *UpdateMediaRuntimeActionOneOf1) UnsetDestinationId() {
	o.DestinationId.Unset()
}

// GetMediaAction returns the MediaAction field value
func (o *UpdateMediaRuntimeActionOneOf1) GetMediaAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MediaAction
}

// GetMediaActionOk returns a tuple with the MediaAction field value
// and a boolean to check if the value has been set.
func (o *UpdateMediaRuntimeActionOneOf1) GetMediaActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaAction, true
}

// SetMediaAction sets field value
func (o *UpdateMediaRuntimeActionOneOf1) SetMediaAction(v string) {
	o.MediaAction = v
}

// GetAmountToSkip returns the AmountToSkip field value
func (o *UpdateMediaRuntimeActionOneOf1) GetAmountToSkip() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AmountToSkip
}

// GetAmountToSkipOk returns a tuple with the AmountToSkip field value
// and a boolean to check if the value has been set.
func (o *UpdateMediaRuntimeActionOneOf1) GetAmountToSkipOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountToSkip, true
}

// SetAmountToSkip sets field value
func (o *UpdateMediaRuntimeActionOneOf1) SetAmountToSkip(v float32) {
	o.AmountToSkip = v
}

func (o UpdateMediaRuntimeActionOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMediaRuntimeActionOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.DestinationId.IsSet() {
		toSerialize["destinationId"] = o.DestinationId.Get()
	}
	toSerialize["mediaAction"] = o.MediaAction
	toSerialize["amountToSkip"] = o.AmountToSkip
	return toSerialize, nil
}

func (o *UpdateMediaRuntimeActionOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"mediaAction",
		"amountToSkip",
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

	varUpdateMediaRuntimeActionOneOf1 := _UpdateMediaRuntimeActionOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateMediaRuntimeActionOneOf1)

	if err != nil {
		return err
	}

	*o = UpdateMediaRuntimeActionOneOf1(varUpdateMediaRuntimeActionOneOf1)

	return err
}

type NullableUpdateMediaRuntimeActionOneOf1 struct {
	value *UpdateMediaRuntimeActionOneOf1
	isSet bool
}

func (v NullableUpdateMediaRuntimeActionOneOf1) Get() *UpdateMediaRuntimeActionOneOf1 {
	return v.value
}

func (v *NullableUpdateMediaRuntimeActionOneOf1) Set(val *UpdateMediaRuntimeActionOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMediaRuntimeActionOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMediaRuntimeActionOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMediaRuntimeActionOneOf1(val *UpdateMediaRuntimeActionOneOf1) *NullableUpdateMediaRuntimeActionOneOf1 {
	return &NullableUpdateMediaRuntimeActionOneOf1{value: val, isSet: true}
}

func (v NullableUpdateMediaRuntimeActionOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMediaRuntimeActionOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


