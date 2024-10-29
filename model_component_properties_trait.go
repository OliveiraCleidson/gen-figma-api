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

// checks if the ComponentPropertiesTrait type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentPropertiesTrait{}

// ComponentPropertiesTrait struct for ComponentPropertiesTrait
type ComponentPropertiesTrait struct {
	// A mapping of name to `ComponentPropertyDefinition` for every component property on this component. Each property has a type, defaultValue, and other optional values.
	ComponentPropertyDefinitions map[string]ComponentPropertyDefinition `json:"componentPropertyDefinitions,omitempty"`
}

// NewComponentPropertiesTrait instantiates a new ComponentPropertiesTrait object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentPropertiesTrait() *ComponentPropertiesTrait {
	this := ComponentPropertiesTrait{}
	return &this
}

// NewComponentPropertiesTraitWithDefaults instantiates a new ComponentPropertiesTrait object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentPropertiesTraitWithDefaults() *ComponentPropertiesTrait {
	this := ComponentPropertiesTrait{}
	return &this
}

// GetComponentPropertyDefinitions returns the ComponentPropertyDefinitions field value if set, zero value otherwise.
func (o *ComponentPropertiesTrait) GetComponentPropertyDefinitions() map[string]ComponentPropertyDefinition {
	if o == nil || IsNil(o.ComponentPropertyDefinitions) {
		var ret map[string]ComponentPropertyDefinition
		return ret
	}
	return o.ComponentPropertyDefinitions
}

// GetComponentPropertyDefinitionsOk returns a tuple with the ComponentPropertyDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentPropertiesTrait) GetComponentPropertyDefinitionsOk() (map[string]ComponentPropertyDefinition, bool) {
	if o == nil || IsNil(o.ComponentPropertyDefinitions) {
		return map[string]ComponentPropertyDefinition{}, false
	}
	return o.ComponentPropertyDefinitions, true
}

// HasComponentPropertyDefinitions returns a boolean if a field has been set.
func (o *ComponentPropertiesTrait) HasComponentPropertyDefinitions() bool {
	if o != nil && !IsNil(o.ComponentPropertyDefinitions) {
		return true
	}

	return false
}

// SetComponentPropertyDefinitions gets a reference to the given map[string]ComponentPropertyDefinition and assigns it to the ComponentPropertyDefinitions field.
func (o *ComponentPropertiesTrait) SetComponentPropertyDefinitions(v map[string]ComponentPropertyDefinition) {
	o.ComponentPropertyDefinitions = v
}

func (o ComponentPropertiesTrait) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentPropertiesTrait) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ComponentPropertyDefinitions) {
		toSerialize["componentPropertyDefinitions"] = o.ComponentPropertyDefinitions
	}
	return toSerialize, nil
}

type NullableComponentPropertiesTrait struct {
	value *ComponentPropertiesTrait
	isSet bool
}

func (v NullableComponentPropertiesTrait) Get() *ComponentPropertiesTrait {
	return v.value
}

func (v *NullableComponentPropertiesTrait) Set(val *ComponentPropertiesTrait) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentPropertiesTrait) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentPropertiesTrait) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentPropertiesTrait(val *ComponentPropertiesTrait) *NullableComponentPropertiesTrait {
	return &NullableComponentPropertiesTrait{value: val, isSet: true}
}

func (v NullableComponentPropertiesTrait) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentPropertiesTrait) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


