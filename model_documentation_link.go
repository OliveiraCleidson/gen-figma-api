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

// checks if the DocumentationLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentationLink{}

// DocumentationLink Represents a link to documentation for a component or component set.
type DocumentationLink struct {
	// Should be a valid URI (e.g. https://www.figma.com).
	Uri string `json:"uri"`
}

type _DocumentationLink DocumentationLink

// NewDocumentationLink instantiates a new DocumentationLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentationLink(uri string) *DocumentationLink {
	this := DocumentationLink{}
	this.Uri = uri
	return &this
}

// NewDocumentationLinkWithDefaults instantiates a new DocumentationLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentationLinkWithDefaults() *DocumentationLink {
	this := DocumentationLink{}
	return &this
}

// GetUri returns the Uri field value
func (o *DocumentationLink) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *DocumentationLink) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *DocumentationLink) SetUri(v string) {
	o.Uri = v
}

func (o DocumentationLink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentationLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uri"] = o.Uri
	return toSerialize, nil
}

func (o *DocumentationLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uri",
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

	varDocumentationLink := _DocumentationLink{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDocumentationLink)

	if err != nil {
		return err
	}

	*o = DocumentationLink(varDocumentationLink)

	return err
}

type NullableDocumentationLink struct {
	value *DocumentationLink
	isSet bool
}

func (v NullableDocumentationLink) Get() *DocumentationLink {
	return v.value
}

func (v *NullableDocumentationLink) Set(val *DocumentationLink) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentationLink) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentationLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentationLink(val *DocumentationLink) *NullableDocumentationLink {
	return &NullableDocumentationLink{value: val, isSet: true}
}

func (v NullableDocumentationLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentationLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

