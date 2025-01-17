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

// checks if the BaseShadowEffectBoundVariables type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseShadowEffectBoundVariables{}

// BaseShadowEffectBoundVariables The variables bound to a particular field on this shadow effect
type BaseShadowEffectBoundVariables struct {
	Radius *VariableAlias `json:"radius,omitempty"`
	Spread *VariableAlias `json:"spread,omitempty"`
	Color *VariableAlias `json:"color,omitempty"`
	OffsetX *VariableAlias `json:"offsetX,omitempty"`
	OffsetY *VariableAlias `json:"offsetY,omitempty"`
}

// NewBaseShadowEffectBoundVariables instantiates a new BaseShadowEffectBoundVariables object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseShadowEffectBoundVariables() *BaseShadowEffectBoundVariables {
	this := BaseShadowEffectBoundVariables{}
	return &this
}

// NewBaseShadowEffectBoundVariablesWithDefaults instantiates a new BaseShadowEffectBoundVariables object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseShadowEffectBoundVariablesWithDefaults() *BaseShadowEffectBoundVariables {
	this := BaseShadowEffectBoundVariables{}
	return &this
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *BaseShadowEffectBoundVariables) GetRadius() VariableAlias {
	if o == nil || IsNil(o.Radius) {
		var ret VariableAlias
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShadowEffectBoundVariables) GetRadiusOk() (*VariableAlias, bool) {
	if o == nil || IsNil(o.Radius) {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *BaseShadowEffectBoundVariables) HasRadius() bool {
	if o != nil && !IsNil(o.Radius) {
		return true
	}

	return false
}

// SetRadius gets a reference to the given VariableAlias and assigns it to the Radius field.
func (o *BaseShadowEffectBoundVariables) SetRadius(v VariableAlias) {
	o.Radius = &v
}

// GetSpread returns the Spread field value if set, zero value otherwise.
func (o *BaseShadowEffectBoundVariables) GetSpread() VariableAlias {
	if o == nil || IsNil(o.Spread) {
		var ret VariableAlias
		return ret
	}
	return *o.Spread
}

// GetSpreadOk returns a tuple with the Spread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShadowEffectBoundVariables) GetSpreadOk() (*VariableAlias, bool) {
	if o == nil || IsNil(o.Spread) {
		return nil, false
	}
	return o.Spread, true
}

// HasSpread returns a boolean if a field has been set.
func (o *BaseShadowEffectBoundVariables) HasSpread() bool {
	if o != nil && !IsNil(o.Spread) {
		return true
	}

	return false
}

// SetSpread gets a reference to the given VariableAlias and assigns it to the Spread field.
func (o *BaseShadowEffectBoundVariables) SetSpread(v VariableAlias) {
	o.Spread = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *BaseShadowEffectBoundVariables) GetColor() VariableAlias {
	if o == nil || IsNil(o.Color) {
		var ret VariableAlias
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShadowEffectBoundVariables) GetColorOk() (*VariableAlias, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *BaseShadowEffectBoundVariables) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given VariableAlias and assigns it to the Color field.
func (o *BaseShadowEffectBoundVariables) SetColor(v VariableAlias) {
	o.Color = &v
}

// GetOffsetX returns the OffsetX field value if set, zero value otherwise.
func (o *BaseShadowEffectBoundVariables) GetOffsetX() VariableAlias {
	if o == nil || IsNil(o.OffsetX) {
		var ret VariableAlias
		return ret
	}
	return *o.OffsetX
}

// GetOffsetXOk returns a tuple with the OffsetX field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShadowEffectBoundVariables) GetOffsetXOk() (*VariableAlias, bool) {
	if o == nil || IsNil(o.OffsetX) {
		return nil, false
	}
	return o.OffsetX, true
}

// HasOffsetX returns a boolean if a field has been set.
func (o *BaseShadowEffectBoundVariables) HasOffsetX() bool {
	if o != nil && !IsNil(o.OffsetX) {
		return true
	}

	return false
}

// SetOffsetX gets a reference to the given VariableAlias and assigns it to the OffsetX field.
func (o *BaseShadowEffectBoundVariables) SetOffsetX(v VariableAlias) {
	o.OffsetX = &v
}

// GetOffsetY returns the OffsetY field value if set, zero value otherwise.
func (o *BaseShadowEffectBoundVariables) GetOffsetY() VariableAlias {
	if o == nil || IsNil(o.OffsetY) {
		var ret VariableAlias
		return ret
	}
	return *o.OffsetY
}

// GetOffsetYOk returns a tuple with the OffsetY field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseShadowEffectBoundVariables) GetOffsetYOk() (*VariableAlias, bool) {
	if o == nil || IsNil(o.OffsetY) {
		return nil, false
	}
	return o.OffsetY, true
}

// HasOffsetY returns a boolean if a field has been set.
func (o *BaseShadowEffectBoundVariables) HasOffsetY() bool {
	if o != nil && !IsNil(o.OffsetY) {
		return true
	}

	return false
}

// SetOffsetY gets a reference to the given VariableAlias and assigns it to the OffsetY field.
func (o *BaseShadowEffectBoundVariables) SetOffsetY(v VariableAlias) {
	o.OffsetY = &v
}

func (o BaseShadowEffectBoundVariables) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaseShadowEffectBoundVariables) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Radius) {
		toSerialize["radius"] = o.Radius
	}
	if !IsNil(o.Spread) {
		toSerialize["spread"] = o.Spread
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.OffsetX) {
		toSerialize["offsetX"] = o.OffsetX
	}
	if !IsNil(o.OffsetY) {
		toSerialize["offsetY"] = o.OffsetY
	}
	return toSerialize, nil
}

type NullableBaseShadowEffectBoundVariables struct {
	value *BaseShadowEffectBoundVariables
	isSet bool
}

func (v NullableBaseShadowEffectBoundVariables) Get() *BaseShadowEffectBoundVariables {
	return v.value
}

func (v *NullableBaseShadowEffectBoundVariables) Set(val *BaseShadowEffectBoundVariables) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseShadowEffectBoundVariables) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseShadowEffectBoundVariables) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseShadowEffectBoundVariables(val *BaseShadowEffectBoundVariables) *NullableBaseShadowEffectBoundVariables {
	return &NullableBaseShadowEffectBoundVariables{value: val, isSet: true}
}

func (v NullableBaseShadowEffectBoundVariables) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseShadowEffectBoundVariables) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


