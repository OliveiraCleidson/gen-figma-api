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

// checks if the ActivityLogWidgetEntity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityLogWidgetEntity{}

// ActivityLogWidgetEntity A Figma widget
type ActivityLogWidgetEntity struct {
	// The type of entity.
	Type string `json:"type"`
	// Unique identifier of the widget.
	Id string `json:"id"`
	// Name of the widget.
	Name string `json:"name"`
	// Indicates if the object is a widget available on Figma Design or FigJam.
	EditorType string `json:"editor_type"`
}

type _ActivityLogWidgetEntity ActivityLogWidgetEntity

// NewActivityLogWidgetEntity instantiates a new ActivityLogWidgetEntity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityLogWidgetEntity(type_ string, id string, name string, editorType string) *ActivityLogWidgetEntity {
	this := ActivityLogWidgetEntity{}
	this.Type = type_
	this.Id = id
	this.Name = name
	this.EditorType = editorType
	return &this
}

// NewActivityLogWidgetEntityWithDefaults instantiates a new ActivityLogWidgetEntity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityLogWidgetEntityWithDefaults() *ActivityLogWidgetEntity {
	this := ActivityLogWidgetEntity{}
	return &this
}

// GetType returns the Type field value
func (o *ActivityLogWidgetEntity) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ActivityLogWidgetEntity) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ActivityLogWidgetEntity) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ActivityLogWidgetEntity) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ActivityLogWidgetEntity) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ActivityLogWidgetEntity) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ActivityLogWidgetEntity) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ActivityLogWidgetEntity) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ActivityLogWidgetEntity) SetName(v string) {
	o.Name = v
}

// GetEditorType returns the EditorType field value
func (o *ActivityLogWidgetEntity) GetEditorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EditorType
}

// GetEditorTypeOk returns a tuple with the EditorType field value
// and a boolean to check if the value has been set.
func (o *ActivityLogWidgetEntity) GetEditorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EditorType, true
}

// SetEditorType sets field value
func (o *ActivityLogWidgetEntity) SetEditorType(v string) {
	o.EditorType = v
}

func (o ActivityLogWidgetEntity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityLogWidgetEntity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["editor_type"] = o.EditorType
	return toSerialize, nil
}

func (o *ActivityLogWidgetEntity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"name",
		"editor_type",
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

	varActivityLogWidgetEntity := _ActivityLogWidgetEntity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivityLogWidgetEntity)

	if err != nil {
		return err
	}

	*o = ActivityLogWidgetEntity(varActivityLogWidgetEntity)

	return err
}

type NullableActivityLogWidgetEntity struct {
	value *ActivityLogWidgetEntity
	isSet bool
}

func (v NullableActivityLogWidgetEntity) Get() *ActivityLogWidgetEntity {
	return v.value
}

func (v *NullableActivityLogWidgetEntity) Set(val *ActivityLogWidgetEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityLogWidgetEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityLogWidgetEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityLogWidgetEntity(val *ActivityLogWidgetEntity) *NullableActivityLogWidgetEntity {
	return &NullableActivityLogWidgetEntity{value: val, isSet: true}
}

func (v NullableActivityLogWidgetEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityLogWidgetEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


