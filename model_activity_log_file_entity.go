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

// checks if the ActivityLogFileEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityLogFileEntity{}

// ActivityLogFileEntity A Figma Design or FigJam file
type ActivityLogFileEntity struct {
	// The type of entity.
	Type string `json:"type"`
	// Unique identifier of the file.
	Key string `json:"key"`
	// Name of the file.
	Name string `json:"name"`
	// Indicates if the object is a file on Figma Design or FigJam.
	EditorType string `json:"editor_type"`
	// Access policy for users who have the link to the file.
	LinkAccess string `json:"link_access"`
	// Access policy for users who have the link to the file's prototype.
	ProtoLinkAccess string `json:"proto_link_access"`
}

type _ActivityLogFileEntity ActivityLogFileEntity

// NewActivityLogFileEntity instantiates a new ActivityLogFileEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityLogFileEntity(type_ string, key string, name string, editorType string, linkAccess string, protoLinkAccess string) *ActivityLogFileEntity {
	this := ActivityLogFileEntity{}
	this.Type = type_
	this.Key = key
	this.Name = name
	this.EditorType = editorType
	this.LinkAccess = linkAccess
	this.ProtoLinkAccess = protoLinkAccess
	return &this
}

// NewActivityLogFileEntityWithDefaults instantiates a new ActivityLogFileEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityLogFileEntityWithDefaults() *ActivityLogFileEntity {
	this := ActivityLogFileEntity{}
	return &this
}

// GetType returns the Type field value
func (o *ActivityLogFileEntity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActivityLogFileEntity) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ActivityLogFileEntity) SetType(v string) {
	o.Type = v
}

// GetKey returns the Key field value
func (o *ActivityLogFileEntity) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ActivityLogFileEntity) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ActivityLogFileEntity) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *ActivityLogFileEntity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ActivityLogFileEntity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ActivityLogFileEntity) SetName(v string) {
	o.Name = v
}

// GetEditorType returns the EditorType field value
func (o *ActivityLogFileEntity) GetEditorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EditorType
}

// GetEditorTypeOk returns a tuple with the EditorType field value
// and a boolean to check if the value has been set.
func (o *ActivityLogFileEntity) GetEditorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditorType, true
}

// SetEditorType sets field value
func (o *ActivityLogFileEntity) SetEditorType(v string) {
	o.EditorType = v
}

// GetLinkAccess returns the LinkAccess field value
func (o *ActivityLogFileEntity) GetLinkAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LinkAccess
}

// GetLinkAccessOk returns a tuple with the LinkAccess field value
// and a boolean to check if the value has been set.
func (o *ActivityLogFileEntity) GetLinkAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LinkAccess, true
}

// SetLinkAccess sets field value
func (o *ActivityLogFileEntity) SetLinkAccess(v string) {
	o.LinkAccess = v
}

// GetProtoLinkAccess returns the ProtoLinkAccess field value
func (o *ActivityLogFileEntity) GetProtoLinkAccess() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProtoLinkAccess
}

// GetProtoLinkAccessOk returns a tuple with the ProtoLinkAccess field value
// and a boolean to check if the value has been set.
func (o *ActivityLogFileEntity) GetProtoLinkAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtoLinkAccess, true
}

// SetProtoLinkAccess sets field value
func (o *ActivityLogFileEntity) SetProtoLinkAccess(v string) {
	o.ProtoLinkAccess = v
}

func (o ActivityLogFileEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityLogFileEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["editor_type"] = o.EditorType
	toSerialize["link_access"] = o.LinkAccess
	toSerialize["proto_link_access"] = o.ProtoLinkAccess
	return toSerialize, nil
}

func (o *ActivityLogFileEntity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"key",
		"name",
		"editor_type",
		"link_access",
		"proto_link_access",
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

	varActivityLogFileEntity := _ActivityLogFileEntity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivityLogFileEntity)

	if err != nil {
		return err
	}

	*o = ActivityLogFileEntity(varActivityLogFileEntity)

	return err
}

type NullableActivityLogFileEntity struct {
	value *ActivityLogFileEntity
	isSet bool
}

func (v NullableActivityLogFileEntity) Get() *ActivityLogFileEntity {
	return v.value
}

func (v *NullableActivityLogFileEntity) Set(val *ActivityLogFileEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityLogFileEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityLogFileEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityLogFileEntity(val *ActivityLogFileEntity) *NullableActivityLogFileEntity {
	return &NullableActivityLogFileEntity{value: val, isSet: true}
}

func (v NullableActivityLogFileEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityLogFileEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


