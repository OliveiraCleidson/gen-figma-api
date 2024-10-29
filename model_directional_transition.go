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

// checks if the DirectionalTransition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectionalTransition{}

// DirectionalTransition Describes an animation used when navigating in a prototype.
type DirectionalTransition struct {
	Type string `json:"type"`
	Direction string `json:"direction"`
	// The duration of the transition in milliseconds.
	Duration float32 `json:"duration"`
	// The easing curve of the transition.
	Easing Easing `json:"easing"`
	// When the transition `type` is `\"SMART_ANIMATE\"` or when `matchLayers` is `true`, then the transition will be performed using smart animate, which attempts to match corresponding layers an interpolate other properties during the animation.
	MatchLayers *bool `json:"matchLayers,omitempty"`
}

type _DirectionalTransition DirectionalTransition

// NewDirectionalTransition instantiates a new DirectionalTransition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectionalTransition(type_ string, direction string, duration float32, easing Easing) *DirectionalTransition {
	this := DirectionalTransition{}
	this.Type = type_
	this.Direction = direction
	this.Duration = duration
	this.Easing = easing
	return &this
}

// NewDirectionalTransitionWithDefaults instantiates a new DirectionalTransition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectionalTransitionWithDefaults() *DirectionalTransition {
	this := DirectionalTransition{}
	return &this
}

// GetType returns the Type field value
func (o *DirectionalTransition) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DirectionalTransition) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DirectionalTransition) SetType(v string) {
	o.Type = v
}

// GetDirection returns the Direction field value
func (o *DirectionalTransition) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *DirectionalTransition) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *DirectionalTransition) SetDirection(v string) {
	o.Direction = v
}

// GetDuration returns the Duration field value
func (o *DirectionalTransition) GetDuration() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *DirectionalTransition) GetDurationOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *DirectionalTransition) SetDuration(v float32) {
	o.Duration = v
}

// GetEasing returns the Easing field value
func (o *DirectionalTransition) GetEasing() Easing {
	if o == nil {
		var ret Easing
		return ret
	}

	return o.Easing
}

// GetEasingOk returns a tuple with the Easing field value
// and a boolean to check if the value has been set.
func (o *DirectionalTransition) GetEasingOk() (*Easing, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Easing, true
}

// SetEasing sets field value
func (o *DirectionalTransition) SetEasing(v Easing) {
	o.Easing = v
}

// GetMatchLayers returns the MatchLayers field value if set, zero value otherwise.
func (o *DirectionalTransition) GetMatchLayers() bool {
	if o == nil || IsNil(o.MatchLayers) {
		var ret bool
		return ret
	}
	return *o.MatchLayers
}

// GetMatchLayersOk returns a tuple with the MatchLayers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectionalTransition) GetMatchLayersOk() (*bool, bool) {
	if o == nil || IsNil(o.MatchLayers) {
		return nil, false
	}
	return o.MatchLayers, true
}

// HasMatchLayers returns a boolean if a field has been set.
func (o *DirectionalTransition) HasMatchLayers() bool {
	if o != nil && !IsNil(o.MatchLayers) {
		return true
	}

	return false
}

// SetMatchLayers gets a reference to the given bool and assigns it to the MatchLayers field.
func (o *DirectionalTransition) SetMatchLayers(v bool) {
	o.MatchLayers = &v
}

func (o DirectionalTransition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectionalTransition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["direction"] = o.Direction
	toSerialize["duration"] = o.Duration
	toSerialize["easing"] = o.Easing
	if !IsNil(o.MatchLayers) {
		toSerialize["matchLayers"] = o.MatchLayers
	}
	return toSerialize, nil
}

func (o *DirectionalTransition) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"direction",
		"duration",
		"easing",
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

	varDirectionalTransition := _DirectionalTransition{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDirectionalTransition)

	if err != nil {
		return err
	}

	*o = DirectionalTransition(varDirectionalTransition)

	return err
}

type NullableDirectionalTransition struct {
	value *DirectionalTransition
	isSet bool
}

func (v NullableDirectionalTransition) Get() *DirectionalTransition {
	return v.value
}

func (v *NullableDirectionalTransition) Set(val *DirectionalTransition) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectionalTransition) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectionalTransition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectionalTransition(val *DirectionalTransition) *NullableDirectionalTransition {
	return &NullableDirectionalTransition{value: val, isSet: true}
}

func (v NullableDirectionalTransition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectionalTransition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


