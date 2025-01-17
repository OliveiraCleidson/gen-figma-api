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

// checks if the HasChildrenTrait type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HasChildrenTrait{}

// HasChildrenTrait struct for HasChildrenTrait
type HasChildrenTrait struct {
	// An array of nodes that are direct children of this node
	Children []SubcanvasNode `json:"children"`
}

type _HasChildrenTrait HasChildrenTrait

// NewHasChildrenTrait instantiates a new HasChildrenTrait object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasChildrenTrait(children []SubcanvasNode) *HasChildrenTrait {
	this := HasChildrenTrait{}
	this.Children = children
	return &this
}

// NewHasChildrenTraitWithDefaults instantiates a new HasChildrenTrait object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasChildrenTraitWithDefaults() *HasChildrenTrait {
	this := HasChildrenTrait{}
	return &this
}

// GetChildren returns the Children field value
func (o *HasChildrenTrait) GetChildren() []SubcanvasNode {
	if o == nil {
		var ret []SubcanvasNode
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *HasChildrenTrait) GetChildrenOk() ([]SubcanvasNode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *HasChildrenTrait) SetChildren(v []SubcanvasNode) {
	o.Children = v
}

func (o HasChildrenTrait) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HasChildrenTrait) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["children"] = o.Children
	return toSerialize, nil
}

func (o *HasChildrenTrait) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"children",
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

	varHasChildrenTrait := _HasChildrenTrait{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHasChildrenTrait)

	if err != nil {
		return err
	}

	*o = HasChildrenTrait(varHasChildrenTrait)

	return err
}

type NullableHasChildrenTrait struct {
	value *HasChildrenTrait
	isSet bool
}

func (v NullableHasChildrenTrait) Get() *HasChildrenTrait {
	return v.value
}

func (v *NullableHasChildrenTrait) Set(val *HasChildrenTrait) {
	v.value = val
	v.isSet = true
}

func (v NullableHasChildrenTrait) IsSet() bool {
	return v.isSet
}

func (v *NullableHasChildrenTrait) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasChildrenTrait(val *HasChildrenTrait) *NullableHasChildrenTrait {
	return &NullableHasChildrenTrait{value: val, isSet: true}
}

func (v NullableHasChildrenTrait) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasChildrenTrait) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


