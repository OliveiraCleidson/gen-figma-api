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

// checks if the RGBA type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RGBA{}

// RGBA An RGBA color
type RGBA struct {
	// Red channel value, between 0 and 1.
	R float32 `json:"r"`
	// Green channel value, between 0 and 1.
	G float32 `json:"g"`
	// Blue channel value, between 0 and 1.
	B float32 `json:"b"`
	// Alpha channel value, between 0 and 1.
	A float32 `json:"a"`
}

type _RGBA RGBA

// NewRGBA instantiates a new RGBA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRGBA(r float32, g float32, b float32, a float32) *RGBA {
	this := RGBA{}
	this.R = r
	this.G = g
	this.B = b
	this.A = a
	return &this
}

// NewRGBAWithDefaults instantiates a new RGBA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRGBAWithDefaults() *RGBA {
	this := RGBA{}
	return &this
}

// GetR returns the R field value
func (o *RGBA) GetR() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.R
}

// GetROk returns a tuple with the R field value
// and a boolean to check if the value has been set.
func (o *RGBA) GetROk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.R, true
}

// SetR sets field value
func (o *RGBA) SetR(v float32) {
	o.R = v
}

// GetG returns the G field value
func (o *RGBA) GetG() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.G
}

// GetGOk returns a tuple with the G field value
// and a boolean to check if the value has been set.
func (o *RGBA) GetGOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.G, true
}

// SetG sets field value
func (o *RGBA) SetG(v float32) {
	o.G = v
}

// GetB returns the B field value
func (o *RGBA) GetB() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.B
}

// GetBOk returns a tuple with the B field value
// and a boolean to check if the value has been set.
func (o *RGBA) GetBOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.B, true
}

// SetB sets field value
func (o *RGBA) SetB(v float32) {
	o.B = v
}

// GetA returns the A field value
func (o *RGBA) GetA() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.A
}

// GetAOk returns a tuple with the A field value
// and a boolean to check if the value has been set.
func (o *RGBA) GetAOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.A, true
}

// SetA sets field value
func (o *RGBA) SetA(v float32) {
	o.A = v
}

func (o RGBA) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RGBA) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["r"] = o.R
	toSerialize["g"] = o.G
	toSerialize["b"] = o.B
	toSerialize["a"] = o.A
	return toSerialize, nil
}

func (o *RGBA) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"r",
		"g",
		"b",
		"a",
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

	varRGBA := _RGBA{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRGBA)

	if err != nil {
		return err
	}

	*o = RGBA(varRGBA)

	return err
}

type NullableRGBA struct {
	value *RGBA
	isSet bool
}

func (v NullableRGBA) Get() *RGBA {
	return v.value
}

func (v *NullableRGBA) Set(val *RGBA) {
	v.value = val
	v.isSet = true
}

func (v NullableRGBA) IsSet() bool {
	return v.isSet
}

func (v *NullableRGBA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRGBA(val *RGBA) *NullableRGBA {
	return &NullableRGBA{value: val, isSet: true}
}

func (v NullableRGBA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRGBA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


