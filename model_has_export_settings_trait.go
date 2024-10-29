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

// checks if the HasExportSettingsTrait type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HasExportSettingsTrait{}

// HasExportSettingsTrait struct for HasExportSettingsTrait
type HasExportSettingsTrait struct {
	// An array of export settings representing images to export from the node.
	ExportSettings []ExportSetting `json:"exportSettings,omitempty"`
}

// NewHasExportSettingsTrait instantiates a new HasExportSettingsTrait object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasExportSettingsTrait() *HasExportSettingsTrait {
	this := HasExportSettingsTrait{}
	return &this
}

// NewHasExportSettingsTraitWithDefaults instantiates a new HasExportSettingsTrait object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasExportSettingsTraitWithDefaults() *HasExportSettingsTrait {
	this := HasExportSettingsTrait{}
	return &this
}

// GetExportSettings returns the ExportSettings field value if set, zero value otherwise.
func (o *HasExportSettingsTrait) GetExportSettings() []ExportSetting {
	if o == nil || IsNil(o.ExportSettings) {
		var ret []ExportSetting
		return ret
	}
	return o.ExportSettings
}

// GetExportSettingsOk returns a tuple with the ExportSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasExportSettingsTrait) GetExportSettingsOk() ([]ExportSetting, bool) {
	if o == nil || IsNil(o.ExportSettings) {
		return nil, false
	}
	return o.ExportSettings, true
}

// HasExportSettings returns a boolean if a field has been set.
func (o *HasExportSettingsTrait) HasExportSettings() bool {
	if o != nil && !IsNil(o.ExportSettings) {
		return true
	}

	return false
}

// SetExportSettings gets a reference to the given []ExportSetting and assigns it to the ExportSettings field.
func (o *HasExportSettingsTrait) SetExportSettings(v []ExportSetting) {
	o.ExportSettings = v
}

func (o HasExportSettingsTrait) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HasExportSettingsTrait) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExportSettings) {
		toSerialize["exportSettings"] = o.ExportSettings
	}
	return toSerialize, nil
}

type NullableHasExportSettingsTrait struct {
	value *HasExportSettingsTrait
	isSet bool
}

func (v NullableHasExportSettingsTrait) Get() *HasExportSettingsTrait {
	return v.value
}

func (v *NullableHasExportSettingsTrait) Set(val *HasExportSettingsTrait) {
	v.value = val
	v.isSet = true
}

func (v NullableHasExportSettingsTrait) IsSet() bool {
	return v.isSet
}

func (v *NullableHasExportSettingsTrait) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasExportSettingsTrait(val *HasExportSettingsTrait) *NullableHasExportSettingsTrait {
	return &NullableHasExportSettingsTrait{value: val, isSet: true}
}

func (v NullableHasExportSettingsTrait) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasExportSettingsTrait) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


