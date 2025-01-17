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

// checks if the ComponentSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentSet{}

// ComponentSet A description of a component set, which is a node containing a set of variants of a component.
type ComponentSet struct {
	// The key of the component set
	Key string `json:"key"`
	// Name of the component set
	Name string `json:"name"`
	// The description of the component set as entered in the editor
	Description string `json:"description"`
	// An array of documentation links attached to this component set
	DocumentationLinks []DocumentationLink `json:"documentationLinks,omitempty"`
	// Whether this component set is a remote component set that doesn't live in this file
	Remote *bool `json:"remote,omitempty"`
}

type _ComponentSet ComponentSet

// NewComponentSet instantiates a new ComponentSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentSet(key string, name string, description string) *ComponentSet {
	this := ComponentSet{}
	this.Key = key
	this.Name = name
	this.Description = description
	return &this
}

// NewComponentSetWithDefaults instantiates a new ComponentSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentSetWithDefaults() *ComponentSet {
	this := ComponentSet{}
	return &this
}

// GetKey returns the Key field value
func (o *ComponentSet) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ComponentSet) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ComponentSet) SetKey(v string) {
	o.Key = v
}

// GetName returns the Name field value
func (o *ComponentSet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ComponentSet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ComponentSet) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ComponentSet) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ComponentSet) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ComponentSet) SetDescription(v string) {
	o.Description = v
}

// GetDocumentationLinks returns the DocumentationLinks field value if set, zero value otherwise.
func (o *ComponentSet) GetDocumentationLinks() []DocumentationLink {
	if o == nil || IsNil(o.DocumentationLinks) {
		var ret []DocumentationLink
		return ret
	}
	return o.DocumentationLinks
}

// GetDocumentationLinksOk returns a tuple with the DocumentationLinks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentSet) GetDocumentationLinksOk() ([]DocumentationLink, bool) {
	if o == nil || IsNil(o.DocumentationLinks) {
		return nil, false
	}
	return o.DocumentationLinks, true
}

// HasDocumentationLinks returns a boolean if a field has been set.
func (o *ComponentSet) HasDocumentationLinks() bool {
	if o != nil && !IsNil(o.DocumentationLinks) {
		return true
	}

	return false
}

// SetDocumentationLinks gets a reference to the given []DocumentationLink and assigns it to the DocumentationLinks field.
func (o *ComponentSet) SetDocumentationLinks(v []DocumentationLink) {
	o.DocumentationLinks = v
}

// GetRemote returns the Remote field value if set, zero value otherwise.
func (o *ComponentSet) GetRemote() bool {
	if o == nil || IsNil(o.Remote) {
		var ret bool
		return ret
	}
	return *o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentSet) GetRemoteOk() (*bool, bool) {
	if o == nil || IsNil(o.Remote) {
		return nil, false
	}
	return o.Remote, true
}

// HasRemote returns a boolean if a field has been set.
func (o *ComponentSet) HasRemote() bool {
	if o != nil && !IsNil(o.Remote) {
		return true
	}

	return false
}

// SetRemote gets a reference to the given bool and assigns it to the Remote field.
func (o *ComponentSet) SetRemote(v bool) {
	o.Remote = &v
}

func (o ComponentSet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	if !IsNil(o.DocumentationLinks) {
		toSerialize["documentationLinks"] = o.DocumentationLinks
	}
	if !IsNil(o.Remote) {
		toSerialize["remote"] = o.Remote
	}
	return toSerialize, nil
}

func (o *ComponentSet) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
		"name",
		"description",
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

	varComponentSet := _ComponentSet{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varComponentSet)

	if err != nil {
		return err
	}

	*o = ComponentSet(varComponentSet)

	return err
}

type NullableComponentSet struct {
	value *ComponentSet
	isSet bool
}

func (v NullableComponentSet) Get() *ComponentSet {
	return v.value
}

func (v *NullableComponentSet) Set(val *ComponentSet) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentSet) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentSet(val *ComponentSet) *NullableComponentSet {
	return &NullableComponentSet{value: val, isSet: true}
}

func (v NullableComponentSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


