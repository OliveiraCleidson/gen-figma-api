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

// checks if the OnMediaHitTrigger type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OnMediaHitTrigger{}

// OnMediaHitTrigger struct for OnMediaHitTrigger
type OnMediaHitTrigger struct {
	Type string `json:"type"`
	MediaHitTime float32 `json:"mediaHitTime"`
}

type _OnMediaHitTrigger OnMediaHitTrigger

// NewOnMediaHitTrigger instantiates a new OnMediaHitTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOnMediaHitTrigger(type_ string, mediaHitTime float32) *OnMediaHitTrigger {
	this := OnMediaHitTrigger{}
	this.Type = type_
	this.MediaHitTime = mediaHitTime
	return &this
}

// NewOnMediaHitTriggerWithDefaults instantiates a new OnMediaHitTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOnMediaHitTriggerWithDefaults() *OnMediaHitTrigger {
	this := OnMediaHitTrigger{}
	return &this
}

// GetType returns the Type field value
func (o *OnMediaHitTrigger) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *OnMediaHitTrigger) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *OnMediaHitTrigger) SetType(v string) {
	o.Type = v
}

// GetMediaHitTime returns the MediaHitTime field value
func (o *OnMediaHitTrigger) GetMediaHitTime() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.MediaHitTime
}

// GetMediaHitTimeOk returns a tuple with the MediaHitTime field value
// and a boolean to check if the value has been set.
func (o *OnMediaHitTrigger) GetMediaHitTimeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MediaHitTime, true
}

// SetMediaHitTime sets field value
func (o *OnMediaHitTrigger) SetMediaHitTime(v float32) {
	o.MediaHitTime = v
}

func (o OnMediaHitTrigger) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OnMediaHitTrigger) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["mediaHitTime"] = o.MediaHitTime
	return toSerialize, nil
}

func (o *OnMediaHitTrigger) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"mediaHitTime",
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

	varOnMediaHitTrigger := _OnMediaHitTrigger{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOnMediaHitTrigger)

	if err != nil {
		return err
	}

	*o = OnMediaHitTrigger(varOnMediaHitTrigger)

	return err
}

type NullableOnMediaHitTrigger struct {
	value *OnMediaHitTrigger
	isSet bool
}

func (v NullableOnMediaHitTrigger) Get() *OnMediaHitTrigger {
	return v.value
}

func (v *NullableOnMediaHitTrigger) Set(val *OnMediaHitTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableOnMediaHitTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableOnMediaHitTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOnMediaHitTrigger(val *OnMediaHitTrigger) *NullableOnMediaHitTrigger {
	return &NullableOnMediaHitTrigger{value: val, isSet: true}
}

func (v NullableOnMediaHitTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOnMediaHitTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

