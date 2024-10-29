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

// checks if the FrameInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrameInfo{}

// FrameInfo Data on the frame a component resides in.
type FrameInfo struct {
	// The ID of the frame node within the file.
	NodeId *string `json:"nodeId,omitempty"`
	// The name of the frame node.
	Name *string `json:"name,omitempty"`
	// The background color of the frame node.
	BackgroundColor *string `json:"backgroundColor,omitempty"`
	// The ID of the page containing the frame node.
	PageId string `json:"pageId"`
	// The name of the page containing the frame node.
	PageName string `json:"pageName"`
}

type _FrameInfo FrameInfo

// NewFrameInfo instantiates a new FrameInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrameInfo(pageId string, pageName string) *FrameInfo {
	this := FrameInfo{}
	this.PageId = pageId
	this.PageName = pageName
	return &this
}

// NewFrameInfoWithDefaults instantiates a new FrameInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrameInfoWithDefaults() *FrameInfo {
	this := FrameInfo{}
	return &this
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *FrameInfo) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameInfo) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *FrameInfo) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *FrameInfo) SetNodeId(v string) {
	o.NodeId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FrameInfo) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameInfo) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FrameInfo) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FrameInfo) SetName(v string) {
	o.Name = &v
}

// GetBackgroundColor returns the BackgroundColor field value if set, zero value otherwise.
func (o *FrameInfo) GetBackgroundColor() string {
	if o == nil || IsNil(o.BackgroundColor) {
		var ret string
		return ret
	}
	return *o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrameInfo) GetBackgroundColorOk() (*string, bool) {
	if o == nil || IsNil(o.BackgroundColor) {
		return nil, false
	}
	return o.BackgroundColor, true
}

// HasBackgroundColor returns a boolean if a field has been set.
func (o *FrameInfo) HasBackgroundColor() bool {
	if o != nil && !IsNil(o.BackgroundColor) {
		return true
	}

	return false
}

// SetBackgroundColor gets a reference to the given string and assigns it to the BackgroundColor field.
func (o *FrameInfo) SetBackgroundColor(v string) {
	o.BackgroundColor = &v
}

// GetPageId returns the PageId field value
func (o *FrameInfo) GetPageId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PageId
}

// GetPageIdOk returns a tuple with the PageId field value
// and a boolean to check if the value has been set.
func (o *FrameInfo) GetPageIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageId, true
}

// SetPageId sets field value
func (o *FrameInfo) SetPageId(v string) {
	o.PageId = v
}

// GetPageName returns the PageName field value
func (o *FrameInfo) GetPageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PageName
}

// GetPageNameOk returns a tuple with the PageName field value
// and a boolean to check if the value has been set.
func (o *FrameInfo) GetPageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageName, true
}

// SetPageName sets field value
func (o *FrameInfo) SetPageName(v string) {
	o.PageName = v
}

func (o FrameInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrameInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NodeId) {
		toSerialize["nodeId"] = o.NodeId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.BackgroundColor) {
		toSerialize["backgroundColor"] = o.BackgroundColor
	}
	toSerialize["pageId"] = o.PageId
	toSerialize["pageName"] = o.PageName
	return toSerialize, nil
}

func (o *FrameInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"pageId",
		"pageName",
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

	varFrameInfo := _FrameInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFrameInfo)

	if err != nil {
		return err
	}

	*o = FrameInfo(varFrameInfo)

	return err
}

type NullableFrameInfo struct {
	value *FrameInfo
	isSet bool
}

func (v NullableFrameInfo) Get() *FrameInfo {
	return v.value
}

func (v *NullableFrameInfo) Set(val *FrameInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFrameInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFrameInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrameInfo(val *FrameInfo) *NullableFrameInfo {
	return &NullableFrameInfo{value: val, isSet: true}
}

func (v NullableFrameInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrameInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


