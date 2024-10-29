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

// checks if the BlurEffectBoundVariables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BlurEffectBoundVariables{}

// BlurEffectBoundVariables The variables bound to a particular field on this blur effect
type BlurEffectBoundVariables struct {
	Radius *VariableAlias `json:"radius,omitempty"`
}

// NewBlurEffectBoundVariables instantiates a new BlurEffectBoundVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlurEffectBoundVariables() *BlurEffectBoundVariables {
	this := BlurEffectBoundVariables{}
	return &this
}

// NewBlurEffectBoundVariablesWithDefaults instantiates a new BlurEffectBoundVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlurEffectBoundVariablesWithDefaults() *BlurEffectBoundVariables {
	this := BlurEffectBoundVariables{}
	return &this
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *BlurEffectBoundVariables) GetRadius() VariableAlias {
	if o == nil || IsNil(o.Radius) {
		var ret VariableAlias
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlurEffectBoundVariables) GetRadiusOk() (*VariableAlias, bool) {
	if o == nil || IsNil(o.Radius) {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *BlurEffectBoundVariables) HasRadius() bool {
	if o != nil && !IsNil(o.Radius) {
		return true
	}

	return false
}

// SetRadius gets a reference to the given VariableAlias and assigns it to the Radius field.
func (o *BlurEffectBoundVariables) SetRadius(v VariableAlias) {
	o.Radius = &v
}

func (o BlurEffectBoundVariables) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BlurEffectBoundVariables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Radius) {
		toSerialize["radius"] = o.Radius
	}
	return toSerialize, nil
}

type NullableBlurEffectBoundVariables struct {
	value *BlurEffectBoundVariables
	isSet bool
}

func (v NullableBlurEffectBoundVariables) Get() *BlurEffectBoundVariables {
	return v.value
}

func (v *NullableBlurEffectBoundVariables) Set(val *BlurEffectBoundVariables) {
	v.value = val
	v.isSet = true
}

func (v NullableBlurEffectBoundVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableBlurEffectBoundVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlurEffectBoundVariables(val *BlurEffectBoundVariables) *NullableBlurEffectBoundVariables {
	return &NullableBlurEffectBoundVariables{value: val, isSet: true}
}

func (v NullableBlurEffectBoundVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlurEffectBoundVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


