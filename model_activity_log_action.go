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

// checks if the ActivityLogAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityLogAction{}

// ActivityLogAction The task or activity the actor performed.
type ActivityLogAction struct {
	// The type of the action.
	Type string `json:"type"`
	// Metadata of the action. Each action type supports its own metadata attributes.
	Details map[string]interface{} `json:"details"`
}

type _ActivityLogAction ActivityLogAction

// NewActivityLogAction instantiates a new ActivityLogAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityLogAction(type_ string, details map[string]interface{}) *ActivityLogAction {
	this := ActivityLogAction{}
	this.Type = type_
	this.Details = details
	return &this
}

// NewActivityLogActionWithDefaults instantiates a new ActivityLogAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityLogActionWithDefaults() *ActivityLogAction {
	this := ActivityLogAction{}
	return &this
}

// GetType returns the Type field value
func (o *ActivityLogAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActivityLogAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ActivityLogAction) SetType(v string) {
	o.Type = v
}

// GetDetails returns the Details field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *ActivityLogAction) GetDetails() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActivityLogAction) GetDetailsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Details) {
		return map[string]interface{}{}, false
	}
	return o.Details, true
}

// SetDetails sets field value
func (o *ActivityLogAction) SetDetails(v map[string]interface{}) {
	o.Details = v
}

func (o ActivityLogAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityLogAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

func (o *ActivityLogAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"details",
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

	varActivityLogAction := _ActivityLogAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivityLogAction)

	if err != nil {
		return err
	}

	*o = ActivityLogAction(varActivityLogAction)

	return err
}

type NullableActivityLogAction struct {
	value *ActivityLogAction
	isSet bool
}

func (v NullableActivityLogAction) Get() *ActivityLogAction {
	return v.value
}

func (v *NullableActivityLogAction) Set(val *ActivityLogAction) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityLogAction) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityLogAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityLogAction(val *ActivityLogAction) *NullableActivityLogAction {
	return &NullableActivityLogAction{value: val, isSet: true}
}

func (v NullableActivityLogAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityLogAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


