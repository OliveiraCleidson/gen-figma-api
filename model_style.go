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

// checks if the Style type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Style{}

// Style A set of properties that can be applied to nodes and published. Styles for a property can be created in the corresponding property's panel while editing a file.
type Style struct {
	// The key of the style
	Key string `json:"key"`
	// Name of the style
	Name string `json:"name"`
	// Description of the style
	Description string `json:"description"`
	// Whether this style is a remote style that doesn't live in this file
	Remote bool `json:"remote"`
	StyleType StyleType `json:"styleType"`
}

type _Style Style

// NewStyle instantiates a new Style object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStyle(key string, name string, description string, remote bool, styleType StyleType) *Style {
	this := Style{}
	this.Key = key
	this.Name = name
	this.Description = description
	this.Remote = remote
	this.StyleType = styleType
	return &this
}

// NewStyleWithDefaults instantiates a new Style object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStyleWithDefaults() *Style {
	this := Style{}
	return &this
}

// GetKey returns the Key field value
func (o *Style) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Style) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Style) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *Style) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Style) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Style) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Style) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Style) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Style) SetDescription(v string) {
	o.Description = v
}

// GetRemote returns the Remote field value
func (o *Style) GetRemote() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value
// and a boolean to check if the value has been set.
func (o *Style) GetRemoteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remote, true
}

// SetRemote sets field value
func (o *Style) SetRemote(v bool) {
	o.Remote = v
}

// GetStyleType returns the StyleType field value
func (o *Style) GetStyleType() StyleType {
	if o == nil {
		var ret StyleType
		return ret
	}

	return o.StyleType
}

// GetStyleTypeOk returns a tuple with the StyleType field value
// and a boolean to check if the value has been set.
func (o *Style) GetStyleTypeOk() (*StyleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StyleType, true
}

// SetStyleType sets field value
func (o *Style) SetStyleType(v StyleType) {
	o.StyleType = v
}

func (o Style) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Style) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["remote"] = o.Remote
	toSerialize["styleType"] = o.StyleType
	return toSerialize, nil
}

func (o *Style) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"name",
		"description",
		"remote",
		"styleType",
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

	varStyle := _Style{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStyle)

	if err != nil {
		return err
	}

	*o = Style(varStyle)

	return err
}

type NullableStyle struct {
	value *Style
	isSet bool
}

func (v NullableStyle) Get() *Style {
	return v.value
}

func (v *NullableStyle) Set(val *Style) {
	v.value = val
	v.isSet = true
}

func (v NullableStyle) IsSet() bool {
	return v.isSet
}

func (v *NullableStyle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStyle(val *Style) *NullableStyle {
	return &NullableStyle{value: val, isSet: true}
}

func (v NullableStyle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStyle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


