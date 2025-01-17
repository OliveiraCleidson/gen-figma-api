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

// checks if the EasingEasingFunctionSpring type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EasingEasingFunctionSpring{}

// EasingEasingFunctionSpring A spring function that defines the easing.
type EasingEasingFunctionSpring struct {
	Mass float32 `json:"mass"`
	Stiffness float32 `json:"stiffness"`
	Damping float32 `json:"damping"`
}

type _EasingEasingFunctionSpring EasingEasingFunctionSpring

// NewEasingEasingFunctionSpring instantiates a new EasingEasingFunctionSpring object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasingEasingFunctionSpring(mass float32, stiffness float32, damping float32) *EasingEasingFunctionSpring {
	this := EasingEasingFunctionSpring{}
	this.Mass = mass
	this.Stiffness = stiffness
	this.Damping = damping
	return &this
}

// NewEasingEasingFunctionSpringWithDefaults instantiates a new EasingEasingFunctionSpring object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasingEasingFunctionSpringWithDefaults() *EasingEasingFunctionSpring {
	this := EasingEasingFunctionSpring{}
	return &this
}

// GetMass returns the Mass field value
func (o *EasingEasingFunctionSpring) GetMass() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Mass
}

// GetMassOk returns a tuple with the Mass field value
// and a boolean to check if the value has been set.
func (o *EasingEasingFunctionSpring) GetMassOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mass, true
}

// SetMass sets field value
func (o *EasingEasingFunctionSpring) SetMass(v float32) {
	o.Mass = v
}

// GetStiffness returns the Stiffness field value
func (o *EasingEasingFunctionSpring) GetStiffness() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Stiffness
}

// GetStiffnessOk returns a tuple with the Stiffness field value
// and a boolean to check if the value has been set.
func (o *EasingEasingFunctionSpring) GetStiffnessOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stiffness, true
}

// SetStiffness sets field value
func (o *EasingEasingFunctionSpring) SetStiffness(v float32) {
	o.Stiffness = v
}

// GetDamping returns the Damping field value
func (o *EasingEasingFunctionSpring) GetDamping() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Damping
}

// GetDampingOk returns a tuple with the Damping field value
// and a boolean to check if the value has been set.
func (o *EasingEasingFunctionSpring) GetDampingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Damping, true
}

// SetDamping sets field value
func (o *EasingEasingFunctionSpring) SetDamping(v float32) {
	o.Damping = v
}

func (o EasingEasingFunctionSpring) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EasingEasingFunctionSpring) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mass"] = o.Mass
	toSerialize["stiffness"] = o.Stiffness
	toSerialize["damping"] = o.Damping
	return toSerialize, nil
}

func (o *EasingEasingFunctionSpring) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mass",
		"stiffness",
		"damping",
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

	varEasingEasingFunctionSpring := _EasingEasingFunctionSpring{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEasingEasingFunctionSpring)

	if err != nil {
		return err
	}

	*o = EasingEasingFunctionSpring(varEasingEasingFunctionSpring)

	return err
}

type NullableEasingEasingFunctionSpring struct {
	value *EasingEasingFunctionSpring
	isSet bool
}

func (v NullableEasingEasingFunctionSpring) Get() *EasingEasingFunctionSpring {
	return v.value
}

func (v *NullableEasingEasingFunctionSpring) Set(val *EasingEasingFunctionSpring) {
	v.value = val
	v.isSet = true
}

func (v NullableEasingEasingFunctionSpring) IsSet() bool {
	return v.isSet
}

func (v *NullableEasingEasingFunctionSpring) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasingEasingFunctionSpring(val *EasingEasingFunctionSpring) *NullableEasingEasingFunctionSpring {
	return &NullableEasingEasingFunctionSpring{value: val, isSet: true}
}

func (v NullableEasingEasingFunctionSpring) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasingEasingFunctionSpring) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


