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
)

// checks if the ColorStopBoundVariables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ColorStopBoundVariables{}

// ColorStopBoundVariables The variables bound to a particular gradient stop
type ColorStopBoundVariables struct {
	Color *VariableAlias `json:"color,omitempty"`
}

// NewColorStopBoundVariables instantiates a new ColorStopBoundVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColorStopBoundVariables() *ColorStopBoundVariables {
	this := ColorStopBoundVariables{}
	return &this
}

// NewColorStopBoundVariablesWithDefaults instantiates a new ColorStopBoundVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColorStopBoundVariablesWithDefaults() *ColorStopBoundVariables {
	this := ColorStopBoundVariables{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *ColorStopBoundVariables) GetColor() VariableAlias {
	if o == nil || IsNil(o.Color) {
		var ret VariableAlias
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ColorStopBoundVariables) GetColorOk() (*VariableAlias, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *ColorStopBoundVariables) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given VariableAlias and assigns it to the Color field.
func (o *ColorStopBoundVariables) SetColor(v VariableAlias) {
	o.Color = &v
}

func (o ColorStopBoundVariables) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ColorStopBoundVariables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	return toSerialize, nil
}

type NullableColorStopBoundVariables struct {
	value *ColorStopBoundVariables
	isSet bool
}

func (v NullableColorStopBoundVariables) Get() *ColorStopBoundVariables {
	return v.value
}

func (v *NullableColorStopBoundVariables) Set(val *ColorStopBoundVariables) {
	v.value = val
	v.isSet = true
}

func (v NullableColorStopBoundVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableColorStopBoundVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColorStopBoundVariables(val *ColorStopBoundVariables) *NullableColorStopBoundVariables {
	return &NullableColorStopBoundVariables{value: val, isSet: true}
}

func (v NullableColorStopBoundVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColorStopBoundVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


