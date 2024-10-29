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

// checks if the VariableDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableDelete{}

// VariableDelete An object that contains details about deleting a `Variable`.
type VariableDelete struct {
	// The action to perform for the variable.
	Action string `json:"action"`
	// The id of the variable to delete.
	Id string `json:"id"`
}

type _VariableDelete VariableDelete

// NewVariableDelete instantiates a new VariableDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableDelete(action string, id string) *VariableDelete {
	this := VariableDelete{}
	this.Action = action
	this.Id = id
	return &this
}

// NewVariableDeleteWithDefaults instantiates a new VariableDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableDeleteWithDefaults() *VariableDelete {
	this := VariableDelete{}
	return &this
}

// GetAction returns the Action field value
func (o *VariableDelete) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *VariableDelete) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *VariableDelete) SetAction(v string) {
	o.Action = v
}

// GetId returns the Id field value
func (o *VariableDelete) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VariableDelete) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VariableDelete) SetId(v string) {
	o.Id = v
}

func (o VariableDelete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *VariableDelete) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"id",
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

	varVariableDelete := _VariableDelete{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVariableDelete)

	if err != nil {
		return err
	}

	*o = VariableDelete(varVariableDelete)

	return err
}

type NullableVariableDelete struct {
	value *VariableDelete
	isSet bool
}

func (v NullableVariableDelete) Get() *VariableDelete {
	return v.value
}

func (v *NullableVariableDelete) Set(val *VariableDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableDelete(val *VariableDelete) *NullableVariableDelete {
	return &NullableVariableDelete{value: val, isSet: true}
}

func (v NullableVariableDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


