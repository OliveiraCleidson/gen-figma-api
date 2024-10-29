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

// checks if the MinimalFillsTrait type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MinimalFillsTrait{}

// MinimalFillsTrait struct for MinimalFillsTrait
type MinimalFillsTrait struct {
	// An array of fill paints applied to the node.
	Fills []Paint `json:"fills"`
	// A mapping of a StyleType to style ID (see Style) of styles present on this node. The style ID can be used to look up more information about the style in the top-level styles field.
	Styles map[string]string `json:"styles,omitempty"`
}

type _MinimalFillsTrait MinimalFillsTrait

// NewMinimalFillsTrait instantiates a new MinimalFillsTrait object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMinimalFillsTrait(fills []Paint) *MinimalFillsTrait {
	this := MinimalFillsTrait{}
	this.Fills = fills
	return &this
}

// NewMinimalFillsTraitWithDefaults instantiates a new MinimalFillsTrait object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMinimalFillsTraitWithDefaults() *MinimalFillsTrait {
	this := MinimalFillsTrait{}
	return &this
}

// GetFills returns the Fills field value
func (o *MinimalFillsTrait) GetFills() []Paint {
	if o == nil {
		var ret []Paint
		return ret
	}

	return o.Fills
}

// GetFillsOk returns a tuple with the Fills field value
// and a boolean to check if the value has been set.
func (o *MinimalFillsTrait) GetFillsOk() ([]Paint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fills, true
}

// SetFills sets field value
func (o *MinimalFillsTrait) SetFills(v []Paint) {
	o.Fills = v
}

// GetStyles returns the Styles field value if set, zero value otherwise.
func (o *MinimalFillsTrait) GetStyles() map[string]string {
	if o == nil || IsNil(o.Styles) {
		var ret map[string]string
		return ret
	}
	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MinimalFillsTrait) GetStylesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Styles) {
		return map[string]string{}, false
	}
	return o.Styles, true
}

// HasStyles returns a boolean if a field has been set.
func (o *MinimalFillsTrait) HasStyles() bool {
	if o != nil && !IsNil(o.Styles) {
		return true
	}

	return false
}

// SetStyles gets a reference to the given map[string]string and assigns it to the Styles field.
func (o *MinimalFillsTrait) SetStyles(v map[string]string) {
	o.Styles = v
}

func (o MinimalFillsTrait) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MinimalFillsTrait) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fills"] = o.Fills
	if !IsNil(o.Styles) {
		toSerialize["styles"] = o.Styles
	}
	return toSerialize, nil
}

func (o *MinimalFillsTrait) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fills",
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

	varMinimalFillsTrait := _MinimalFillsTrait{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMinimalFillsTrait)

	if err != nil {
		return err
	}

	*o = MinimalFillsTrait(varMinimalFillsTrait)

	return err
}

type NullableMinimalFillsTrait struct {
	value *MinimalFillsTrait
	isSet bool
}

func (v NullableMinimalFillsTrait) Get() *MinimalFillsTrait {
	return v.value
}

func (v *NullableMinimalFillsTrait) Set(val *MinimalFillsTrait) {
	v.value = val
	v.isSet = true
}

func (v NullableMinimalFillsTrait) IsSet() bool {
	return v.isSet
}

func (v *NullableMinimalFillsTrait) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMinimalFillsTrait(val *MinimalFillsTrait) *NullableMinimalFillsTrait {
	return &NullableMinimalFillsTrait{value: val, isSet: true}
}

func (v NullableMinimalFillsTrait) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMinimalFillsTrait) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

