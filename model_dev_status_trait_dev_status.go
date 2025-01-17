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

// checks if the DevStatusTraitDevStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DevStatusTraitDevStatus{}

// DevStatusTraitDevStatus Represents whether or not a node has a particular handoff (or dev) status applied to it.
type DevStatusTraitDevStatus struct {
	Type string `json:"type"`
	// An optional field where the designer can add more information about the design and what has changed.
	Description *string `json:"description,omitempty"`
}

type _DevStatusTraitDevStatus DevStatusTraitDevStatus

// NewDevStatusTraitDevStatus instantiates a new DevStatusTraitDevStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDevStatusTraitDevStatus(type_ string) *DevStatusTraitDevStatus {
	this := DevStatusTraitDevStatus{}
	this.Type = type_
	return &this
}

// NewDevStatusTraitDevStatusWithDefaults instantiates a new DevStatusTraitDevStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDevStatusTraitDevStatusWithDefaults() *DevStatusTraitDevStatus {
	this := DevStatusTraitDevStatus{}
	return &this
}

// GetType returns the Type field value
func (o *DevStatusTraitDevStatus) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DevStatusTraitDevStatus) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DevStatusTraitDevStatus) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DevStatusTraitDevStatus) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DevStatusTraitDevStatus) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DevStatusTraitDevStatus) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DevStatusTraitDevStatus) SetDescription(v string) {
	o.Description = &v
}

func (o DevStatusTraitDevStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DevStatusTraitDevStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

func (o *DevStatusTraitDevStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varDevStatusTraitDevStatus := _DevStatusTraitDevStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDevStatusTraitDevStatus)

	if err != nil {
		return err
	}

	*o = DevStatusTraitDevStatus(varDevStatusTraitDevStatus)

	return err
}

type NullableDevStatusTraitDevStatus struct {
	value *DevStatusTraitDevStatus
	isSet bool
}

func (v NullableDevStatusTraitDevStatus) Get() *DevStatusTraitDevStatus {
	return v.value
}

func (v *NullableDevStatusTraitDevStatus) Set(val *DevStatusTraitDevStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDevStatusTraitDevStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDevStatusTraitDevStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDevStatusTraitDevStatus(val *DevStatusTraitDevStatus) *NullableDevStatusTraitDevStatus {
	return &NullableDevStatusTraitDevStatus{value: val, isSet: true}
}

func (v NullableDevStatusTraitDevStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDevStatusTraitDevStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


