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

// checks if the VariableAlias type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableAlias{}

// VariableAlias Contains a variable alias
type VariableAlias struct {
	Type string `json:"type"`
	// The id of the variable that the current variable is aliased to. This variable can be a local or remote variable, and both can be retrieved via the GET /v1/files/:file_key/variables/local endpoint.
	Id string `json:"id"`
}

type _VariableAlias VariableAlias

// NewVariableAlias instantiates a new VariableAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableAlias(type_ string, id string) *VariableAlias {
	this := VariableAlias{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewVariableAliasWithDefaults instantiates a new VariableAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableAliasWithDefaults() *VariableAlias {
	this := VariableAlias{}
	return &this
}

// GetType returns the Type field value
func (o *VariableAlias) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VariableAlias) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VariableAlias) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *VariableAlias) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VariableAlias) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VariableAlias) SetId(v string) {
	o.Id = v
}

func (o VariableAlias) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableAlias) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *VariableAlias) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varVariableAlias := _VariableAlias{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVariableAlias)

	if err != nil {
		return err
	}

	*o = VariableAlias(varVariableAlias)

	return err
}

type NullableVariableAlias struct {
	value *VariableAlias
	isSet bool
}

func (v NullableVariableAlias) Get() *VariableAlias {
	return v.value
}

func (v *NullableVariableAlias) Set(val *VariableAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableAlias(val *VariableAlias) *NullableVariableAlias {
	return &NullableVariableAlias{value: val, isSet: true}
}

func (v NullableVariableAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


