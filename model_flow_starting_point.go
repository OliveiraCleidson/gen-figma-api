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

// checks if the FlowStartingPoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowStartingPoint{}

// FlowStartingPoint A flow starting point used when launching a prototype to enter Presentation view.
type FlowStartingPoint struct {
	// Unique identifier specifying the frame.
	NodeId string `json:"nodeId"`
	// Name of flow.
	Name string `json:"name"`
}

type _FlowStartingPoint FlowStartingPoint

// NewFlowStartingPoint instantiates a new FlowStartingPoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowStartingPoint(nodeId string, name string) *FlowStartingPoint {
	this := FlowStartingPoint{}
	this.NodeId = nodeId
	this.Name = name
	return &this
}

// NewFlowStartingPointWithDefaults instantiates a new FlowStartingPoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowStartingPointWithDefaults() *FlowStartingPoint {
	this := FlowStartingPoint{}
	return &this
}

// GetNodeId returns the NodeId field value
func (o *FlowStartingPoint) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *FlowStartingPoint) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *FlowStartingPoint) SetNodeId(v string) {
	o.NodeId = v
}

// GetName returns the Name field value
func (o *FlowStartingPoint) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FlowStartingPoint) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FlowStartingPoint) SetName(v string) {
	o.Name = v
}

func (o FlowStartingPoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowStartingPoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodeId"] = o.NodeId
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *FlowStartingPoint) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nodeId",
		"name",
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

	varFlowStartingPoint := _FlowStartingPoint{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFlowStartingPoint)

	if err != nil {
		return err
	}

	*o = FlowStartingPoint(varFlowStartingPoint)

	return err
}

type NullableFlowStartingPoint struct {
	value *FlowStartingPoint
	isSet bool
}

func (v NullableFlowStartingPoint) Get() *FlowStartingPoint {
	return v.value
}

func (v *NullableFlowStartingPoint) Set(val *FlowStartingPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowStartingPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowStartingPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowStartingPoint(val *FlowStartingPoint) *NullableFlowStartingPoint {
	return &NullableFlowStartingPoint{value: val, isSet: true}
}

func (v NullableFlowStartingPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowStartingPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

