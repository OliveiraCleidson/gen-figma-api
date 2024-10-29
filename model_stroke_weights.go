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

// checks if the StrokeWeights type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StrokeWeights{}

// StrokeWeights Individual stroke weights
type StrokeWeights struct {
	// The top stroke weight.
	Top float32 `json:"top"`
	// The right stroke weight.
	Right float32 `json:"right"`
	// The bottom stroke weight.
	Bottom float32 `json:"bottom"`
	// The left stroke weight.
	Left float32 `json:"left"`
}

type _StrokeWeights StrokeWeights

// NewStrokeWeights instantiates a new StrokeWeights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStrokeWeights(top float32, right float32, bottom float32, left float32) *StrokeWeights {
	this := StrokeWeights{}
	this.Top = top
	this.Right = right
	this.Bottom = bottom
	this.Left = left
	return &this
}

// NewStrokeWeightsWithDefaults instantiates a new StrokeWeights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStrokeWeightsWithDefaults() *StrokeWeights {
	this := StrokeWeights{}
	return &this
}

// GetTop returns the Top field value
func (o *StrokeWeights) GetTop() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Top
}

// GetTopOk returns a tuple with the Top field value
// and a boolean to check if the value has been set.
func (o *StrokeWeights) GetTopOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Top, true
}

// SetTop sets field value
func (o *StrokeWeights) SetTop(v float32) {
	o.Top = v
}

// GetRight returns the Right field value
func (o *StrokeWeights) GetRight() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Right
}

// GetRightOk returns a tuple with the Right field value
// and a boolean to check if the value has been set.
func (o *StrokeWeights) GetRightOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Right, true
}

// SetRight sets field value
func (o *StrokeWeights) SetRight(v float32) {
	o.Right = v
}

// GetBottom returns the Bottom field value
func (o *StrokeWeights) GetBottom() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Bottom
}

// GetBottomOk returns a tuple with the Bottom field value
// and a boolean to check if the value has been set.
func (o *StrokeWeights) GetBottomOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bottom, true
}

// SetBottom sets field value
func (o *StrokeWeights) SetBottom(v float32) {
	o.Bottom = v
}

// GetLeft returns the Left field value
func (o *StrokeWeights) GetLeft() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Left
}

// GetLeftOk returns a tuple with the Left field value
// and a boolean to check if the value has been set.
func (o *StrokeWeights) GetLeftOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Left, true
}

// SetLeft sets field value
func (o *StrokeWeights) SetLeft(v float32) {
	o.Left = v
}

func (o StrokeWeights) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StrokeWeights) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["top"] = o.Top
	toSerialize["right"] = o.Right
	toSerialize["bottom"] = o.Bottom
	toSerialize["left"] = o.Left
	return toSerialize, nil
}

func (o *StrokeWeights) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"top",
		"right",
		"bottom",
		"left",
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

	varStrokeWeights := _StrokeWeights{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStrokeWeights)

	if err != nil {
		return err
	}

	*o = StrokeWeights(varStrokeWeights)

	return err
}

type NullableStrokeWeights struct {
	value *StrokeWeights
	isSet bool
}

func (v NullableStrokeWeights) Get() *StrokeWeights {
	return v.value
}

func (v *NullableStrokeWeights) Set(val *StrokeWeights) {
	v.value = val
	v.isSet = true
}

func (v NullableStrokeWeights) IsSet() bool {
	return v.isSet
}

func (v *NullableStrokeWeights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStrokeWeights(val *StrokeWeights) *NullableStrokeWeights {
	return &NullableStrokeWeights{value: val, isSet: true}
}

func (v NullableStrokeWeights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStrokeWeights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


