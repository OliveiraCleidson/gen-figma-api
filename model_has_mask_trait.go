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
)

// checks if the HasMaskTrait type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HasMaskTrait{}

// HasMaskTrait struct for HasMaskTrait
type HasMaskTrait struct {
	// Does this node mask sibling nodes in front of it?
	IsMask *bool `json:"isMask,omitempty"`
	// If this layer is a mask, this property describes the operation used to mask the layer's siblings. The value may be one of the following:  - ALPHA: the mask node's alpha channel will be used to determine the opacity of each pixel in the masked result. - VECTOR: if the mask node has visible fill paints, every pixel inside the node's fill regions will be fully visible in the masked result. If the mask has visible stroke paints, every pixel inside the node's stroke regions will be fully visible in the masked result. - LUMINANCE: the luminance value of each pixel of the mask node will be used to determine the opacity of that pixel in the masked result.
	MaskType *string `json:"maskType,omitempty"`
	// True if maskType is VECTOR. This field is deprecated; use maskType instead.
	// Deprecated
	IsMaskOutline *bool `json:"isMaskOutline,omitempty"`
}

// NewHasMaskTrait instantiates a new HasMaskTrait object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasMaskTrait() *HasMaskTrait {
	this := HasMaskTrait{}
	var isMask bool = false
	this.IsMask = &isMask
	var isMaskOutline bool = false
	this.IsMaskOutline = &isMaskOutline
	return &this
}

// NewHasMaskTraitWithDefaults instantiates a new HasMaskTrait object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasMaskTraitWithDefaults() *HasMaskTrait {
	this := HasMaskTrait{}
	var isMask bool = false
	this.IsMask = &isMask
	var isMaskOutline bool = false
	this.IsMaskOutline = &isMaskOutline
	return &this
}

// GetIsMask returns the IsMask field value if set, zero value otherwise.
func (o *HasMaskTrait) GetIsMask() bool {
	if o == nil || IsNil(o.IsMask) {
		var ret bool
		return ret
	}
	return *o.IsMask
}

// GetIsMaskOk returns a tuple with the IsMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasMaskTrait) GetIsMaskOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMask) {
		return nil, false
	}
	return o.IsMask, true
}

// HasIsMask returns a boolean if a field has been set.
func (o *HasMaskTrait) HasIsMask() bool {
	if o != nil && !IsNil(o.IsMask) {
		return true
	}

	return false
}

// SetIsMask gets a reference to the given bool and assigns it to the IsMask field.
func (o *HasMaskTrait) SetIsMask(v bool) {
	o.IsMask = &v
}

// GetMaskType returns the MaskType field value if set, zero value otherwise.
func (o *HasMaskTrait) GetMaskType() string {
	if o == nil || IsNil(o.MaskType) {
		var ret string
		return ret
	}
	return *o.MaskType
}

// GetMaskTypeOk returns a tuple with the MaskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasMaskTrait) GetMaskTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MaskType) {
		return nil, false
	}
	return o.MaskType, true
}

// HasMaskType returns a boolean if a field has been set.
func (o *HasMaskTrait) HasMaskType() bool {
	if o != nil && !IsNil(o.MaskType) {
		return true
	}

	return false
}

// SetMaskType gets a reference to the given string and assigns it to the MaskType field.
func (o *HasMaskTrait) SetMaskType(v string) {
	o.MaskType = &v
}

// GetIsMaskOutline returns the IsMaskOutline field value if set, zero value otherwise.
// Deprecated
func (o *HasMaskTrait) GetIsMaskOutline() bool {
	if o == nil || IsNil(o.IsMaskOutline) {
		var ret bool
		return ret
	}
	return *o.IsMaskOutline
}

// GetIsMaskOutlineOk returns a tuple with the IsMaskOutline field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *HasMaskTrait) GetIsMaskOutlineOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMaskOutline) {
		return nil, false
	}
	return o.IsMaskOutline, true
}

// HasIsMaskOutline returns a boolean if a field has been set.
func (o *HasMaskTrait) HasIsMaskOutline() bool {
	if o != nil && !IsNil(o.IsMaskOutline) {
		return true
	}

	return false
}

// SetIsMaskOutline gets a reference to the given bool and assigns it to the IsMaskOutline field.
// Deprecated
func (o *HasMaskTrait) SetIsMaskOutline(v bool) {
	o.IsMaskOutline = &v
}

func (o HasMaskTrait) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HasMaskTrait) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsMask) {
		toSerialize["isMask"] = o.IsMask
	}
	if !IsNil(o.MaskType) {
		toSerialize["maskType"] = o.MaskType
	}
	if !IsNil(o.IsMaskOutline) {
		toSerialize["isMaskOutline"] = o.IsMaskOutline
	}
	return toSerialize, nil
}

type NullableHasMaskTrait struct {
	value *HasMaskTrait
	isSet bool
}

func (v NullableHasMaskTrait) Get() *HasMaskTrait {
	return v.value
}

func (v *NullableHasMaskTrait) Set(val *HasMaskTrait) {
	v.value = val
	v.isSet = true
}

func (v NullableHasMaskTrait) IsSet() bool {
	return v.isSet
}

func (v *NullableHasMaskTrait) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasMaskTrait(val *HasMaskTrait) *NullableHasMaskTrait {
	return &NullableHasMaskTrait{value: val, isSet: true}
}

func (v NullableHasMaskTrait) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasMaskTrait) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

