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

// checks if the ActivityLogWorkspaceEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityLogWorkspaceEntity{}

// ActivityLogWorkspaceEntity Part of the organizational hierarchy of managing files and users within Figma, only available on the Enterprise Plan
type ActivityLogWorkspaceEntity struct {
	// The type of entity.
	Type string `json:"type"`
	// Unique identifier of the workspace.
	Id string `json:"id"`
	// Name of the workspace.
	Name string `json:"name"`
}

type _ActivityLogWorkspaceEntity ActivityLogWorkspaceEntity

// NewActivityLogWorkspaceEntity instantiates a new ActivityLogWorkspaceEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityLogWorkspaceEntity(type_ string, id string, name string) *ActivityLogWorkspaceEntity {
	this := ActivityLogWorkspaceEntity{}
	this.Type = type_
	this.Id = id
	this.Name = name
	return &this
}

// NewActivityLogWorkspaceEntityWithDefaults instantiates a new ActivityLogWorkspaceEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityLogWorkspaceEntityWithDefaults() *ActivityLogWorkspaceEntity {
	this := ActivityLogWorkspaceEntity{}
	return &this
}

// GetType returns the Type field value
func (o *ActivityLogWorkspaceEntity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActivityLogWorkspaceEntity) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ActivityLogWorkspaceEntity) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ActivityLogWorkspaceEntity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ActivityLogWorkspaceEntity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ActivityLogWorkspaceEntity) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ActivityLogWorkspaceEntity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ActivityLogWorkspaceEntity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ActivityLogWorkspaceEntity) SetName(v string) {
	o.Name = v
}

func (o ActivityLogWorkspaceEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityLogWorkspaceEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *ActivityLogWorkspaceEntity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
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

	varActivityLogWorkspaceEntity := _ActivityLogWorkspaceEntity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivityLogWorkspaceEntity)

	if err != nil {
		return err
	}

	*o = ActivityLogWorkspaceEntity(varActivityLogWorkspaceEntity)

	return err
}

type NullableActivityLogWorkspaceEntity struct {
	value *ActivityLogWorkspaceEntity
	isSet bool
}

func (v NullableActivityLogWorkspaceEntity) Get() *ActivityLogWorkspaceEntity {
	return v.value
}

func (v *NullableActivityLogWorkspaceEntity) Set(val *ActivityLogWorkspaceEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityLogWorkspaceEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityLogWorkspaceEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityLogWorkspaceEntity(val *ActivityLogWorkspaceEntity) *NullableActivityLogWorkspaceEntity {
	return &NullableActivityLogWorkspaceEntity{value: val, isSet: true}
}

func (v NullableActivityLogWorkspaceEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityLogWorkspaceEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


