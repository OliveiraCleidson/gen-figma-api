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

// checks if the FrameOffset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameOffset{}

// FrameOffset Position of a comment relative to the frame to which it is attached.
type FrameOffset struct {
	// Unique id specifying the frame.
	NodeId string `json:"node_id"`
	// 2D vector offset within the frame from the top-left corner.
	NodeOffset Vector `json:"node_offset"`
}

type _FrameOffset FrameOffset

// NewFrameOffset instantiates a new FrameOffset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameOffset(nodeId string, nodeOffset Vector) *FrameOffset {
	this := FrameOffset{}
	this.NodeId = nodeId
	this.NodeOffset = nodeOffset
	return &this
}

// NewFrameOffsetWithDefaults instantiates a new FrameOffset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameOffsetWithDefaults() *FrameOffset {
	this := FrameOffset{}
	return &this
}

// GetNodeId returns the NodeId field value
func (o *FrameOffset) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *FrameOffset) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *FrameOffset) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeOffset returns the NodeOffset field value
func (o *FrameOffset) GetNodeOffset() Vector {
	if o == nil {
		var ret Vector
		return ret
	}

	return o.NodeOffset
}

// GetNodeOffsetOk returns a tuple with the NodeOffset field value
// and a boolean to check if the value has been set.
func (o *FrameOffset) GetNodeOffsetOk() (*Vector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeOffset, true
}

// SetNodeOffset sets field value
func (o *FrameOffset) SetNodeOffset(v Vector) {
	o.NodeOffset = v
}

func (o FrameOffset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameOffset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_offset"] = o.NodeOffset
	return toSerialize, nil
}

func (o *FrameOffset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_id",
		"node_offset",
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

	varFrameOffset := _FrameOffset{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFrameOffset)

	if err != nil {
		return err
	}

	*o = FrameOffset(varFrameOffset)

	return err
}

type NullableFrameOffset struct {
	value *FrameOffset
	isSet bool
}

func (v NullableFrameOffset) Get() *FrameOffset {
	return v.value
}

func (v *NullableFrameOffset) Set(val *FrameOffset) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameOffset) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameOffset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameOffset(val *FrameOffset) *NullableFrameOffset {
	return &NullableFrameOffset{value: val, isSet: true}
}

func (v NullableFrameOffset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameOffset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


