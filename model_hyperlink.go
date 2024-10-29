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

// checks if the Hyperlink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Hyperlink{}

// Hyperlink A link to either a URL or another frame (node) in the document.
type Hyperlink struct {
	// The type of hyperlink. Can be either `URL` or `NODE`.
	Type string `json:"type"`
	// The URL that the hyperlink points to, if `type` is `URL`.
	Url *string `json:"url,omitempty"`
	// The ID of the node that the hyperlink points to, if `type` is `NODE`.
	NodeID *string `json:"nodeID,omitempty"`
}

type _Hyperlink Hyperlink

// NewHyperlink instantiates a new Hyperlink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperlink(type_ string) *Hyperlink {
	this := Hyperlink{}
	this.Type = type_
	return &this
}

// NewHyperlinkWithDefaults instantiates a new Hyperlink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperlinkWithDefaults() *Hyperlink {
	this := Hyperlink{}
	return &this
}

// GetType returns the Type field value
func (o *Hyperlink) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Hyperlink) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Hyperlink) SetType(v string) {
	o.Type = v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Hyperlink) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hyperlink) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Hyperlink) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Hyperlink) SetUrl(v string) {
	o.Url = &v
}

// GetNodeID returns the NodeID field value if set, zero value otherwise.
func (o *Hyperlink) GetNodeID() string {
	if o == nil || IsNil(o.NodeID) {
		var ret string
		return ret
	}
	return *o.NodeID
}

// GetNodeIDOk returns a tuple with the NodeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hyperlink) GetNodeIDOk() (*string, bool) {
	if o == nil || IsNil(o.NodeID) {
		return nil, false
	}
	return o.NodeID, true
}

// HasNodeID returns a boolean if a field has been set.
func (o *Hyperlink) HasNodeID() bool {
	if o != nil && !IsNil(o.NodeID) {
		return true
	}

	return false
}

// SetNodeID gets a reference to the given string and assigns it to the NodeID field.
func (o *Hyperlink) SetNodeID(v string) {
	o.NodeID = &v
}

func (o Hyperlink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Hyperlink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.NodeID) {
		toSerialize["nodeID"] = o.NodeID
	}
	return toSerialize, nil
}

func (o *Hyperlink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varHyperlink := _Hyperlink{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHyperlink)

	if err != nil {
		return err
	}

	*o = Hyperlink(varHyperlink)

	return err
}

type NullableHyperlink struct {
	value *Hyperlink
	isSet bool
}

func (v NullableHyperlink) Get() *Hyperlink {
	return v.value
}

func (v *NullableHyperlink) Set(val *Hyperlink) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperlink) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperlink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperlink(val *Hyperlink) *NullableHyperlink {
	return &NullableHyperlink{value: val, isSet: true}
}

func (v NullableHyperlink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperlink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

