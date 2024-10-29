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

// checks if the Component type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Component{}

// Component A description of a main component. Helps you identify which component instances are attached to.
type Component struct {
	// The key of the component
	Key string `json:"key"`
	// Name of the component
	Name string `json:"name"`
	// The description of the component as entered in the editor
	Description string `json:"description"`
	// The ID of the component set if the component belongs to one
	ComponentSetId *string `json:"componentSetId,omitempty"`
	// An array of documentation links attached to this component
	DocumentationLinks []DocumentationLink `json:"documentationLinks"`
	// Whether this component is a remote component that doesn't live in this file
	Remote bool `json:"remote"`
}

type _Component Component

// NewComponent instantiates a new Component object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponent(key string, name string, description string, documentationLinks []DocumentationLink, remote bool) *Component {
	this := Component{}
	this.Key = key
	this.Name = name
	this.Description = description
	this.DocumentationLinks = documentationLinks
	this.Remote = remote
	return &this
}

// NewComponentWithDefaults instantiates a new Component object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentWithDefaults() *Component {
	this := Component{}
	return &this
}

// GetKey returns the Key field value
func (o *Component) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Component) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *Component) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *Component) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Component) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Component) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Component) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Component) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Component) SetDescription(v string) {
	o.Description = v
}

// GetComponentSetId returns the ComponentSetId field value if set, zero value otherwise.
func (o *Component) GetComponentSetId() string {
	if o == nil || IsNil(o.ComponentSetId) {
		var ret string
		return ret
	}
	return *o.ComponentSetId
}

// GetComponentSetIdOk returns a tuple with the ComponentSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Component) GetComponentSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.ComponentSetId) {
		return nil, false
	}
	return o.ComponentSetId, true
}

// HasComponentSetId returns a boolean if a field has been set.
func (o *Component) HasComponentSetId() bool {
	if o != nil && !IsNil(o.ComponentSetId) {
		return true
	}

	return false
}

// SetComponentSetId gets a reference to the given string and assigns it to the ComponentSetId field.
func (o *Component) SetComponentSetId(v string) {
	o.ComponentSetId = &v
}

// GetDocumentationLinks returns the DocumentationLinks field value
func (o *Component) GetDocumentationLinks() []DocumentationLink {
	if o == nil {
		var ret []DocumentationLink
		return ret
	}

	return o.DocumentationLinks
}

// GetDocumentationLinksOk returns a tuple with the DocumentationLinks field value
// and a boolean to check if the value has been set.
func (o *Component) GetDocumentationLinksOk() ([]DocumentationLink, bool) {
	if o == nil {
		return nil, false
	}
	return o.DocumentationLinks, true
}

// SetDocumentationLinks sets field value
func (o *Component) SetDocumentationLinks(v []DocumentationLink) {
	o.DocumentationLinks = v
}

// GetRemote returns the Remote field value
func (o *Component) GetRemote() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value
// and a boolean to check if the value has been set.
func (o *Component) GetRemoteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remote, true
}

// SetRemote sets field value
func (o *Component) SetRemote(v bool) {
	o.Remote = v
}

func (o Component) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Component) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.ComponentSetId) {
		toSerialize["componentSetId"] = o.ComponentSetId
	}
	toSerialize["documentationLinks"] = o.DocumentationLinks
	toSerialize["remote"] = o.Remote
	return toSerialize, nil
}

func (o *Component) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"name",
		"description",
		"documentationLinks",
		"remote",
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

	varComponent := _Component{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varComponent)

	if err != nil {
		return err
	}

	*o = Component(varComponent)

	return err
}

type NullableComponent struct {
	value *Component
	isSet bool
}

func (v NullableComponent) Get() *Component {
	return v.value
}

func (v *NullableComponent) Set(val *Component) {
	v.value = val
	v.isSet = true
}

func (v NullableComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponent(val *Component) *NullableComponent {
	return &NullableComponent{value: val, isSet: true}
}

func (v NullableComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


