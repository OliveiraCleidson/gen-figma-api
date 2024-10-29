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

// checks if the ComponentProperty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentProperty{}

// ComponentProperty A property of a component.
type ComponentProperty struct {
	// Type of this component property.
	Type ComponentPropertyType `json:"type"`
	Value ComponentPropertyValue `json:"value"`
	// Preferred values for this property. Only applicable if type is `INSTANCE_SWAP`.
	PreferredValues []InstanceSwapPreferredValue `json:"preferredValues,omitempty"`
	BoundVariables *ComponentPropertyBoundVariables `json:"boundVariables,omitempty"`
}

type _ComponentProperty ComponentProperty

// NewComponentProperty instantiates a new ComponentProperty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentProperty(type_ ComponentPropertyType, value ComponentPropertyValue) *ComponentProperty {
	this := ComponentProperty{}
	this.Type = type_
	this.Value = value
	return &this
}

// NewComponentPropertyWithDefaults instantiates a new ComponentProperty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentPropertyWithDefaults() *ComponentProperty {
	this := ComponentProperty{}
	return &this
}

// GetType returns the Type field value
func (o *ComponentProperty) GetType() ComponentPropertyType {
	if o == nil {
		var ret ComponentPropertyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ComponentProperty) GetTypeOk() (*ComponentPropertyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ComponentProperty) SetType(v ComponentPropertyType) {
	o.Type = v
}

// GetValue returns the Value field value
func (o *ComponentProperty) GetValue() ComponentPropertyValue {
	if o == nil {
		var ret ComponentPropertyValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ComponentProperty) GetValueOk() (*ComponentPropertyValue, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ComponentProperty) SetValue(v ComponentPropertyValue) {
	o.Value = v
}

// GetPreferredValues returns the PreferredValues field value if set, zero value otherwise.
func (o *ComponentProperty) GetPreferredValues() []InstanceSwapPreferredValue {
	if o == nil || IsNil(o.PreferredValues) {
		var ret []InstanceSwapPreferredValue
		return ret
	}
	return o.PreferredValues
}

// GetPreferredValuesOk returns a tuple with the PreferredValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentProperty) GetPreferredValuesOk() ([]InstanceSwapPreferredValue, bool) {
	if o == nil || IsNil(o.PreferredValues) {
		return nil, false
	}
	return o.PreferredValues, true
}

// HasPreferredValues returns a boolean if a field has been set.
func (o *ComponentProperty) HasPreferredValues() bool {
	if o != nil && !IsNil(o.PreferredValues) {
		return true
	}

	return false
}

// SetPreferredValues gets a reference to the given []InstanceSwapPreferredValue and assigns it to the PreferredValues field.
func (o *ComponentProperty) SetPreferredValues(v []InstanceSwapPreferredValue) {
	o.PreferredValues = v
}

// GetBoundVariables returns the BoundVariables field value if set, zero value otherwise.
func (o *ComponentProperty) GetBoundVariables() ComponentPropertyBoundVariables {
	if o == nil || IsNil(o.BoundVariables) {
		var ret ComponentPropertyBoundVariables
		return ret
	}
	return *o.BoundVariables
}

// GetBoundVariablesOk returns a tuple with the BoundVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentProperty) GetBoundVariablesOk() (*ComponentPropertyBoundVariables, bool) {
	if o == nil || IsNil(o.BoundVariables) {
		return nil, false
	}
	return o.BoundVariables, true
}

// HasBoundVariables returns a boolean if a field has been set.
func (o *ComponentProperty) HasBoundVariables() bool {
	if o != nil && !IsNil(o.BoundVariables) {
		return true
	}

	return false
}

// SetBoundVariables gets a reference to the given ComponentPropertyBoundVariables and assigns it to the BoundVariables field.
func (o *ComponentProperty) SetBoundVariables(v ComponentPropertyBoundVariables) {
	o.BoundVariables = &v
}

func (o ComponentProperty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentProperty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["value"] = o.Value
	if !IsNil(o.PreferredValues) {
		toSerialize["preferredValues"] = o.PreferredValues
	}
	if !IsNil(o.BoundVariables) {
		toSerialize["boundVariables"] = o.BoundVariables
	}
	return toSerialize, nil
}

func (o *ComponentProperty) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"value",
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

	varComponentProperty := _ComponentProperty{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varComponentProperty)

	if err != nil {
		return err
	}

	*o = ComponentProperty(varComponentProperty)

	return err
}

type NullableComponentProperty struct {
	value *ComponentProperty
	isSet bool
}

func (v NullableComponentProperty) Get() *ComponentProperty {
	return v.value
}

func (v *NullableComponentProperty) Set(val *ComponentProperty) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentProperty) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentProperty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentProperty(val *ComponentProperty) *NullableComponentProperty {
	return &NullableComponentProperty{value: val, isSet: true}
}

func (v NullableComponentProperty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentProperty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

